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
