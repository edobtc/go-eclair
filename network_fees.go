package eclair

import (
	"encoding/json"
)

const (
	networkFeesPath = "/networkfees"
)

type NetworkFeesResponse struct {
	RemoteNodeId string    `json:"remoteNodeId"`
	ChannelId    string    `json:"channelId"`
	TxId         string    `json:"txId"`
	Fee          int       `json:"fee"`
	TxType       string    `json:"txType"`
	Timestamp    Timestamp `json:"timestamp"`
}

func (c *Client) NetworkFees() (*NetworkFeesResponse, error) {
	networkFees := NetworkFeesResponse{}
	data, err := c.Post(networkFeesPath, nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &networkFees)
	if err != nil {
		return nil, err
	}

	return &networkFees, nil
}
