package mailafrica

import (
	"context"
	"net/http"
)

// SendEmail sends a single email.
func (c *Client) SendEmail(ctx context.Context, req SendEmailRequest) (*SentMessage, error) {
	var msg SentMessage
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/outbound/emails", req, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

// BatchSend sends an email to multiple recipients.
func (c *Client) BatchSend(ctx context.Context, req BatchSendRequest) (*BatchResult, error) {
	var result BatchResult
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/outbound/emails/batch", req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListSentEmails returns sent emails with pagination.
func (c *Client) ListSentEmails(ctx context.Context, opts ListOpts) ([]SentMessage, *Pagination, error) {
	opts.applyDefaults()
	url := c.cfg.BaseURL + "/api/outbound/emails?page=" + itoa(int64(opts.Page)) + "&per_page=" + itoa(int64(opts.PerPage))

	var resp struct {
		Success    bool        `json:"success"`
		Data       []SentMessage `json:"data"`
		Pagination *Pagination `json:"pagination"`
	}

	_, err := c.doJSON(ctx, http.MethodGet, url, nil, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Data, resp.Pagination, nil
}

// GetSentEmail retrieves a single sent email with recipient statuses.
func (c *Client) GetSentEmail(ctx context.Context, id int64) (*MessageDetail, error) {
	var detail MessageDetail
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/outbound/emails/"+itoa(id), nil, &detail)
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

// CreateTemplate creates a new email template.
func (c *Client) CreateTemplate(ctx context.Context, req TemplateRequest) (*Template, error) {
	var tpl Template
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/outbound/templates", req, &tpl)
	if err != nil {
		return nil, err
	}
	return &tpl, nil
}

// ListTemplates returns all templates.
func (c *Client) ListTemplates(ctx context.Context) ([]Template, error) {
	var tpls []Template
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/outbound/templates", nil, &tpls)
	if err != nil {
		return nil, err
	}
	return tpls, nil
}

// GetTemplate retrieves a single template.
func (c *Client) GetTemplate(ctx context.Context, id int64) (*Template, error) {
	var tpl Template
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/outbound/templates/"+itoa(id), nil, &tpl)
	if err != nil {
		return nil, err
	}
	return &tpl, nil
}

// UpdateTemplate updates a template.
func (c *Client) UpdateTemplate(ctx context.Context, id int64, req TemplateRequest) (*Template, error) {
	var tpl Template
	_, err := c.doJSON(ctx, http.MethodPatch, c.cfg.BaseURL+"/api/outbound/templates/"+itoa(id), req, &tpl)
	if err != nil {
		return nil, err
	}
	return &tpl, nil
}

// DeleteTemplate deletes a template.
func (c *Client) DeleteTemplate(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodDelete, c.cfg.BaseURL+"/api/outbound/templates/"+itoa(id), nil, nil)
	return err
}
