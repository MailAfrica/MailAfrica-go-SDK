package mailafrica

import "time"

// ComplianceProfile represents the user's compliance/PDPC profile.
type ComplianceProfile struct {
	ID                    int64      `json:"id"`
	UserID                int64      `json:"user_id"`
	PDPCRegistered        bool       `json:"pdpc_registered"`
	PDPCCertificateNumber *string    `json:"pdpc_certificate_number,omitempty"`
	PDPCRegisteredAt      *time.Time `json:"pdpc_registered_at,omitempty"`
	DefaultRetentionDays  int        `json:"default_retention_days"`
	DataConsentAt         *time.Time `json:"data_consent_at,omitempty"`
	PrivacyPolicyVersion  string     `json:"privacy_policy_version"`
	UpdatedAt             time.Time  `json:"updated_at"`
}

// UpdateComplianceProfileRequest is the payload for PATCH /api/compliance/profile.
type UpdateComplianceProfileRequest struct {
	PDPCRegistered        *bool   `json:"pdpc_registered,omitempty"`
	PDPCCertificateNumber *string `json:"pdpc_certificate_number,omitempty"`
	PDPCRegisteredAt      *string `json:"pdpc_registered_at,omitempty"`
	DefaultRetentionDays  *int    `json:"default_retention_days,omitempty"`
}

// AuditExport represents the compliance audit export.
type AuditExport struct {
	PDPCRegistered        bool      `json:"pdpc_registered"`
	PDPCCertificateNumber string    `json:"pdpc_certificate_number,omitempty"`
	DefaultRetentionDays  int       `json:"default_retention_days"`
	AddressCount          int64     `json:"address_count"`
	MessageCount          int64     `json:"message_count"`
	GeneratedAt           time.Time `json:"generated_at"`
}
