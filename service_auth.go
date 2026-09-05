package mailafrica

import (
	"context"
	"net/http"
)

// Register creates a new user account.
func (c *Client) Register(ctx context.Context, req RegisterRequest) (*AuthResponse, error) {
	var resp AuthResponse
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/auth/register", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Login authenticates a user and returns a JWT.
func (c *Client) Login(ctx context.Context, req LoginRequest) (*AuthResponse, error) {
	var resp AuthResponse
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/auth/login", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GoogleLogin exchanges a Google ID token for a MailAfrica session.
func (c *Client) GoogleLogin(ctx context.Context, idToken string) (*AuthResponse, error) {
	var resp AuthResponse
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/auth/google", GoogleLoginRequest{IDToken: idToken}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Refresh exchanges a refresh token for a new access token.
func (c *Client) Refresh(ctx context.Context, refreshToken string) (*AuthResponse, error) {
	var resp AuthResponse
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/auth/refresh", struct {
		RefreshToken string `json:"refresh_token"`
	}{RefreshToken: refreshToken}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Me returns the current authenticated user profile.
func (c *Client) Me(ctx context.Context) (*User, error) {
	var user User
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/auth/me", nil, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateMe updates the caller's profile.
func (c *Client) UpdateMe(ctx context.Context, req UpdateProfileRequest) (*User, error) {
	var user User
	_, err := c.doJSON(ctx, http.MethodPatch, c.cfg.BaseURL+"/api/auth/me", req, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// SetEmail adds or replaces the caller's email and triggers verification.
func (c *Client) SetEmail(ctx context.Context, email string) (*User, error) {
	var user User
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/auth/email", SetEmailRequest{Email: email}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// VerifyEmail confirms an email via the verification token.
func (c *Client) VerifyEmail(ctx context.Context, token string) error {
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/auth/email/verify", VerifyEmailRequest{Token: token}, nil)
	return err
}

// ResendEmailVerification re-sends the verification email.
func (c *Client) ResendEmailVerification(ctx context.Context) error {
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/auth/email/resend", nil, nil)
	return err
}

// SetPhone adds or replaces the caller's phone and triggers OTP.
func (c *Client) SetPhone(ctx context.Context, phone string) (*User, error) {
	var user User
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/auth/phone", SetPhoneRequest{Phone: phone}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// VerifyPhone confirms a phone via the 6-digit OTP.
func (c *Client) VerifyPhone(ctx context.Context, code string) error {
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/auth/phone/verify", VerifyPhoneRequest{Code: code}, nil)
	return err
}

// ResendPhoneOTP re-sends the phone OTP.
func (c *Client) ResendPhoneOTP(ctx context.Context) error {
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/auth/phone/resend", nil, nil)
	return err
}
