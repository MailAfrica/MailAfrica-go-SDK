package mailafrica

import (
	"context"
	"net/http"
)

// CreateWebhook creates a new webhook.
func (c *Client) CreateWebhook(ctx context.Context, req CreateWebhookRequest) (*Webhook, error) {
	var wh Webhook
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/webhook/webhooks", req, &wh)
	if err != nil {
		return nil, err
	}
	return &wh, nil
}

// ListWebhooks returns webhooks for an address.
func (c *Client) ListWebhooks(ctx context.Context, addressID int64) ([]Webhook, error) {
	var webhooks []Webhook
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/webhook/webhooks?address_id="+itoa(addressID), nil, &webhooks)
	if err != nil {
		return nil, err
	}
	return webhooks, nil
}

// DeleteWebhook deletes a webhook.
func (c *Client) DeleteWebhook(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodDelete, c.cfg.BaseURL+"/api/webhook/webhooks/"+itoa(id), nil, nil)
	return err
}

// ListWebhookDeliveries returns delivery attempts for a webhook.
func (c *Client) ListWebhookDeliveries(ctx context.Context, webhookID int64) ([]WebhookDelivery, error) {
	var deliveries []WebhookDelivery
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/webhook/webhooks/"+itoa(webhookID)+"/deliveries", nil, &deliveries)
	if err != nil {
		return nil, err
	}
	return deliveries, nil
}

// TestWebhook sends a test ping to a webhook.
func (c *Client) TestWebhook(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/webhook/webhooks/"+itoa(id)+"/test", nil, nil)
	return err
}

// TriggerWebhook manually triggers a webhook.
func (c *Client) TriggerWebhook(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/webhook/webhooks/trigger/"+itoa(id), nil, nil)
	return err
}
