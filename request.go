package eclair

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) Get(path string, response interface{}) ([]byte, error) {
	return c.Request(http.MethodGet, path, nil, response)
}

func (c *Client) Post(path string, data io.Reader, response interface{}) ([]byte, error) {
	return c.Request(http.MethodPost, path, data, response)
}
func (c *Client) Request(method, path string, data io.Reader, response interface{}) ([]byte, error) {
	path = strings.TrimPrefix(path, "/")
	url := strings.TrimSuffix(c.BaseURL, "/") + "/" + path

	req, err := http.NewRequest(
		method,
		url,
		data,
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

	d, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(d))
	}

	return d, nil
}
