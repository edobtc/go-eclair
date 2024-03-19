package eclair

import (
	"encoding/json"
)

const (
	auditPath = "/audit"
)

type AuditResponse struct {
	Sent     []PaymentSent     `json:"sent"`
	Received []PaymentReceived `json:"received"`
	Relayed  []PaymentRelayed  `json:"relayed"`
}

func (c *Client) Audit() (*AuditResponse, error) {
	var response AuditResponse
	data, err := c.Post(auditPath, nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
