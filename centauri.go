package centauri

import (
	"github.com/kkrypt0nn/centauri/rest"
	"github.com/kkrypt0nn/logger.go"
	"net/http"
	"time"
)

// NewRestClient returns a new rest.Client to make REST API calls only, may be a bot or a user token
func NewRestClient(token string, isBot bool) *rest.Client {
	header := token
	if isBot {
		header = "Bot " + header
	}
	restClient := &rest.Client{
		HttpClient: &http.Client{
			Timeout: 20 * time.Second,
		},
		Logger:      logger.NewLogger(),
		RateLimiter: rest.NewRateLimiter(),
	}
	restClient.SetAuthorizationHeader(header)
	return restClient
}
