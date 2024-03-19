package eclair

type PaymentRelayed struct {
	Type          string    `json:"type"`
	AmountIn      int       `json:"amountIn"`
	AmountOut     int       `json:"amountOut"`
	PaymentHash   string    `json:"paymentHash"`
	FromChannelId string    `json:"fromChannelId"`
	ToChannelId   string    `json:"toChannelId"`
	Timestamp     Timestamp `json:"timestamp"`
}

type PaymentReceived struct {
	Type        string        `json:"type"`
	PaymentHash string        `json:"paymentHash"`
	Parts       []PaymentPart `json:"parts"`
}

type PaymentFailed struct {
	Type        string    `json:"type"`
	ID          string    `json:"id"`
	PaymentHash string    `json:"paymentHash"`
	Failures    []Failure `json:"failures"`
	Timestamp   Timestamp `json:"timestamp"`
}

type PaymentSent struct {
	Type            string        `json:"type"`
	ID              string        `json:"id"`
	PaymentHash     string        `json:"paymentHash"`
	PaymentPreimage string        `json:"paymentPreimage"`
	RecipientAmount int           `json:"recipientAmount"`
	RecipientNodeId string        `json:"recipientNodeId"`
	Parts           []PaymentPart `json:"parts"`
}

type PaymentPart struct {
	ID            string    `json:"id"`
	Amount        int       `json:"amount"`
	FeesPaid      int       `json:"feesPaid"`
	FromChannelId string    `json:"fromChannelId"`
	ToChannelId   string    `json:"toChannelId"`
	Timestamp     Timestamp `json:"timestamp"`
}

type PaymentSettlingOnchain struct {
	Type        string    `json:"type"`
	ID          string    `json:"id"`
	Amount      int       `json:"amount"`
	PaymentHash string    `json:"paymentHash"`
	Timestamp   Timestamp `json:"timestamp"`
}
