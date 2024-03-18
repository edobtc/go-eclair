package eclair

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

const (
	createPath = "/createinvoice"
	deletePath = "/deleteinvoice"
	payInvoice = "/payinvoice"
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
	Description     string `json:"description,omitempty"`
	DescriptionHash string `json:"descriptionHash,omitempty"`
	Expiry          int    `json:"expireIn,omitempty"`
	Amount          int    `json:"amountMsat"`
	FallbackAddress string `json:"fallbackAddress,omitempty"`
	PaymentPreimage string `json:"paymentPreimage,omitempty"`
}

func (c *CreateInvoiceRequest) FormData() (io.Reader, error) {
	formData := url.Values{}

	if c.Description != "" {
		formData.Set("description", c.Description)
	}

	if c.DescriptionHash != "" {
		formData.Set("descriptionHash", c.DescriptionHash)
	}
	if c.Expiry != 0 {
		formData.Set("expireIn", strconv.Itoa(c.Expiry))
	}
	if c.FallbackAddress != "" {
		formData.Set("fallbackAddress", c.FallbackAddress)
	}
	if c.PaymentPreimage != "" {
		formData.Set("paymentPreimage", c.PaymentPreimage)
	}

	if c.Amount != 0 {
		formData.Set("amountMsat", fmt.Sprintf("%d", c.Amount))
	}

	return strings.NewReader(formData.Encode()), nil
}

type InvoiceFeatures struct {
	Activated map[string]string `json:"activated"`
	Unknown   []interface{}     `json:"unknown"`
}

func (c *Client) CreateInvoice(settings CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	resp := CreateInvoiceResponse{}
	data, err := settings.FormData()
	if err != nil {
		return nil, err
	}

	r, err := c.Post(createPath, data, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(r, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

type DeleteInvoiceResponse struct {
	Message string `json:"message"`
}

type DeleteInvoiceRequest struct {
	PaymentHash string `json:"paymentHash"`
}

func (c *DeleteInvoiceRequest) FormData() (io.Reader, error) {
	formData := url.Values{}
	formData.Set("paymentHash", c.PaymentHash)
	return strings.NewReader(formData.Encode()), nil
}

func (c *Client) DeleteInvoice(settings DeleteInvoiceRequest) (*DeleteInvoiceResponse, error) {

	data, err := settings.FormData()
	if err != nil {
		return nil, err
	}

	r, err := c.Post(deletePath, data, nil)
	if err != nil {
		return nil, err
	}

	return &DeleteInvoiceResponse{
		Message: string(r),
	}, nil
}

type PayInvoiceRequest struct {
	Invoice                   string `json:"invoice,omitempty"`
	AmountMsat                int    `json:"amountMsat"`
	MaxAttempts               int    `json:"maxAttempts"`
	MaxFeeFlatSat             int    `json:"maxFeeFlatSat"`
	MaxFeePct                 int    `json:"maxFeePct"`
	ExternalId                string `json:"externalId"`
	PathFindingExperimentName string `json:"pathFindingExperimentName"`
	Blocking                  bool   `json:"blocking"`
}

type PayInvoiceResponse struct {
	Message string `json:"message"`
}

func (c *PayInvoiceRequest) FormData() (io.Reader, error) {
	formData := url.Values{}

	if c.Invoice != "" {
		formData.Set("invoice", c.Invoice)
	}
	if c.AmountMsat != 0 {
		formData.Set("amountMsat", strconv.Itoa(c.AmountMsat))
	}
	if c.MaxAttempts != 0 {
		formData.Set("maxAttempts", strconv.Itoa(c.MaxAttempts))
	}
	if c.MaxFeeFlatSat != 0 {
		formData.Set("maxFeeFlatSat", strconv.Itoa(c.MaxFeeFlatSat))
	}
	if c.MaxFeePct != 0 {
		formData.Set("maxFeePct", strconv.Itoa(c.MaxFeePct))
	}
	if c.ExternalId != "" {
		formData.Set("externalId", c.ExternalId)
	}
	if c.PathFindingExperimentName != "" {
		formData.Set("pathFindingExperimentName", c.PathFindingExperimentName)
	}
	formData.Set("blocking", strconv.FormatBool(c.Blocking))

	return strings.NewReader(formData.Encode()), nil
}

func (c *Client) PayInvoice(settings PayInvoiceRequest) (*PayInvoiceResponse, error) {
	data, err := settings.FormData()
	if err != nil {
		return nil, err
	}

	r, err := c.Post(payInvoice, data, nil)
	if err != nil {
		return nil, err
	}

	return &PayInvoiceResponse{
		Message: string(r),
	}, nil
}
