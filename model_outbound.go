package mailafrica

import (
	"time"
)

// SendEmailRequest is the payload for POST /api/outbound/emails.
type SendEmailRequest struct {
	To           []string               `json:"to"`
	Cc           []string               `json:"cc,omitempty"`
	Bcc          []string               `json:"bcc,omitempty"`
	Subject      string                 `json:"subject"`
	HTMLBody     string                 `json:"html_body,omitempty"`
	TextBody     string                 `json:"text_body,omitempty"`
	Attachments  []Attachment           `json:"attachments,omitempty"`
	TemplateID   *int64                 `json:"template_id,omitempty"`
	Variables    map[string]string      `json:"variables,omitempty"`
	FromDomainID *int64                 `json:"from_domain_id,omitempty"`
	FromAddress  *string                `json:"from_address,omitempty"`
}

// BatchSendRequest is the payload for POST /api/outbound/emails/batch.
type BatchSendRequest struct {
	To           []string          `json:"to"`
	Subject      string            `json:"subject"`
	HTMLBody     string            `json:"html_body,omitempty"`
	TextBody     string            `json:"text_body,omitempty"`
	Attachments  []Attachment      `json:"attachments,omitempty"`
	TemplateID   *int64            `json:"template_id,omitempty"`
	Variables    map[string]string `json:"variables,omitempty"`
	FromDomainID *int64            `json:"from_domain_id,omitempty"`
	FromAddress  *string           `json:"from_address,omitempty"`
}

// Attachment represents an email attachment.
type Attachment struct {
	Filename    string `json:"filename"`
	DataBase64  string `json:"data_base64"`
	ContentType string `json:"content_type,omitempty"`
}

// SentMessage represents an outbound email.
type SentMessage struct {
	ID                int64     `json:"id"`
	UserID            int64     `json:"user_id"`
	FromAddress       string    `json:"from_address"`
	ToAddresses       []string  `json:"to_addresses"`
	CcAddresses       []string  `json:"cc_addresses,omitempty"`
	BccAddresses      []string  `json:"bcc_addresses,omitempty"`
	Subject           string    `json:"subject"`
	HTMLBody          *string   `json:"html_body,omitempty"`
	TextBody          *string   `json:"text_body,omitempty"`
	Status            string    `json:"status"`
	ErrorCode         *string   `json:"error_code,omitempty"`
	ProviderMessageID *string   `json:"provider_message_id,omitempty"`
	AmountTZS         int64     `json:"amount_tzs"`
	CreatedAt         time.Time `json:"created_at"`
}

// MessageDetail represents a sent email with recipient statuses.
type MessageDetail struct {
	Message    SentMessage      `json:"message"`
	Recipients []RecipientStatus `json:"recipients"`
}

// RecipientStatus represents the delivery status for a single recipient.
type RecipientStatus struct {
	ID           int64     `json:"id"`
	MessageID    int64     `json:"message_id"`
	Recipient    string    `json:"recipient"`
	Status       string    `json:"status"`
	ProviderCode *string   `json:"provider_code,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

// BatchResult represents the result of a batch send.
type BatchResult struct {
	Total    int          `json:"total"`
	Sent     int          `json:"sent"`
	Failed   int          `json:"failed"`
	Messages []SentMessage `json:"messages"`
}

// TemplateRequest is the payload for creating/updating a template.
type TemplateRequest struct {
	Name     string `json:"name"`
	Subject  string `json:"subject"`
	HTMLBody string `json:"html_body,omitempty"`
	TextBody string `json:"text_body,omitempty"`
}

// Template represents an email template.
type Template struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Name      string    `json:"name"`
	Subject   string    `json:"subject"`
	HTMLBody  string    `json:"html_body,omitempty"`
	TextBody  string    `json:"text_body,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
