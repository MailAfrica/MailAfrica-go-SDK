package mailafrica

import (
	"context"
	"net/http"
)

// ListAgentConfigs returns all agent configurations for the user.
func (c *Client) ListAgentConfigs(ctx context.Context) ([]AgentConfig, error) {
	var configs []AgentConfig
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/agent/configs", nil, &configs)
	if err != nil {
		return nil, err
	}
	return configs, nil
}

// GetAgentConfig retrieves the agent config for a specific address.
func (c *Client) GetAgentConfig(ctx context.Context, addressID int64) (*AgentConfig, error) {
	var config AgentConfig
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/agent/configs/"+itoa(addressID), nil, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// UpdateAgentConfig updates the agent config for a specific address.
func (c *Client) UpdateAgentConfig(ctx context.Context, addressID int64, req UpdateAgentConfigRequest) (*AgentConfig, error) {
	var config AgentConfig
	_, err := c.doJSON(ctx, http.MethodPut, c.cfg.BaseURL+"/api/agent/configs/"+itoa(addressID), req, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// GenerateAgentDraft generates an AI reply draft for a specific address.
func (c *Client) GenerateAgentDraft(ctx context.Context, addressID int64, req AgentDraftRequest) (*AgentDraft, error) {
	var draft AgentDraft
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/agent/configs/"+itoa(addressID)+"/draft", req, &draft)
	if err != nil {
		return nil, err
	}
	return &draft, nil
}
