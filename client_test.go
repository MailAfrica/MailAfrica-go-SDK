package mailafrica

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewDefaults(t *testing.T) {
	c := New(Config{})
	if c.cfg.BaseURL != defaultBaseURL {
		t.Errorf("expected base URL %q, got %q", defaultBaseURL, c.cfg.BaseURL)
	}
	if c.cfg.Timeout != defaultTimeout {
		t.Errorf("expected timeout %v, got %v", defaultTimeout, c.cfg.Timeout)
	}
	if c.cfg.UserAgent != defaultUserAgent {
		t.Errorf("expected user agent %q, got %q", defaultUserAgent, c.cfg.UserAgent)
	}
}

func TestDoSuccess(t *testing.T) {
	var authHeader string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"success":true,"message":"ok","data":{"id":1,"name":"Test"},"request_id":"req-1","timestamp":"2024-01-01T00:00:00Z"}`)
	}))
	defer srv.Close()

	c := New(Config{BaseURL: srv.URL, JWT: "test-jwt"})
	var resp struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	req, _ := http.NewRequestWithContext(context.Background(), "GET", srv.URL+"/test", nil)
	_, err := c.Do(context.Background(), req, &resp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if authHeader != "Bearer test-jwt" {
		t.Errorf("expected Bearer auth, got %q", authHeader)
	}
	if resp.ID != 1 || resp.Name != "Test" {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestDoAPIKey(t *testing.T) {
	var authHeader string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader = r.Header.Get("X-API-Key")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"success":true,"data":null}`)
	}))
	defer srv.Close()

	c := New(Config{BaseURL: srv.URL, APIKey: "MAIL_test_key"})
	req, _ := http.NewRequestWithContext(context.Background(), "GET", srv.URL+"/test", nil)
	_, err := c.Do(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if authHeader != "MAIL_test_key" {
		t.Errorf("expected API key auth, got %q", authHeader)
	}
}

func TestDoError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprint(w, `{"success":false,"message":"too many requests","errors":[{"code":"RATE_LIMITED","message":"too many requests"}],"request_id":"req-2","timestamp":"2024-01-01T00:00:00Z"}`)
	}))
	defer srv.Close()

	c := New(Config{BaseURL: srv.URL})
	req, _ := http.NewRequestWithContext(context.Background(), "GET", srv.URL+"/test", nil)
	_, err := c.Do(context.Background(), req, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.Code != "RATE_LIMITED" {
		t.Errorf("expected code RATE_LIMITED, got %s", apiErr.Code)
	}
	if apiErr.HTTPStatus != http.StatusTooManyRequests {
		t.Errorf("expected status 429, got %d", apiErr.HTTPStatus)
	}
	if apiErr.RequestID != "req-2" {
		t.Errorf("expected request_id req-2, got %s", apiErr.RequestID)
	}
}

func TestDoTokenRefresh(t *testing.T) {
	var refreshCalled int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "Bearer old-jwt" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, `{"success":false,"message":"unauthorized","errors":[{"code":"UNAUTHORIZED"}],"request_id":"req-401","timestamp":"2024-01-01T00:00:00Z"}`)
			return
		}
		if auth == "Bearer new-jwt" {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"success":true,"data":{"id":1},"request_id":"req-ok","timestamp":"2024-01-01T00:00:00Z"}`)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	c := New(Config{
		BaseURL: srv.URL,
		JWT:     "old-jwt",
		TokenRefresher: func(ctx context.Context) (string, error) {
			atomic.AddInt32(&refreshCalled, 1)
			return "new-jwt", nil
		},
	})

	var resp struct{ ID int64 }
	req, _ := http.NewRequestWithContext(context.Background(), "GET", srv.URL+"/test", nil)
	_, err := c.Do(context.Background(), req, &resp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if atomic.LoadInt32(&refreshCalled) != 1 {
		t.Error("expected TokenRefresher to be called once")
	}
	if resp.ID != 1 {
		t.Errorf("expected id 1, got %d", resp.ID)
	}
}

func TestDoContextCancel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
	}))
	defer srv.Close()

	c := New(Config{BaseURL: srv.URL, Timeout: 5 * time.Second})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	req, _ := http.NewRequestWithContext(context.Background(), "GET", srv.URL+"/test", nil)
	_, err := c.Do(ctx, req, nil)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
}

func TestDoHooks(t *testing.T) {
	var reqCount, respCount, errCount int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"success":true,"data":null}`)
	}))
	defer srv.Close()

	c := New(Config{
		BaseURL: srv.URL,
		Hooks: &Hooks{
			OnRequest: func(req *http.Request) { reqCount++ },
			OnResponse: func(resp *http.Response, d time.Duration) { respCount++ },
			OnError:    func(err error) { errCount++ },
		},
	})
	req, _ := http.NewRequestWithContext(context.Background(), "GET", srv.URL+"/test", nil)
	_, err := c.Do(context.Background(), req, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reqCount != 1 {
		t.Errorf("expected 1 request hook call, got %d", reqCount)
	}
	if respCount != 1 {
		t.Errorf("expected 1 response hook call, got %d", respCount)
	}
	if errCount != 0 {
		t.Errorf("expected 0 error hook calls, got %d", errCount)
	}
}

func TestIsInsufficientBalance(t *testing.T) {
	err := &APIError{Code: "INSUFFICIENT_BALANCE", Message: "no funds", HTTPStatus: 402}
	if !IsInsufficientBalance(err) {
		t.Error("expected IsInsufficientBalance to be true")
	}
}

func TestIsRateLimited(t *testing.T) {
	err := &APIError{Code: "RATE_LIMITED", Message: "slow down", HTTPStatus: 429}
	if !IsRateLimited(err) {
		t.Error("expected IsRateLimited to be true")
	}
}

func TestIsNotFound(t *testing.T) {
	err := &APIError{Code: "NOT_FOUND", Message: "missing", HTTPStatus: 404}
	if !IsNotFound(err) {
		t.Error("expected IsNotFound to be true")
	}
}

func TestDoMarshalError(t *testing.T) {
	c := New(Config{BaseURL: "http://invalid"})
	_, err := c.doJSON(context.Background(), "GET", "http://invalid", func() {}, nil)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestListOptsDefaults(t *testing.T) {
	opts := ListOpts{}
	opts.applyDefaults()
	if opts.Page != 1 {
		t.Errorf("expected default page 1, got %d", opts.Page)
	}
	if opts.PerPage != 25 {
		t.Errorf("expected default per_page 25, got %d", opts.PerPage)
	}
}

func TestListOptsCap(t *testing.T) {
	opts := ListOpts{Page: 5, PerPage: 200}
	opts.applyDefaults()
	if opts.PerPage != 100 {
		t.Errorf("expected capped per_page 100, got %d", opts.PerPage)
	}
}
