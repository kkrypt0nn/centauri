package centauri

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/gateway"
	"net/http"
	"time"

	"github.com/kkrypt0nn/centauri/ext/tasks"
	"github.com/kkrypt0nn/centauri/rest"
	"github.com/kkrypt0nn/logger.go"
)

// NewRestClient returns a new REST client (rest.Client) to make REST API calls only, may be a bot or a user token
func NewRestClient(token string) *rest.Client {
	restClient := &rest.Client{
		HttpClient: &rest.HttpClient{
			Client: &http.Client{
				Timeout: 20 * time.Second,
			},
			Interceptor: nil,
		},
		Logger:      logger.NewLogger(),
		RateLimiter: rest.NewRateLimiter(),
		TaskManager: tasks.NewTaskManager(),
	}
	restClient.SetAuthorizationHeader(token)
	return restClient
}

// NewGatewayClient returns a new Gateway API client (gateway.Client) to handle real-time events from Discord
func NewGatewayClient(token string, intents discord.Intents) *gateway.Client {
	return gateway.New(token, intents)
}
