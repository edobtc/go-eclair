package eclair

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

func (c *Client) websocketURL() string {
	u, _ := url.Parse(c.BaseURL)
	u.Scheme = "ws"
	u.Path = "/ws"
	return u.String()
}

func (c *Client) Subscribe() (<-chan (Message), error) {
	ch := make(chan Message)

	go func() {
		defer close(ch)

		conn, _, err := websocket.DefaultDialer.Dial(
			c.websocketURL(),
			c.settings.AuthHeaders(),
		)

		if err != nil {
			err = fmt.Errorf("error connecting to WebSocket: %v", err)
			fmt.Println(err)
			return
		}
		defer conn.Close()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				err = fmt.Errorf("error reading from WebSocket: %v", err)
				fmt.Println(err)
				return
			}

			event, err := UnmarshalEvent(msg)
			if err != nil {
				err = fmt.Errorf("error unmarshaling event: %v", err)
				fmt.Println(err)
				return
			}

			ch <- *event
		}
	}()

	return ch, nil
}

// FilteredSubscription allow you to subscribe to a specific set of events
// and only receive messages for those, vs a firehose of all of them on the
// websocket
// This is a convenience method that wraps Subscribe and filters the messages
// based on the event type
func (c *Client) FilteredSubscription(evts ...eventType) (<-chan (Message), error) {
	ch, err := c.Subscribe()
	if err != nil {
		return nil, err
	}

	filteredCh := make(chan Message)

	go func() {
		defer close(filteredCh)

		for msg := range ch {
			for _, evt := range evts {
				if msg.Type == evt {
					filteredCh <- msg
				}
			}
		}
	}()

	return filteredCh, nil
}
