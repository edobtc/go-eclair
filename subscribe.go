package eclair

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

const (
	recoveryInterval = 5 * time.Second
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

		for {
			conn, _, err := websocket.DefaultDialer.Dial(
				c.websocketURL(),
				c.settings.AuthHeaders(),
			)

			if err != nil {
				err = fmt.Errorf("error connecting to WebSocket: %v", err)
				fmt.Println(err)
				time.Sleep(recoveryInterval)
				continue
			}

			func() {
				defer conn.Close()

				for {
					_, msg, err := conn.ReadMessage()
					if err != nil {
						if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
							fmt.Printf("error reading from WebSocket: %v\n", err)
						}
						return // Exit the inner function to reconnect
					}

					event, err := UnmarshalEvent(msg)
					if err != nil {
						fmt.Printf("error unmarshaling event: %v\n", err)
						continue
					}

					ch <- *event
				}
			}()
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
