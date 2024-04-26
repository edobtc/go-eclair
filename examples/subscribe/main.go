package main

import (
	"encoding/json"
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
		data, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}

		// output nicely for demostation purposes
		fmt.Println(string(data))

		switch message.Type {
		case eclair.ChannelOpenedEvent:
			fmt.Println("channel opened")
			// handle channel opened event
		case eclair.ChannelClosedEvent:
			// handle channel closed event
			fmt.Println("channel closed")
		case eclair.PaymentReceivedEvent:
			// handle channel closed even
			fmt.Println("payment received")
		}
	}
}
