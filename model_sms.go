package mailafrica

import "time"

// CreateSMSNotificationRequest is the payload for POST /api/sms/notifications.
type CreateSMSNotificationRequest struct {
	AddressID   int64  `json:"address_id"`
	PhoneNumber string `json:"phone_number"`
	APIKey      string `json:"api_key"`
}

// SMSNotification represents an SMS notification configuration.
type SMSNotification struct {
	ID          int64     `json:"id"`
	AddressID   int64     `json:"address_id"`
	PhoneNumber string    `json:"phone_number"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}

// SMSNotificationWithKey includes the plaintext API key returned on creation.
type SMSNotificationWithKey struct {
	SMSNotification
	APIKey string `json:"api_key"`
}

// SMSDelivery represents an SMS delivery attempt.
type SMSDelivery struct {
	ID                int64     `json:"id"`
	NotificationID    int64     `json:"sms_notification_id"`
	MessageID         int64     `json:"message_id"`
	Status            string    `json:"status"`
	ErrorCode         *string   `json:"error_code,omitempty"`
	ProviderMessageID *string   `json:"provider_message_id,omitempty"`
	Attempt           int       `json:"attempt"`
	CreatedAt         time.Time `json:"created_at"`
}
