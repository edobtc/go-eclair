package main

import (
	"fmt"

	eclair "github.com/edobtc/go-eclair"
)

func main() {
	client := eclair.NewClient()

	// polar node
	client = client.WithBaseURL("http://localhost:8282")

	info, err := client.GetInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println(info)
}
