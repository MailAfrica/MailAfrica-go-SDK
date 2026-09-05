package mailafrica

import "time"

// RegisterRequest is the payload for POST /api/auth/register.
type RegisterRequest struct {
	Email       *string `json:"email,omitempty"`
	Phone       *string `json:"phone_number,omitempty"`
	Password    string  `json:"password"`
	Name        string  `json:"name"`
	CompanyName string  `json:"company_name"`
}

// LoginRequest is the payload for POST /api/auth/login.
type LoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

// GoogleLoginRequest is the payload for POST /api/auth/google.
type GoogleLoginRequest struct {
	IDToken string `json:"id_token"`
}

// AuthResponse is returned by register, login, refresh, and Google login.
type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	User         User   `json:"user"`
}

// UpdateProfileRequest is the payload for PATCH /api/auth/me.
type UpdateProfileRequest struct {
	Name        *string `json:"name,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
}

// SetEmailRequest is the payload for POST /api/auth/email.
type SetEmailRequest struct {
	Email string `json:"email"`
}

// SetPhoneRequest is the payload for POST /api/auth/phone.
type SetPhoneRequest struct {
	Phone string `json:"phone_number"`
}

// VerifyEmailRequest is the payload for POST /api/auth/email/verify.
type VerifyEmailRequest struct {
	Token string `json:"token"`
}

// VerifyPhoneRequest is the payload for POST /api/auth/phone/verify.
type VerifyPhoneRequest struct {
	Code string `json:"code"`
}

// User represents the authenticated user profile.
type User struct {
	ID              int64      `json:"id"`
	Email           *string    `json:"email,omitempty"`
	Phone           *string    `json:"phone_number,omitempty"`
	Name            string     `json:"name"`
	CompanyName     *string    `json:"company_name,omitempty"`
	IsAdmin         bool       `json:"is_admin"`
	DisabledAt      *time.Time `json:"disabled_at,omitempty"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	PhoneVerifiedAt *time.Time `json:"phone_verified_at,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
}
