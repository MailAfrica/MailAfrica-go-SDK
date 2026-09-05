package mailafrica

import (
	"context"
	"net/http"
)

// GetBalance retrieves the user's wallet balance.
func (c *Client) GetBalance(ctx context.Context) (*BalanceResponse, error) {
	var resp BalanceResponse
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/billing/balance", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InitiateTopup initiates a top-up payment.
func (c *Client) InitiateTopup(ctx context.Context, amount int64) (*TopupResponse, error) {
	var resp TopupResponse
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/billing/topup", TopupRequest{AmountTZS: amount}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InitiatePhoneTopup initiates a top-up via phone payment.
func (c *Client) InitiatePhoneTopup(ctx context.Context, amount int64) (*TopupResponse, error) {
	var resp TopupResponse
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/billing/topup/phone", TopupRequest{AmountTZS: amount}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
