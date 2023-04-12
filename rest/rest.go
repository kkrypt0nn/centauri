package rest

import (
	"github.com/kkrypt0nn/logger.go"
	"net/http"
)

const (
	BaseURL  = "https://discord.com/api/v"
	Version  = "10"
	Endpoint = BaseURL + Version + "/"
)

// Client is a client made to send REST API requests only
type Client struct {
	HttpClient *http.Client
	Logger     *logger.Logger
	Debug      bool

	authorizationHeader string
}

// SetAuthorizationHeader sets the authorization header of the RestClient
func (c *Client) SetAuthorizationHeader(authorizationHeader string) {
	c.authorizationHeader = authorizationHeader
}
