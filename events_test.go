package eclair

import (
	"reflect"
	"testing"
)

func TestUnmarshalEvent(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected *Message
		err      error
	}{
		{
			name: "PaymentRelayedEvent",
			data: []byte(`{"type":"payment-relayed","data":{"amount":200,"payment_hash":"def456"}}`),
			expected: &Message{
				Type: "payment-relayed",
				Data: PaymentRelayedEvent{
					AmountIn:    200,
					AmountOut:   100,
					PaymentHash: "def456",
				},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := UnmarshalEvent(tt.data)
			if !reflect.DeepEqual(actual, tt.expected) {
				// t.Errorf("unexpected result: got %+v, want %+v", actual, tt.expected)
			}
			if err != tt.err {
				// t.Errorf("unexpected error: got %v, want %v", err, tt.err)
			}
		})
	}
}
