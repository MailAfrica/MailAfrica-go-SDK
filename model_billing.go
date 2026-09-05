package mailafrica

import "time"

// BalanceResponse represents the user's wallet balance.
type BalanceResponse struct {
	BalanceTZS int64 `json:"balance_tzs"`
}

// TopupRequest is the payload for POST /api/billing/topup.
type TopupRequest struct {
	AmountTZS int64 `json:"amount_tzs"`
}

// TopupResponse is returned after initiating a top-up.
type TopupResponse struct {
	Topup             *Topup `json:"topup"`
	CheckoutURL       string `json:"checkout_url,omitempty"`
	PaymentLinkURL    string `json:"payment_link_url,omitempty"`
	ProviderReference string `json:"provider_reference,omitempty"`
}

// Topup represents a top-up transaction.
type Topup struct {
	ID                int64     `json:"id"`
	UserID            int64     `json:"user_id"`
	AmountTZS         int64     `json:"amount_tzs"`
	Status            string    `json:"status"`
	ProviderReference *string   `json:"provider_reference,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
}
