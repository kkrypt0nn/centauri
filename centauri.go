package centauri

import (
	"github.com/kkrypt0nn/centauri/api"
	"github.com/kkrypt0nn/logger.go"
	"net/http"
	"time"
)

const VERSION = "0.0.1"

// NewRestClient returns a new api.RestClient to make REST API calls only, may be a bot or a user token
func NewRestClient(token string, isBot bool) *api.RestClient {
	header := token
	if isBot {
		header = "Bot " + header
	}
	restClient := &api.RestClient{
		HttpClient: &http.Client{
			Timeout: 20 * time.Second,
		},
		Logger: logger.NewLogger(),
	}
	restClient.SetAuthorizationHeader(header)
	return restClient
}
