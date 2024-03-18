package eclair

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
)

const (
	createPath = "/createinvoice"
)

type CreateInvoiceResponse struct {
	Prefix             string          `json:"prefix"`
	Timestamp          int64           `json:"timestamp"`
	NodeId             string          `json:"nodeId"`
	Serialized         string          `json:"serialized"`
	Description        string          `json:"description"`
	PaymentHash        string          `json:"paymentHash"`
	PaymentMetadata    string          `json:"paymentMetadata"`
	Expiry             int             `json:"expiry"`
	MinFinalCltvExpiry int             `json:"minFinalCltvExpiry"`
	Amount             int             `json:"amount"`
	Features           InvoiceFeatures `json:"features"`
	RoutingInfo        []interface{}   `json:"routingInfo"`
}

type CreateInvoiceRequest struct {
	Description     string `json:"description"`
	DescriptionHash string `json:"descriptionHash"`
	Expiry          int    `json:"expireIn"`
	Amount          int    `json:"amountMsat"`
	FallbackAddress string `json:"fallbackAddress"`
	PaymentPreimage string `json:"paymentPreimage"`
}

func (c *CreateInvoiceRequest) ToBuffer() (io.ReadWriter, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	err := w.WriteField("description", c.Description)
	if err != nil {
		return nil, err
	}

	err = w.WriteField("amountMsat", fmt.Sprintf("%d", c.Amount))
	if err != nil {
		return nil, err
	}

	w.Close()
	return &b, nil
}

type InvoiceFeatures struct {
	Activated map[string]string `json:"activated"`
	Unknown   []interface{}     `json:"unknown"`
}

func (c *Client) CreateInvoice(settings CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	resp := CreateInvoiceResponse{}
	body, err := settings.ToBuffer()
	if err != nil {
		return nil, err
	}

	data, err := c.Post(createPath, &body, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
