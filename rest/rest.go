package rest

import (
	"encoding/json"
	"fmt"
	"github.com/kkrypt0nn/centauri/constants"
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/errors"
	"github.com/kkrypt0nn/logger.go"
	"io"
	"math"
	"net/http"
	"strings"
	"time"
)

const (
	BaseURL  = "https://discord.com/api/v"
	Version  = "10"
	Endpoint = BaseURL + Version + "/"
)

// Client is a client made to send REST API requests only
type Client struct {
	HttpClient  *http.Client
	Logger      *logger.Logger
	Debug       bool
	RateLimiter *RateLimiter

	authorizationHeader string
}

// SetAuthorizationHeader sets the authorization header of the RestClient
func (c *Client) SetAuthorizationHeader(authorizationHeader string) {
	c.authorizationHeader = authorizationHeader
}

// DoRequest performs a request to the given URL with the given method, it will make sure to follow rate limits, and returns the body and the response individually
func (c *Client) DoRequest(method, url string, attempt int) ([]byte, *http.Response, error) {
	if c.Debug {
		c.Logger.Debug(fmt.Sprintf("%s %s", method, url))
	}

	// Get the bucket for the URL
	key := strings.Split(url, "?")[0]
	bucket := c.RateLimiter.LockBucket(key)

	// Prepare the request and set the relevant headers
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		_ = c.RateLimiter.UnlockBucket(bucket, nil)
		return nil, nil, err
	}
	request.Header.Set("Authorization", c.authorizationHeader)
	request.Header.Set("User-Agent", fmt.Sprintf("DiscordBot (%s, %s)", constants.GitHubURL, constants.Version))

	response, err := c.HttpClient.Do(request)
	if err != nil {
		_ = c.RateLimiter.UnlockBucket(bucket, nil)
		return nil, nil, err
	}

	err = c.RateLimiter.UnlockBucket(bucket, response.Header)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	switch response.StatusCode {
	case http.StatusUnauthorized:
		return nil, nil, errors.Unauthorized
	case http.StatusTooManyRequests:
		rateLimitExceeded := discord.RateLimitExceeded{}
		err = json.Unmarshal(body, &rateLimitExceeded)
		if err != nil {
			return nil, nil, err
		}
		if attempt > c.RateLimiter.MaxRetries {
			return nil, nil, &errors.RateLimitExceeded{RateLimitExceeded: &rateLimitExceeded}
		}

		c.Logger.Info(fmt.Sprintf("Got rate limited, trying again in %v [Attempt %d of %d]", rateLimitExceeded.RetryAfter, attempt, c.RateLimiter.MaxRetries))

		integer, frac := math.Modf(rateLimitExceeded.RetryAfter)
		time.Sleep(time.Duration(integer)*time.Second + time.Duration(frac*1000)*time.Millisecond)
		return c.DoRequest(method, url, attempt+1)
	}

	return body, response, nil
}
