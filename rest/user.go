package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"strconv"
)

const (
	UsersEndpoint = Endpoint + "users"
)

// GetCurrentUser returns the current user structure (discord.User)
func (c *Client) GetCurrentUser() (*discord.User, error) {
	return DoRequestAsStructure[discord.User](c, "GET", UsersEndpoint+"/@me", nil, nil, 1)
}

// GetSelfUser is an alias of GetCurrentUser
func (c *Client) GetSelfUser() (*discord.User, error) {
	return c.GetCurrentUser()
}

// GetUser returns a user structure (discord.User) for the given user ID
func (c *Client) GetUser(userID string) (*discord.User, error) {
	return DoRequestAsStructure[discord.User](c, "GET", UsersEndpoint+"/"+userID, nil, nil, 1)
}

// GetUserGuilds returns a list of partial guild structures (discord.PartialGuild)
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
	return DoRequestAsList[discord.PartialGuild](c, "GET", UsersEndpoint+"/@me/guilds", nil, queryParams, 1)
}
