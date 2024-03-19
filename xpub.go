package eclair

import "encoding/json"

const (
	getMasterXPubPath = "/getmasterxpub"
)

type GetMasterXPubResponse struct {
	Xpub string `json:"xpub"`
}

// GetMasterXPub returns the master extended public key for the node.
func (c *Client) GetMasterXPub() (*GetMasterXPubResponse, error) {
	var xpub GetMasterXPubResponse
	data, err := c.Post(getMasterXPubPath, nil, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &xpub)
	if err != nil {
		return nil, err
	}

	return &xpub, nil
}
