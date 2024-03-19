package eclair

type OnionMessageReceived struct {
	Type        string            `json:"type"`
	PathId      string            `json:"pathId"`
	UnknownTlVS map[string]string `json:"unknownTlvs"`
}
