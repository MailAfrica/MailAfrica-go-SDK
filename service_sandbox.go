package mailafrica

import (
	"context"
	"net/http"
)

// CreateSandboxCredential creates a new sandbox credential.
func (c *Client) CreateSandboxCredential(ctx context.Context, req CreateCredentialRequest) (*Credential, error) {
	var cred Credential
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/sandbox/credentials", req, &cred)
	if err != nil {
		return nil, err
	}
	return &cred, nil
}

// ListSandboxCredentials returns all sandbox credentials.
func (c *Client) ListSandboxCredentials(ctx context.Context) ([]Credential, error) {
	var creds []Credential
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/sandbox/credentials", nil, &creds)
	if err != nil {
		return nil, err
	}
	return creds, nil
}

// RevokeSandboxCredential revokes a sandbox credential.
func (c *Client) RevokeSandboxCredential(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/sandbox/credentials/"+itoa(id)+"/revoke", nil, nil)
	return err
}

// GetSMTPSandboxCredentials retrieves SMTP credentials for the sandbox.
func (c *Client) GetSMTPSandboxCredentials(ctx context.Context) (*SMTPSandboxCredentials, error) {
	var creds SMTPSandboxCredentials
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/sandbox/credentials/smtp", nil, &creds)
	if err != nil {
		return nil, err
	}
	return &creds, nil
}

// RegenerateSMTPSandboxPassword regenerates the sandbox SMTP password.
func (c *Client) RegenerateSMTPSandboxPassword(ctx context.Context) (*SMTPSandboxCredentials, error) {
	var creds SMTPSandboxCredentials
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/sandbox/credentials/smtp/regenerate", nil, &creds)
	if err != nil {
		return nil, err
	}
	return &creds, nil
}

// ListSandboxMessages returns sandbox messages with pagination.
func (c *Client) ListSandboxMessages(ctx context.Context, opts ListOpts) ([]SandboxMessage, *Pagination, error) {
	opts.applyDefaults()
	url := c.cfg.BaseURL + "/api/sandbox/messages?page=" + itoa(int64(opts.Page)) + "&per_page=" + itoa(int64(opts.PerPage))

	var resp struct {
		Success    bool        `json:"success"`
		Data       []SandboxMessage `json:"data"`
		Pagination *Pagination `json:"pagination"`
	}

	_, err := c.doJSON(ctx, http.MethodGet, url, nil, &resp)
	if err != nil {
		return nil, nil, err
	}
	return resp.Data, resp.Pagination, nil
}

// GetSandboxMessage retrieves a single sandbox message.
func (c *Client) GetSandboxMessage(ctx context.Context, id int64) (*SandboxMessage, error) {
	var msg SandboxMessage
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/sandbox/messages/"+itoa(id), nil, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

// ClearSandboxMessages deletes all sandbox messages.
func (c *Client) ClearSandboxMessages(ctx context.Context) error {
	_, err := c.doJSON(ctx, http.MethodDelete, c.cfg.BaseURL+"/api/sandbox/messages", nil, nil)
	return err
}
