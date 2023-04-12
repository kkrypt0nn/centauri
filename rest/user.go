package rest

import (
	"encoding/json"
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/errors"
	"io"
	"net/http"
)

const (
	UsersEndpoint = Endpoint + "users"
)

// GetCurrentUser returns the current discord.User
func (c *Client) GetCurrentUser() (*discord.User, error) {
	requestEndpoint := UsersEndpoint + "/@me"
	if c.Debug {
		c.Logger.Debug(fmt.Sprintf("GET %s", requestEndpoint))
	}

	request, err := http.NewRequest("GET", requestEndpoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", c.authorizationHeader)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusUnauthorized:
		return nil, errors.Unauthorized
	}
	defer response.Body.Close()

	var user *discord.User
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return user, err
}

// GetSelfUser is an alias of GetCurrentUser and returns the current discord.User
func (c *Client) GetSelfUser() (*discord.User, error) {
	return c.GetCurrentUser()
}
