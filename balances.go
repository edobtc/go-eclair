package eclair

import "encoding/json"

const (
	usableBalancesPath  = "/usablebalances"
	channelBalancesPath = "/channelbalances"
	globalBalancePath   = "/globalbalance"
)

type UsableBalance struct {
	ChannelId string `json:"channelId"`
	ToRemote  int    `json:"toRemote"`
	ToLocal   int    `json:"toLocal"`
}

type UsableBalancesResponse []UsableBalance

type ChannelBalance struct {
	ChannelId string `json:"channelId"`
	Balance   int    `json:"balance"`
	Capacity  int    `json:"capacity"`
}

type ChannelBalancesResponse []ChannelBalance

type GlobalBalanceResponse struct {
	TotalBalance int        `json:"totalBalance"`
	Onchain      ChainState `json:"onchain"`
	Offchain     ChainState `json:"offchain"`
}

func (c *Client) UsableBalances() (*UsableBalancesResponse, error) {
	usableBalances := UsableBalancesResponse{}
	data, err := c.Post(usableBalancesPath, nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &usableBalances)
	if err != nil {
		return nil, err
	}

	return &usableBalances, nil
}

func (c *Client) ChannelBalances() (*ChannelBalancesResponse, error) {
	channelBalances := ChannelBalancesResponse{}
	data, err := c.Post(channelBalancesPath, nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &channelBalances)
	if err != nil {
		return nil, err
	}

	return &channelBalances, nil
}

func (c *Client) GlobalBalance() (*GlobalBalanceResponse, error) {
	globalBalance := GlobalBalanceResponse{}
	data, err := c.Post(globalBalancePath, nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &globalBalance)
	if err != nil {
		return nil, err
	}

	return &globalBalance, nil
}
