package main

import (
	"fmt"

	eclair "github.com/edobtc/go-eclair"
)

func main() {
	client := eclair.NewClient()
	client = client.WithBaseURL("http://localhost:8282")
	client.Debug = true

	resp, err := client.CreateInvoice(eclair.CreateInvoiceRequest{
		Description: "test4",
		Amount:      10000,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	dresp, err := client.DeleteInvoice(eclair.DeleteInvoiceRequest{
		PaymentHash: resp.PaymentHash,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dresp.Message)

	// Create Invoice

	resp, err = client.CreateInvoice(eclair.CreateInvoiceRequest{
		Description: "test6",
		Amount:      169420,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	presp, err := client.PayInvoice(eclair.PayInvoiceRequest{
		Invoice: resp.Serialized,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(presp.Message)

}
