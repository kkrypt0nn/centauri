package api

import (
	"github.com/kkrypt0nn/logger.go"
	"net/http"
)

const (
	BaseURL      = "https://discord.com"
	Version      = "v10"
	Path         = "/api/" + Version
	RestEndpoint = BaseURL + Path + "/"
)

// RestClient is a client made to send REST API requests only
type RestClient struct {
	HttpClient *http.Client
	Logger     *logger.Logger
	Debug      bool

	authorizationHeader string
}

// SetAuthorizationHeader sets the authorization header of the RestClient
func (c *RestClient) SetAuthorizationHeader(authorizationHeader string) {
	c.authorizationHeader = authorizationHeader
}
