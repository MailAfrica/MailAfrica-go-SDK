package mailafrica

import "time"

// CreateCredentialRequest is the payload for POST /api/sandbox/credentials.
type CreateCredentialRequest struct {
	Scopes *string `json:"scopes,omitempty"`
}

// Credential represents a sandbox API credential.
type Credential struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	Scopes       *string   `json:"scopes,omitempty"`
	Revoked      bool      `json:"revoked"`
	CreatedAt    time.Time `json:"created_at"`
}

// SMTPSandboxCredentials represents SMTP credentials for the sandbox.
type SMTPSandboxCredentials struct {
	Host        string     `json:"host"`
	Port        int        `json:"port"`
	Username    string     `json:"username"`
	Password    *string    `json:"password,omitempty"`
	PasswordSet bool       `json:"password_set"`
	GeneratedAt *time.Time `json:"credentials_generated_at,omitempty"`
}

// SandboxMessage represents a sandbox email message.
type SandboxMessage struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	From        string    `json:"from_addr"`
	To          string    `json:"to_addr"`
	Subject     string    `json:"subject"`
	TextBody    *string   `json:"text_body,omitempty"`
	HTMLBody    *string   `json:"html_body,omitempty"`
	Headers     any       `json:"headers"`
	Attachments any       `json:"attachments"`
	ReceivedAt  time.Time `json:"received_at"`
}
