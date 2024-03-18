package main

import (
	"fmt"

	eclair "github.com/edobtc/go-eclair"
)

func main() {
	client := eclair.NewClient()

	// polar node
	client = client.WithBaseURL("http://localhost:8282")

	channel, err := client.Subscribe()
	if err != nil {
		panic(err)
	}

	for message := range channel {
		fmt.Println(message)
		// handle messages here

		switch message.Type {
		case eclair.ChannelOpened:
			// handle channel opened event
		case eclair.ChannelClosed:
			// handle channel closed event
		case eclair.PaymentReceived:
			// handle channel closed event
		}
	}
}
