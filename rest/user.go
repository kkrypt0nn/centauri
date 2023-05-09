package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
)

const (
	UsersEndpoint = Endpoint + "users"
)

// GetCurrentUser returns the current discord.User
func (c *Client) GetCurrentUser() (*discord.User, error) {
	return DoRequestAs[discord.User](c, "GET", UsersEndpoint+"/@me", 1)
}

// GetSelfUser is an alias of GetCurrentUser and returns the current discord.User
func (c *Client) GetSelfUser() (*discord.User, error) {
	return c.GetCurrentUser()
}

// GetUser returns a discord.User based on the given ID
func (c *Client) GetUser(id string) (*discord.User, error) {
	return DoRequestAs[discord.User](c, "GET", UsersEndpoint+"/"+id, 1)
}
