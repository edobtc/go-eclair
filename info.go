package eclair

import "encoding/json"

const (
	infoPath = "/getinfo"
)

type Info struct {
	Version         string   `json:"version"`
	NodeId          string   `json:"nodeId"`
	Alias           string   `json:"alias"`
	Color           string   `json:"color"`
	Features        Features `json:"features"`
	ChainHash       string   `json:"chainHash"`
	Network         string   `json:"network"`
	BlockHeight     int      `json:"blockHeight"`
	PublicAddresses []string `json:"publicAddresses"`
	InstanceId      string   `json:"instanceId"`
}

type Features struct {
	Activated map[string]string `json:"activated"`
	Unknown   []interface{}     `json:"unknown"`
}

func (c *Client) GetInfo() (*Info, error) {
	info := Info{}
	data, err := c.Post(infoPath, nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
