package mailafrica

import (
	"errors"
	"fmt"
)

// APIError represents an API error response.
type APIError struct {
	Code       string
	Message    string
	HTTPStatus int
	RequestID  string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("mailafrica: %s (status=%d, request_id=%s): %s", e.Code, e.HTTPStatus, e.RequestID, e.Message)
}

// Sentinel errors for common failure modes.
var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrRateLimited         = errors.New("rate limited")
	ErrNotVerified         = errors.New("account not verified")
	ErrAccountDisabled     = errors.New("account disabled")
	ErrNotFound            = errors.New("not found")
)

// IsInsufficientBalance reports whether err is an insufficient balance error.
func IsInsufficientBalance(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code == "INSUFFICIENT_BALANCE"
	}
	return errors.Is(err, ErrInsufficientBalance)
}

// IsRateLimited reports whether err is a rate limit error.
func IsRateLimited(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code == "RATE_LIMITED"
	}
	return errors.Is(err, ErrRateLimited)
}

// IsNotVerified reports whether err is a not verified error.
func IsNotVerified(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code == "NOT_VERIFIED"
	}
	return errors.Is(err, ErrNotVerified)
}

// IsAccountDisabled reports whether err is an account disabled error.
func IsAccountDisabled(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code == "ACCOUNT_DISABLED"
	}
	return errors.Is(err, ErrAccountDisabled)
}

// IsNotFound reports whether err is a not found error.
func IsNotFound(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code == "NOT_FOUND"
	}
	return errors.Is(err, ErrNotFound)
}

// AsAPIError returns the underlying *APIError if present.
func AsAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	return apiErr, errors.As(err, &apiErr)
}
