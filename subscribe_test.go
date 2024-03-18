package eclair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_websocketURL(t *testing.T) {
	testCases := []struct {
		name     string
		baseURL  string
		expected string
	}{
		{
			name:     "Default BaseURL",
			baseURL:  "http://localhost:8080",
			expected: "ws://localhost:8080/ws",
		},
		{
			name:     "Custom BaseURL",
			baseURL:  "http://example.com",
			expected: "ws://example.com/ws",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := &Client{
				BaseURL: tc.baseURL,
			}
			actual := client.websocketURL()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
