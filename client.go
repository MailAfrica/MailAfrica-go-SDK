package mailafrica

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/MailAfrica/go-sdk/internal/pjson"
)

const (
	defaultBaseURL  = "https://api.mailafrica.online"
	defaultTimeout  = 30 * time.Second
	defaultUserAgent = "mailafrica-go/0.1.0"
)

// Config holds SDK configuration.
type Config struct {
	BaseURL        string
	APIKey         string
	JWT            string
	Timeout        time.Duration
	UserAgent      string
	TokenRefresher func(ctx context.Context) (string, error)
	Hooks          *Hooks
}

func (c *Config) defaults() {
	if c.BaseURL == "" {
		c.BaseURL = defaultBaseURL
	}
	if c.Timeout <= 0 {
		c.Timeout = defaultTimeout
	}
	if c.UserAgent == "" {
		c.UserAgent = defaultUserAgent
	}
}

// Client is the MailAfrica API client.
type Client struct {
	cfg    Config
	http   *http.Client
	mu     sync.RWMutex
	jwt    string
}

// New creates a new MailAfrica client.
func New(cfg Config) *Client {
	cfg.defaults()
	return &Client{
		cfg:  cfg,
		http: &http.Client{Timeout: cfg.Timeout},
		jwt:  cfg.JWT,
	}
}

// Do executes an HTTP request with auth headers and envelope unwrapping.
func (c *Client) Do(ctx context.Context, req *http.Request, v any) (*APIResponse, error) {
	if c.cfg.Hooks != nil && c.cfg.Hooks.OnRequest != nil {
		c.cfg.Hooks.OnRequest(req)
	}

	c.setAuthHeaders(req)

	resp, err := c.http.Do(req.WithContext(ctx))
	if err != nil {
		if c.cfg.Hooks != nil && c.cfg.Hooks.OnError != nil {
			c.cfg.Hooks.OnError(err)
		}
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	if c.cfg.Hooks != nil && c.cfg.Hooks.OnResponse != nil {
		c.cfg.Hooks.OnResponse(resp, 0)
	}

	apiResp, err := decodeEnvelope(resp, v)
	if err != nil {
		if c.cfg.Hooks != nil && c.cfg.Hooks.OnError != nil {
			c.cfg.Hooks.OnError(err)
		}
		return nil, err
	}

	if apiResp.Success {
		return apiResp, nil
	}

	apiErr := &APIError{
		Code:       apiResp.Errors[0].Code,
		Message:    apiResp.Message,
		HTTPStatus: resp.StatusCode,
		RequestID:  apiResp.RequestID,
	}

	if apiErr.HTTPStatus == http.StatusUnauthorized && c.cfg.TokenRefresher != nil {
		if err := tryRefreshRetry(ctx, req, apiErr, c, v); err == nil {
			return nil, nil
		} else if !errors.Is(err, apiErr) {
			return nil, err
		}
	}

	return nil, apiErr
}

func (c *Client) setAuthHeaders(req *http.Request) {
	if c.cfg.APIKey != "" {
		req.Header.Set("X-API-Key", c.cfg.APIKey)
		return
	}

	c.mu.RLock()
	jwt := c.jwt
	c.mu.RUnlock()

	if jwt != "" {
		req.Header.Set("Authorization", "Bearer "+jwt)
	}
}

func tryRefreshRetry(ctx context.Context, req *http.Request, apiErr *APIError, c *Client, v any) error {
	if apiErr.HTTPStatus != http.StatusUnauthorized {
		return apiErr
	}

	c.mu.RLock()
	refresher := c.cfg.TokenRefresher
	c.mu.RUnlock()

	if refresher == nil {
		return apiErr
	}

	newJWT, err := refresher(ctx)
	if err != nil {
		return fmt.Errorf("%w (token refresh failed: %v)", apiErr, err)
	}

	c.mu.Lock()
	c.jwt = newJWT
	c.mu.Unlock()

	req2 := req.Clone(ctx)
	c.setAuthHeaders(req2)

	resp, err := c.http.Do(req2.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("%w (retry failed: %v)", apiErr, err)
	}
	defer resp.Body.Close()

	apiResp, err := decodeEnvelope(resp, v)
	if err != nil {
		return err
	}

	if !apiResp.Success {
		return &APIError{
			Code:       apiResp.Errors[0].Code,
			Message:    apiResp.Message,
			HTTPStatus: resp.StatusCode,
			RequestID:  apiResp.RequestID,
		}
	}

	return nil
}

func decodeEnvelope(resp *http.Response, v any) (*APIResponse, error) {
	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if v != nil && apiResp.Success && apiResp.Data != nil {
		b, err := json.Marshal(apiResp.Data)
		if err != nil {
			return nil, fmt.Errorf("marshal data: %w", err)
		}
		if err := json.Unmarshal(b, v); err != nil {
			return nil, fmt.Errorf("decode data: %w", err)
		}
	}

	return &apiResp, nil
}

// APIResponse mirrors the backend envelope.
type APIResponse struct {
	Success    bool           `json:"success"`
	Message    string         `json:"message,omitempty"`
	Data       any            `json:"data,omitempty"`
	Errors     []FieldError   `json:"errors,omitempty"`
	Pagination *Pagination    `json:"pagination,omitempty"`
	RequestID  string         `json:"request_id,omitempty"`
	Timestamp  pjson.Time     `json:"timestamp"`
}

// FieldError describes a single validation or field-level issue.
type FieldError struct {
	Field   string `json:"field,omitempty"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) doJSON(ctx context.Context, method, url string, body, v any) (*APIResponse, error) {
	var bodyReader *bytes.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.cfg.UserAgent)

	return c.Do(ctx, req, v)
}

func itoa(i int64) string {
	return fmt.Sprintf("%d", i)
}
