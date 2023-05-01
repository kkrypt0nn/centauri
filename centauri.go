package centauri

import (
	"github.com/kkrypt0nn/centauri/ext/tasks"
	"github.com/kkrypt0nn/centauri/rest"
	"github.com/kkrypt0nn/logger.go"
	"net/http"
	"time"
)

// NewRestClient returns a new rest.Client to make REST API calls only, may be a bot or a user token
func NewRestClient(token string) *rest.Client {
	restClient := &rest.Client{
		HttpClient: &http.Client{
			Timeout: 20 * time.Second,
		},
		Logger:      logger.NewLogger(),
		RateLimiter: rest.NewRateLimiter(),
		TaskManager: tasks.NewTaskManager(),
	}
	restClient.SetAuthorizationHeader(token)
	return restClient
}
