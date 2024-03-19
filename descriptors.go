package eclair

import "encoding/json"

const (
	getDescriptorsPath = "/getdescriptors"
)

type Descriptor struct {
	Desc      string `json:"desc"`
	Internal  bool   `json:"internal"`
	Active    bool   `json:"active"`
	Timestamp int    `json:"timestamp"`
}

type GetDescriptorsResponse []Descriptor

func (c *Client) GetDescriptors() (*GetDescriptorsResponse, error) {
	var desc GetDescriptorsResponse

	data, err := c.Post(getDescriptorsPath, nil, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &desc)
	if err != nil {
		return nil, err
	}

	return &desc, nil
}
