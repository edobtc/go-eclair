package eclair

import (
	"net/http"

	"github.com/edobtc/go-eclair/httpclient"
)

type Client struct {
	settings   *Settings
	BaseURL    string
	HTTPClient *http.Client

	Debug bool
}

func NewClient() *Client {
	return &Client{
		settings:   NewDefaultSettings(),
		HTTPClient: httpclient.New(),
	}
}

func NewClientFromSettings(settings *Settings) *Client {
	return &Client{
		settings:   settings,
		HTTPClient: httpclient.New(),
	}
}

func (c *Client) _settings() *Settings {
	if c.settings == nil {
		c.settings = NewDefaultSettings()
	}

	if c.settings.Credentials == nil {
		c.settings.Credentials = &Credentials{}
	}

	return c.settings
}

func (c *Client) WithBaseURL(baseURL string) *Client {
	c.BaseURL = baseURL
	return c
}

func (c *Client) WithCredentials(creds *Credentials) *Client {
	c._settings().Credentials = creds
	return c
}

func (c *Client) WithUser(user string) *Client {
	c._settings().Credentials.User = user
	return c
}

func (c *Client) WithPassword(password string) *Client {

	c.settings.Credentials.Password = password
	return c
}

func (c *Client) WithHTTPClient(client *http.Client) *Client {
	c.HTTPClient = client
	return c
}

func (c *Client) WithDebug(debug bool) *Client {
	c.Debug = debug
	return c
}
