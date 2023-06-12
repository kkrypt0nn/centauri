package centauri

import (
	"net/http"
	"time"

	"github.com/kkrypt0nn/centauri/ext/tasks"
	"github.com/kkrypt0nn/centauri/rest"
	"github.com/kkrypt0nn/logger.go"
)

// NewRestClient returns a new rest.Client to make REST API calls only, may be a bot or a user token
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
