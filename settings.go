package eclair

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	DefaultBaseURL  = "http://localhost:8080"
	DefaultUser     = "eclair"
	DefaultPassword = "eclairpw"
)

type Settings struct {
	BaseURL     string
	Credentials *Credentials
}

func NewDefaultSettings() *Settings {
	return &Settings{
		BaseURL: DefaultBaseURL,
		Credentials: &Credentials{
			User:     DefaultUser,
			Password: DefaultPassword,
		},
	}
}

func (s *Settings) AuthHeaders() http.Header {
	return s.Credentials.HTTPAuthHeaders()
}

// Credentials
type Credentials struct {
	User     string
	Password string
}

func (c *Credentials) BasicAuth() string {
	auth := fmt.Sprintf("%s:%s", c.User, c.Password)
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	return encodedAuth
}

func (c *Credentials) HTTPAuthHeaders() http.Header {
	headers := http.Header{}
	headers.Set("Authorization", "Basic "+c.BasicAuth())
	return headers
}
