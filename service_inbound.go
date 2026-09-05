package mailafrica

import (
	"context"
	"net/http"
)

// CreateAddress creates a new inbound email address.
func (c *Client) CreateAddress(ctx context.Context, req CreateAddressRequest) (*Address, error) {
	var addr Address
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/inbound/addresses", req, &addr)
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

// ListAddresses returns all inbound addresses for the user.
func (c *Client) ListAddresses(ctx context.Context) ([]Address, error) {
	var addrs []Address
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/inbound/addresses", nil, &addrs)
	if err != nil {
		return nil, err
	}
	return addrs, nil
}

// DeleteAddress deletes an inbound address.
func (c *Client) DeleteAddress(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodDelete, c.cfg.BaseURL+"/api/inbound/addresses/"+itoa(id), nil, nil)
	return err
}

// ListMessages returns inbound messages with optional filtering and pagination.
func (c *Client) ListMessages(ctx context.Context, opts MessageListOpts) ([]Message, *Pagination, error) {
	opts.applyDefaults()
	url := c.cfg.BaseURL + "/api/inbound/messages?page=" + itoa(int64(opts.Page)) + "&per_page=" + itoa(int64(opts.PerPage))
	if opts.AddressID != nil {
		url += "&address_id=" + itoa(*opts.AddressID)
	}
	if opts.Unread != nil {
		if *opts.Unread {
			url += "&unread=true"
		}
	}

	var resp struct {
		Success    bool        `json:"success"`
		Data       []Message   `json:"data"`
		Pagination *Pagination `json:"pagination"`
	}

	_, err := c.doJSON(ctx, http.MethodGet, url, nil, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Data, resp.Pagination, nil
}

// GetMessage retrieves a single inbound message.
func (c *Client) GetMessage(ctx context.Context, id int64) (*Message, error) {
	var msg Message
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/inbound/messages/"+itoa(id), nil, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

// MarkMessageRead marks a message as read.
func (c *Client) MarkMessageRead(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodPatch, c.cfg.BaseURL+"/api/inbound/messages/"+itoa(id)+"/read", nil, nil)
	return err
}

// CreateInboundDomain adds a new inbound receiving domain.
func (c *Client) CreateInboundDomain(ctx context.Context, domain string) (*DomainWithVerification, error) {
	var resp DomainWithVerification
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/inbound/domains", AddInboundDomainRequest{Domain: domain}, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInboundDomains returns all inbound receiving domains.
func (c *Client) ListInboundDomains(ctx context.Context) ([]InboundDomain, error) {
	var domains []InboundDomain
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/inbound/domains", nil, &domains)
	if err != nil {
		return nil, err
	}
	return domains, nil
}

// VerifyInboundDomain initiates domain verification.
func (c *Client) VerifyInboundDomain(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/inbound/domains/"+itoa(id)+"/verify", nil, nil)
	return err
}

// DeleteInboundDomain deletes an inbound receiving domain.
func (c *Client) DeleteInboundDomain(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodDelete, c.cfg.BaseURL+"/api/inbound/domains/"+itoa(id), nil, nil)
	return err
}
