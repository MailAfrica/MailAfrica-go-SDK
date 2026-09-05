package mailafrica

import (
	"context"
	"net/http"
)

// GetComplianceProfile retrieves the user's compliance profile.
func (c *Client) GetComplianceProfile(ctx context.Context) (*ComplianceProfile, error) {
	var profile ComplianceProfile
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/compliance/profile", nil, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// UpdateComplianceProfile updates the user's compliance profile.
func (c *Client) UpdateComplianceProfile(ctx context.Context, req UpdateComplianceProfileRequest) (*ComplianceProfile, error) {
	var profile ComplianceProfile
	_, err := c.doJSON(ctx, http.MethodPatch, c.cfg.BaseURL+"/api/compliance/profile", req, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// GetAuditExport retrieves the compliance audit export.
func (c *Client) GetAuditExport(ctx context.Context) (*AuditExport, error) {
	var export AuditExport
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/compliance/audit-export", nil, &export)
	if err != nil {
		return nil, err
	}
	return &export, nil
}
