package rest

import (
	"encoding/json"
	"github.com/kkrypt0nn/centauri/discord"
)

const (
	UsersEndpoint = Endpoint + "users"
)

// GetCurrentUser returns the current discord.User
func (c *Client) GetCurrentUser() (*discord.User, error) {
	responseBody, _, err := c.DoRequest("GET", UsersEndpoint+"/@me", 1)
	if err != nil {
		return nil, err
	}

	var user *discord.User
	err = json.Unmarshal(responseBody, &user)
	if err != nil {
		return nil, err
	}
	return user, err
}

// GetSelfUser is an alias of GetCurrentUser and returns the current discord.User
func (c *Client) GetSelfUser() (*discord.User, error) {
	return c.GetCurrentUser()
}
