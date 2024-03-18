package eclair

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) Get(path string, response interface{}) ([]byte, error) {
	return c.Request(http.MethodGet, path, nil, response)
}

func (c *Client) Post(path string, body interface{}, response interface{}) ([]byte, error) {
	return c.Request(http.MethodPost, path, body, response)
}
func (c *Client) Request(method, path string, body interface{}, response interface{}) ([]byte, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	path = strings.TrimPrefix(path, "/")
	url := strings.TrimSuffix(c.BaseURL, "/") + "/" + path

	req, err := http.NewRequest(
		method,
		url,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if c.settings.Credentials != nil {
		req.SetBasicAuth(
			c.settings.Credentials.User,
			c.settings.Credentials.Password,
		)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		fmt.Println(string(data))
	}

	return data, nil
}
