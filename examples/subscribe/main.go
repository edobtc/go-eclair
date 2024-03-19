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
		case eclair.ChannelOpenedEvent:
			// handle channel opened event
		case eclair.ChannelClosedEvent:
			// handle channel closed event
		case eclair.PaymentReceivedEvent:
			// handle channel closed event
		}
	}
}
