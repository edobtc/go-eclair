package eclair

import "encoding/json"

const (
	peersPath = "/peers"
)

type Peer struct {
	NodeId   string `json:"nodeId"`
	State    string `json:"state"`
	Address  string `json:"address"`
	Channels int    `json:"channels"`
}

type PeersResponse []Peer

func (c *Client) Peers() (*PeersResponse, error) {
	var peers PeersResponse
	data, err := c.Post(peersPath, nil, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &peers)
	if err != nil {
		return nil, err
	}

	return &peers, nil
}
