// Package rest handles the REST API of Discord
package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kkrypt0nn/centauri/constants"
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/errors"
	"github.com/kkrypt0nn/centauri/ext/tasks"
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
	HttpClient  *HttpClient
	Logger      *logger.Logger
	Debug       bool
	RateLimiter *RateLimiter
	TaskManager *tasks.TaskManager

	// Extras is a field you can use freely that will **never** be touched by the library itself
	Extras interface{}

	isBot               bool
	authorizationHeader string
}

type QueryParameters map[string]string

// SetAuthorizationHeader sets the authorization header of the RestClient
func (c *Client) SetAuthorizationHeader(authorizationHeader string) {
	c.authorizationHeader = authorizationHeader
	if strings.HasPrefix(c.authorizationHeader, "Bot ") {
		c.isBot = true
	}
}

// DoRequest performs a request to the given URL with the given method, it will make sure to follow rate limits, and returns the body and the response individually
func (c *Client) DoRequest(method, url string, requestBody any, queryParams QueryParameters, attempt int, requestOptions ...RequestOption) ([]byte, *http.Response, error) {
	// Get the bucket for the URL
	key := strings.Split(url, "?")[0]
	bucket := c.RateLimiter.LockBucket(key)

	// Marshal the request body
	var bytesBody []byte
	if requestBody != nil {
		var err error
		if bytesBody, err = json.Marshal(requestBody); err != nil {
			return nil, nil, err
		}
	}

	// Prepare the request and set the relevant headers
	request, err := http.NewRequest(method, url, bytes.NewReader(bytesBody))
	if err != nil {
		_ = c.RateLimiter.UnlockBucket(bucket, nil)
		return nil, nil, err
	}
	request.Header.Set("Authorization", c.authorizationHeader)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", fmt.Sprintf("DiscordBot (%s, %s)", constants.GitHubURL, constants.Version))

	// Set the query parameters
	if queryParams != nil {
		query := request.URL.Query()
		for key, value := range queryParams {
			query.Add(key, value)
		}
		request.URL.RawQuery = query.Encode()
	}

	// Apply the request options by calling them
	for _, option := range requestOptions {
		option(request)
	}

	// Modify the request if an interceptor is set
	if c.HttpClient.Interceptor != nil {
		c.HttpClient.Interceptor.ModifyRequest(request)
	}

	// Perform the request
	if c.Debug {
		if requestBody != nil {
			c.Logger.Debug(fmt.Sprintf("%s %s, body: %s", method, request.URL.String(), string(bytesBody)))
		} else {
			c.Logger.Debug(fmt.Sprintf("%s %s", method, request.URL.String()))
		}
	}
	response, err := c.HttpClient.Do(request)
	if err != nil {
		_ = c.RateLimiter.UnlockBucket(bucket, nil)
		return nil, nil, err
	}

	err = c.RateLimiter.UnlockBucket(bucket, response.Header)
	if err != nil {
		return nil, nil, err
	}

	// Modify the response if an interceptor is set
	if c.HttpClient.Interceptor != nil {
		c.HttpClient.Interceptor.ModifyResponse(response)
	}

	defer func() {
		_ = response.Body.Close()
	}()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		return body, response, nil
	case http.StatusTooManyRequests:
		var rateLimitExceeded discord.RateLimitExceeded
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
		return c.DoRequest(method, url, requestBody, queryParams, attempt+1, requestOptions...)
	default:
		return body, response, discord.NewError(body, request, response)
	}
}

func DoRequestAsStructure[T any](client *Client, method, url string, requestBody any, queryParams QueryParameters, attempt int, requestOptions ...RequestOption) (*T, error) {
	responseBody, _, err := client.DoRequest(method, url, requestBody, queryParams, attempt, requestOptions...)
	if err != nil {
		return nil, err
	}

	var entity *T
	err = json.Unmarshal(responseBody, &entity)
	if err != nil {
		return nil, err
	}
	return entity, err
}

func DoRequestAsList[T any](client *Client, method, url string, requestBody any, queryParams QueryParameters, attempt int, requestOptions ...RequestOption) ([]T, error) {
	responseBody, _, err := client.DoRequest(method, url, requestBody, queryParams, attempt, requestOptions...)
	if err != nil {
		return nil, err
	}

	var entity []T
	err = json.Unmarshal(responseBody, &entity)
	if err != nil {
		return nil, err
	}
	return entity, err
}
