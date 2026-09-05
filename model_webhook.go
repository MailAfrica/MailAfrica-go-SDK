package mailafrica

import "time"

// CreateWebhookRequest is the payload for POST /api/webhook/webhooks.
type CreateWebhookRequest struct {
	AddressID int64  `json:"address_id"`
	URL       string `json:"url"`
	Secret    string `json:"secret,omitempty"`
}

// Webhook represents a webhook configuration.
type Webhook struct {
	ID        int64     `json:"id"`
	AddressID int64     `json:"address_id"`
	URL       string    `json:"url"`
	Secret    string    `json:"secret"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

// WebhookDelivery represents a webhook delivery attempt.
type WebhookDelivery struct {
	ID          int64      `json:"id"`
	WebhookID   int64      `json:"webhook_id"`
	MessageID   int64      `json:"message_id"`
	StatusCode  *int       `json:"status_code,omitempty"`
	Attempt     int        `json:"attempt"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
	Status      string     `json:"status"`
	NextRetryAt *time.Time `json:"next_retry_at,omitempty"`
	LastError   *string    `json:"last_error,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}
