package mailafrica

import (
	"context"
	"net/http"
)

// CreateSMSNotification creates an SMS notification rule.
func (c *Client) CreateSMSNotification(ctx context.Context, req CreateSMSNotificationRequest) (*SMSNotificationWithKey, error) {
	var resp SMSNotificationWithKey
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/sms/notifications", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSMSNotifications returns SMS notifications for an address.
func (c *Client) ListSMSNotifications(ctx context.Context, addressID int64) ([]SMSNotification, error) {
	var notifs []SMSNotification
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/sms/notifications?address_id="+itoa(addressID), nil, &notifs)
	if err != nil {
		return nil, err
	}
	return notifs, nil
}

// RevokeSMSNotification revokes an SMS notification rule.
func (c *Client) RevokeSMSNotification(ctx context.Context, id int64) error {
	_, err := c.doJSON(ctx, http.MethodPost, c.cfg.BaseURL+"/api/sms/notifications/"+itoa(id)+"/revoke", nil, nil)
	return err
}

// ListSMSDeliveries returns delivery attempts for an SMS notification.
func (c *Client) ListSMSDeliveries(ctx context.Context, id int64) ([]SMSDelivery, error) {
	var deliveries []SMSDelivery
	_, err := c.doJSON(ctx, http.MethodGet, c.cfg.BaseURL+"/api/sms/notifications/"+itoa(id)+"/deliveries", nil, &deliveries)
	if err != nil {
		return nil, err
	}
	return deliveries, nil
}
