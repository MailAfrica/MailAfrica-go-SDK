package mailafrica

import (
	"context"
	"net/http"
)

// CreateAPIKey creates a new API key.
func (c *Client) CreateAPIKey(ctx context.Context, req CreateAPIKeyRequest) (*CreateAPIKeyResponse, error) {
	var resp CreateAPIKeyResponse
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/apikeys", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListAPIKeys returns all API keys for the user.
func (c *Client) ListAPIKeys(ctx context.Context) ([]APIKey, error) {
	var keys []APIKey
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/apikeys", nil, &keys)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

// RevokeAPIKey revokes an API key.
func (c *Client) RevokeAPIKey(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodDelete, c.cfg.BaseURL+"/api/apikeys/"+itoa(id), nil, nil)
	return err
}
