package eclair

import "encoding/json"

const (
	channelStatsPath = "/channelstats"
)

type ChannelSummary struct {
	ChannelId        string `json:"channelId"`
	Direction        string `json:"direction"`
	AvgPaymentAmount int    `json:"avgPaymentAmount"`
	PaymentCount     int    `json:"paymentCount"`
	RelayFee         int    `json:"relayFee"`
	NetworkFee       int    `json:"networkFee"`
}

type ChannelStatsResponse []ChannelSummary

func (c *Client) ChannelStats() (*ChannelStatsResponse, error) {
	chStats := ChannelStatsResponse{}
	data, err := c.Post(channelStatsPath, nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &chStats)
	if err != nil {
		return nil, err
	}

	return &chStats, nil
}
