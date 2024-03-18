package eclair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	assert.NotNil(t, client)
	assert.NotNil(t, client.settings)
	assert.NotNil(t, client.HTTPClient)
}

func TestNewClientFromSettings(t *testing.T) {
	settings := NewDefaultSettings()
	client := NewClientFromSettings(settings)
	assert.NotNil(t, client)
	assert.Equal(t, settings, client.settings)
	assert.NotNil(t, client.HTTPClient)
}

func TestClient_settings(t *testing.T) {
	client := NewClient()
	assert.NotNil(t, client._settings())
	assert.Equal(t, client.settings, client._settings())
}

func TestClient_WithBaseURL(t *testing.T) {
	client := NewClient()
	baseURL := "http://localhost:8080"
	client.WithBaseURL(baseURL)
	assert.Equal(t, baseURL, client.BaseURL)
}

func TestClient_WithCredentials(t *testing.T) {
	client := NewClient()
	creds := &Credentials{
		User:     "user",
		Password: "password",
	}
	client.WithCredentials(creds)
	assert.Equal(t, creds, client.settings.Credentials)
}

func TestClient_WithUser(t *testing.T) {
	client := NewClient()
	user := "user"
	client.WithUser(user)
	assert.Equal(t, user, client.settings.Credentials.User)
}

func TestClient_WithPassword(t *testing.T) {
	client := NewClient()
	password := "password"
	client.WithPassword(password)
	assert.Equal(t, password, client.settings.Credentials.Password)
}
