package rest

import (
	"strconv"

	"github.com/kkrypt0nn/centauri/discord"
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
func (c *Client) GetUser(userID discord.Snowflake) (*discord.User, error) {
	return DoRequestAsStructure[discord.User](c, "GET", UsersEndpoint+"/"+userID.String(), nil, nil, 1)
}

// ModifyCurrentUser modifies the current user (discord.User) and returns its new structure
func (c *Client) ModifyCurrentUser(user discord.ModifyCurrentUser) (*discord.User, error) {
	return DoRequestAsStructure[discord.User](c, "PATCH", UsersEndpoint+"/@me", user, nil, 1)
}

// GetCurrentUserGuilds returns a list of partial guild structures (discord.PartialGuild)
func (c *Client) GetCurrentUserGuilds(before, after discord.Snowflake, limit int) ([]discord.PartialGuild, error) {
	queryParams := make(QueryParameters)
	if before != 0 {
		queryParams["before"] = before.String()
	}
	if after != 0 {
		queryParams["after"] = after.String()
	}
	if limit >= 1 && limit <= 200 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsList[discord.PartialGuild](c, "GET", UsersEndpoint+"/@me/guilds", nil, queryParams, 1)
}

// LeaveGuild leaves a guild (discord.Guild) from the given guild ID
func (c *Client) LeaveGuild(guildID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", GuildsEndpoint+"/"+guildID.String(), nil, nil, 1)
}

// CreateDM creates a new DM channel (discord.Channel) for the given user ID and returns its structure
func (c *Client) CreateDM(userID discord.Snowflake) (*discord.Channel, error) {
	payload := struct {
		RecipientID discord.Snowflake `json:"recipient_id"`
	}{
		RecipientID: userID,
	}
	return DoRequestAsStructure[discord.Channel](c, "POST", UsersEndpoint+"/@me/channels", payload, nil, 1)
}
