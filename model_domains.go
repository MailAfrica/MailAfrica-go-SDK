package mailafrica

import (
	"time"
)

// AddSendingDomainRequest is the payload for POST /api/domains.
type AddSendingDomainRequest struct {
	Domain        string `json:"domain"`
	FromLocalPart string `json:"from_local_part,omitempty"`
}

// AddSendingDomainResponse is returned when adding a sending domain.
type AddSendingDomainResponse struct {
	Domain        SendingDomain  `json:"domain"`
	DNSRecords    FullDnsRecords `json:"dns_records"`
}

// SendingDomain represents a sending domain.
type SendingDomain struct {
	ID                int64      `json:"id"`
	UserID            int64      `json:"user_id"`
	Domain            string     `json:"domain"`
	VerificationToken string     `json:"verification_token,omitempty"`
	Purpose           string     `json:"purpose"`
	Status            string     `json:"status"`
	AgentKey          *string    `json:"agent_key,omitempty"`
	DomainKey         *string    `json:"domain_key,omitempty"`
	BouncePrefix      string     `json:"bounce_prefix"`
	FromLocalPart     string     `json:"from_local_part"`
	DkimHost          *string    `json:"dkim_host,omitempty"`
	DkimValue         *string    `json:"dkim_value,omitempty"`
	CnameHost         *string    `json:"cname_host,omitempty"`
	CnameValue        *string    `json:"cname_value,omitempty"`
	VerifiedAt        *time.Time `json:"verified_at,omitempty"`
	LastCheckAt       *time.Time `json:"last_check_at,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
}

// SenderAddress represents a sender address for a domain.
type SenderAddress struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	DomainID  int64     `json:"domain_id"`
	Domain    string    `json:"domain"`
	LocalPart string    `json:"local_part"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateSenderAddressRequest is the payload for POST /api/domains/{id}/senders.
type CreateSenderAddressRequest struct {
	LocalPart string `json:"local_part"`
}

// DNSRecord describes a DNS record.
type DNSRecord struct {
	Type  string `json:"type"`
	Host  string `json:"host"`
	Value string `json:"value"`
}

// FullDnsRecords contains DKIM, SPF, and DMARC records.
type FullDnsRecords struct {
	DKIM  DNSRecord `json:"dkim"`
	SPF   DNSRecord `json:"spf"`
	DMARC DNSRecord `json:"dmarc"`
}

// VerificationResult represents the result of domain verification checks.
type VerificationResult struct {
	DKIM  bool `json:"dkim"`
	SPF   bool `json:"spf"`
	DMARC bool `json:"dmarc"`
}
