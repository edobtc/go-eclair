package eclair

import (
	"encoding/json"
	"fmt"
)

type eventType int

const (
	PaymentReceived eventType = iota
	PaymentRelayed
	PaymentSent
	PaymentSettlingOnchain
	PaymentFailed
	ChannelCreated
	ChannelOpened
	ChannelStateChanged
	ChannelClosed
	OnionMessageReceived
)

func ToString(kind eventType) string {
	switch kind {
	case PaymentReceived:
		return "payment-received"
	case PaymentRelayed:
		return "payment-relayed"
	case PaymentSent:
		return "payment-sent"
	case PaymentSettlingOnchain:
		return "payment-settling-onchain"
	case PaymentFailed:
		return "payment-failed"
	case ChannelCreated:
		return "channel-created"
	case ChannelOpened:
		return "channel-opened"
	case ChannelStateChanged:
		return "channel-state-changed"
	case ChannelClosed:
		return "channel-closed"
	case OnionMessageReceived:
		return "onion-message-received"
	default:
		return "unknown"
	}
}

type Message struct {
	Type string `json:"type"`
	Data interface{}
}

func ToEventType(eventType string) eventType {
	switch eventType {
	case "payment-received":
		return PaymentReceived
	case "payment-relayed":
		return PaymentRelayed
	case "payment-sent":
		return PaymentSent
	case "payment-settling-onchain":
		return PaymentSettlingOnchain
	case "payment-failed":
		return PaymentFailed
	case "channel-created":
		return ChannelCreated
	case "channel-opened":
		return ChannelOpened
	case "channel-state-changed":
		return ChannelStateChanged
	case "channel-closed":
		return ChannelClosed
	case "onion-message-received":
		return OnionMessageReceived
	default:
		return -1
	}
}

// BaseEvent is used to unmarshal the initial type information
type BaseEvent struct {
	Type string `json:"type"`
}

// UnmarshalEvent takes a byte slice and returns the correct event struct
func UnmarshalEvent(data []byte) (*Message, error) {
	var baseEvent BaseEvent

	if err := json.Unmarshal(data, &baseEvent); err != nil {
		return nil, err
	}

	switch baseEvent.Type {
	case "payment-received":
		var event PaymentReceivedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: baseEvent.Type,
			Data: event,
		}, nil
	case "payment-relayed":
		var event PaymentRelayedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: baseEvent.Type,
			Data: event,
		}, nil

	case "payment-sent":
		var event PaymentSentEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: baseEvent.Type,
			Data: event,
		}, nil

	case "payment-settling-onchain":
		var event PaymentSettlingOnchainEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: baseEvent.Type,
			Data: event,
		}, nil

	case "payment-failed":
		var event PaymentFailedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: baseEvent.Type,
			Data: event,
		}, nil

	case "channel-created":
		var event ChannelCreatedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: baseEvent.Type,
			Data: event,
		}, nil

	case "channel-opened":
		var event ChannelOpenedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: baseEvent.Type,
			Data: event,
		}, nil

	case "channel-state-changed":
		var event ChannelStateChangedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: baseEvent.Type,
			Data: event,
		}, nil

	case "channel-closed":
		var event ChannelClosedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: baseEvent.Type,
			Data: event,
		}, nil

	case "onion-message-received":
		var event OnionMessageReceivedEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: baseEvent.Type,
			Data: event,
		}, nil

	default:
		return nil, fmt.Errorf("unknown event type: %s", baseEvent.Type)
	}
}

type PaymentRelayedEvent struct {
	Type          string    `json:"type"`
	AmountIn      int       `json:"amountIn"`
	AmountOut     int       `json:"amountOut"`
	PaymentHash   string    `json:"paymentHash"`
	FromChannelId string    `json:"fromChannelId"`
	ToChannelId   string    `json:"toChannelId"`
	Timestamp     Timestamp `json:"timestamp"`
}

type Timestamp struct {
	Iso  string `json:"iso"`
	Unix int64  `json:"unix"`
}

type PaymentReceivedEvent struct {
	Type        string        `json:"type"`
	PaymentHash string        `json:"paymentHash"`
	Parts       []PaymentPart `json:"parts"`
}

type PaymentFailedEvent struct {
	Type        string    `json:"type"`
	ID          string    `json:"id"`
	PaymentHash string    `json:"paymentHash"`
	Failures    []Failure `json:"failures"`
	Timestamp   Timestamp `json:"timestamp"`
}

type Failure struct {
	FailureType    string   `json:"failureType"`
	FailureMessage string   `json:"failureMessage"`
	FailedRoute    []string `json:"failedRoute"`
}

type PaymentSentEvent struct {
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

type PaymentSettlingOnchainEvent struct {
	Type        string    `json:"type"`
	ID          string    `json:"id"`
	Amount      int       `json:"amount"`
	PaymentHash string    `json:"paymentHash"`
	Timestamp   Timestamp `json:"timestamp"`
}

type ChannelCreatedEvent struct {
	Type                  string `json:"type"`
	RemoteNodeId          string `json:"remoteNodeId"`
	IsInitiator           bool   `json:"isInitiator"`
	TemporaryChannelId    string `json:"temporaryChannelId"`
	InitialFeeRatePerKw   int    `json:"initialFeeratePerKw"`
	FundingTxFeeRatePerKw int    `json:"fundingTxFeeratePerKw"`
}

type ChannelOpenedEvent struct {
	Type         string `json:"type"`
	RemoteNodeId string `json:"remoteNodeId"`
	ChannelId    string `json:"channelId"`
}

type ChannelStateChangedEvent struct {
	Type          string `json:"type"`
	ChannelId     string `json:"channelId"`
	RemoteNodeId  string `json:"remoteNodeId"`
	PreviousState string `json:"previousState"`
	CurrentState  string `json:"currentState"`
}

type ChannelClosedEvent struct {
	Type        string `json:"type"`
	ChannelId   string `json:"channelId"`
	ClosingType string `json:"closingType"`
}

type OnionMessageReceivedEvent struct {
	Type        string            `json:"type"`
	PathId      string            `json:"pathId"`
	UnknownTlVS map[string]string `json:"unknownTlvs"`
}
