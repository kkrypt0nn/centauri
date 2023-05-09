package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"strconv"
)

const (
	UsersEndpoint = Endpoint + "users"
)

// GetCurrentUser returns the current discord.User
func (c *Client) GetCurrentUser() (*discord.User, error) {
	return DoRequestAs[discord.User](c, "GET", UsersEndpoint+"/@me", nil, 1)
}

// GetSelfUser is an alias of GetCurrentUser and returns the current discord.User
func (c *Client) GetSelfUser() (*discord.User, error) {
	return c.GetCurrentUser()
}

// GetUser returns a discord.User based on the given ID
func (c *Client) GetUser(id string) (*discord.User, error) {
	return DoRequestAs[discord.User](c, "GET", UsersEndpoint+"/"+id, nil, 1)
}

// GetUserGuilds returns a list of discord.PartialGuild
func (c *Client) GetUserGuilds(before, after string, limit int) ([]discord.PartialGuild, error) {
	queryParams := make(QueryParameters)
	if before != "" {
		queryParams["before"] = before
	}
	if after != "" {
		queryParams["after"] = after
	}
	if limit >= 1 && limit <= 200 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsList[discord.PartialGuild](c, "GET", UsersEndpoint+"/@me/guilds", queryParams, 1)
}
