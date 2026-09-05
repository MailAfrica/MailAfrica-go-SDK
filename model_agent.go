package mailafrica

import "time"

// AgentConfig represents the AI auto-reply configuration for an address.
type AgentConfig struct {
	AddressID         int64     `json:"address_id"`
	UserID            int64     `json:"user_id"`
	Mode              string    `json:"mode"`
	Persona           *string   `json:"persona,omitempty"`
	Enabled           bool      `json:"enabled"`
	ReplyFromDomainID *int64    `json:"reply_from_domain_id,omitempty"`
	ReplyFromAddress  *string   `json:"reply_from_address,omitempty"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// UpdateAgentConfigRequest is the payload for PUT /api/agent/configs/{id}.
type UpdateAgentConfigRequest struct {
	Mode              string  `json:"mode"`
	Persona           *string `json:"persona,omitempty"`
	Enabled           *bool   `json:"enabled,omitempty"`
	ReplyFromDomainID *int64  `json:"reply_from_domain_id,omitempty"`
	ReplyFromAddress  *string `json:"reply_from_address,omitempty"`
}

// AgentDraftRequest is the payload for POST /api/agent/configs/{id}/draft.
type AgentDraftRequest struct {
	Subject string `json:"subject"`
	Body    string `json:"text_body"`
}

// AgentDraft represents a generated AI reply draft.
type AgentDraft struct {
	Draft string `json:"draft"`
}
