package eclair

type ChainState struct {
	Confirmed   int `json:"confirmed"`
	Unconfirmed int `json:"unconfirmed"`
}
