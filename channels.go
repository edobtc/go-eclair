package eclair

type ChannelCreated struct {
	Type                  string `json:"type"`
	RemoteNodeId          string `json:"remoteNodeId"`
	IsInitiator           bool   `json:"isInitiator"`
	TemporaryChannelId    string `json:"temporaryChannelId"`
	InitialFeeRatePerKw   int    `json:"initialFeeratePerKw"`
	FundingTxFeeRatePerKw int    `json:"fundingTxFeeratePerKw"`
}

type ChannelOpened struct {
	Type         string `json:"type"`
	RemoteNodeId string `json:"remoteNodeId"`
	ChannelId    string `json:"channelId"`
}

type ChannelStateChanged struct {
	Type          string `json:"type"`
	ChannelId     string `json:"channelId"`
	RemoteNodeId  string `json:"remoteNodeId"`
	PreviousState string `json:"previousState"`
	CurrentState  string `json:"currentState"`
}

type ChannelClosed struct {
	Type        string `json:"type"`
	ChannelId   string `json:"channelId"`
	ClosingType string `json:"closingType"`
}
