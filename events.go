package eclair

import (
	"encoding/json"
	"fmt"
)

type eventType int

const (
	PaymentReceivedEvent eventType = iota
	PaymentRelayedEvent
	PaymentSentEvent
	PaymentSettlingOnchainEvent
	PaymentFailedEvent
	ChannelCreatedEvent
	ChannelOpenedEvent
	ChannelStateChangedEvent
	ChannelClosedEvent
	OnionMessageReceivedEvent
)

func ToString(kind eventType) string {
	switch kind {
	case PaymentReceivedEvent:
		return "payment-received"
	case PaymentRelayedEvent:
		return "payment-relayed"
	case PaymentSentEvent:
		return "payment-sent"
	case PaymentSettlingOnchainEvent:
		return "payment-settling-onchain"
	case PaymentFailedEvent:
		return "payment-failed"
	case ChannelCreatedEvent:
		return "channel-created"
	case ChannelOpenedEvent:
		return "channel-opened"
	case ChannelStateChangedEvent:
		return "channel-state-changed"
	case ChannelClosedEvent:
		return "channel-closed"
	case OnionMessageReceivedEvent:
		return "onion-message-received"
	default:
		return "unknown"
	}
}

type Message struct {
	Type eventType `json:"type"`
	Data interface{}
}

func ToEventType(eventType string) eventType {
	switch eventType {
	case "payment-received":
		return PaymentReceivedEvent
	case "payment-relayed":
		return PaymentRelayedEvent
	case "payment-sent":
		return PaymentSentEvent
	case "payment-settling-onchain":
		return PaymentSettlingOnchainEvent
	case "payment-failed":
		return PaymentFailedEvent
	case "channel-created":
		return ChannelCreatedEvent
	case "channel-opened":
		return ChannelOpenedEvent
	case "channel-state-changed":
		return ChannelStateChangedEvent
	case "channel-closed":
		return ChannelClosedEvent
	case "onion-message-received":
		return OnionMessageReceivedEvent
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
		var event PaymentReceived
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: ToEventType(baseEvent.Type),
			Data: event,
		}, nil
	case "payment-relayed":
		var event PaymentRelayed
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: ToEventType(baseEvent.Type),
			Data: event,
		}, nil

	case "payment-sent":
		var event PaymentSent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: ToEventType(baseEvent.Type),
			Data: event,
		}, nil

	case "payment-settling-onchain":
		var event PaymentSettlingOnchain
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: ToEventType(baseEvent.Type),
			Data: event,
		}, nil

	case "payment-failed":
		var event PaymentFailed
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: ToEventType(baseEvent.Type),
			Data: event,
		}, nil

	case "channel-created":
		var event ChannelCreated
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: ToEventType(baseEvent.Type),
			Data: event,
		}, nil

	case "channel-opened":
		var event ChannelOpened
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: ToEventType(baseEvent.Type),
			Data: event,
		}, nil

	case "channel-state-changed":
		var event ChannelStateChanged
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: ToEventType(baseEvent.Type),
			Data: event,
		}, nil

	case "channel-closed":
		var event ChannelClosed
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: ToEventType(baseEvent.Type),
			Data: event,
		}, nil

	case "onion-message-received":
		var event OnionMessageReceived
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return &Message{
			Type: ToEventType(baseEvent.Type),
			Data: event,
		}, nil

	default:
		return nil, fmt.Errorf("unknown event type: %s", baseEvent.Type)
	}
}
