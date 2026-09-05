package mailafrica

import "time"

// Address represents an inbound email address.
type Address struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	LocalPart     string    `json:"local_part"`
	Label         *string   `json:"label,omitempty"`
	DomainID      *int64    `json:"domain_id,omitempty"`
	RetentionDays int       `json:"retention_days"`
	CreatedAt     time.Time `json:"created_at"`
}

// CreateAddressRequest is the payload for POST /api/inbound/addresses.
type CreateAddressRequest struct {
	LocalPart string  `json:"local_part"`
	Label     *string `json:"label,omitempty"`
	DomainID  *int64  `json:"domain_id,omitempty"`
}

// Message represents an inbound email message.
type Message struct {
	ID         int64      `json:"id"`
	AddressID  int64      `json:"address_id"`
	From       string     `json:"from_addr"`
	To         string     `json:"to_addr"`
	Subject    string     `json:"subject"`
	TextBody   *string    `json:"text_body,omitempty"`
	HTMLBody   *string    `json:"html_body,omitempty"`
	Headers    any        `json:"headers"`
	Attachments any       `json:"attachments"`
	IsRead     bool       `json:"is_read"`
	ReceivedAt time.Time  `json:"received_at"`
}

// MessageListOpts configures pagination for inbound messages.
type MessageListOpts struct {
	ListOpts
	AddressID *int64
	Unread    *bool
}

// InboundDomain represents an inbound receiving domain.
type InboundDomain struct {
	ID                int64      `json:"id"`
	UserID            int64      `json:"user_id"`
	Domain            string     `json:"domain"`
	VerificationToken string     `json:"verification_token,omitempty"`
	VerifiedAt        *time.Time `json:"verified_at,omitempty"`
	LastCheckAt       *time.Time `json:"last_check_at,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
}

// AddInboundDomainRequest is the payload for POST /api/inbound/domains.
type AddInboundDomainRequest struct {
	Domain string `json:"domain"`
}

// DomainWithVerification is returned when creating an inbound domain.
type DomainWithVerification struct {
	InboundDomain
	VerificationRecord  VerificationRecord  `json:"verification_record"`
	VerificationRecords []VerificationRecord `json:"verification_records"`
}

// VerificationRecord describes a DNS verification record.
type VerificationRecord struct {
	Type  string `json:"type"`
	Host  string `json:"host"`
	Value string `json:"value"`
}
