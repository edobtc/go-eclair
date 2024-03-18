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
		return
	}

	for message := range channel {
		fmt.Println(message)
	}
}
