package mailafrica

import (
	"context"
	"net/http"
)

// AddSendingDomain adds a new sending domain.
func (c *Client) AddSendingDomain(ctx context.Context, req AddSendingDomainRequest) (*AddSendingDomainResponse, error) {
	var resp AddSendingDomainResponse
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/domains", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSendingDomains returns all sending domains.
func (c *Client) ListSendingDomains(ctx context.Context) ([]SendingDomain, error) {
	var domains []SendingDomain
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/domains", nil, &domains)
	if err != nil {
		return nil, err
	}
	return domains, nil
}

// VerifySendingDomain initiates verification for a sending domain.
func (c *Client) VerifySendingDomain(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/domains/"+itoa(id)+"/verify", nil, nil)
	return err
}

// DeleteSendingDomain deletes a sending domain.
func (c *Client) DeleteSendingDomain(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodDelete, c.cfg.BaseURL+"/api/domains/"+itoa(id), nil, nil)
	return err
}

// CreateSenderAddress creates a new sender address for a domain.
func (c *Client) CreateSenderAddress(ctx context.Context, domainID int64, localPart string) (*SenderAddress, error) {
	var addr SenderAddress
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/domains/"+itoa(domainID)+"/senders", CreateSenderAddressRequest{LocalPart: localPart}, &addr)
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

// ListSenderAddresses returns all sender addresses.
func (c *Client) ListSenderAddresses(ctx context.Context) ([]SenderAddress, error) {
	var addrs []SenderAddress
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/domains/senders", nil, &addrs)
	if err != nil {
		return nil, err
	}
	return addrs, nil
}

// DeleteSenderAddress deletes a sender address.
func (c *Client) DeleteSenderAddress(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodDelete, c.cfg.BaseURL+"/api/domains/senders/"+itoa(id), nil, nil)
	return err
}
