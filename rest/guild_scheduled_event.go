package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"strconv"
)

// ListGuildScheduledEvents returns a list of guild scheduled event structures (discord.GuildScheduledEvent) for the given guild ID
func (c *Client) ListGuildScheduledEvents(guildID string, withUserCount bool) ([]discord.GuildScheduledEvent, error) {
	queryParams := make(QueryParameters)
	if withUserCount {
		queryParams["with_user_count"] = "true"
	}
	return DoRequestAsList[discord.GuildScheduledEvent](c, "GET", GuildsEndpoint+"/"+guildID+"/scheduled-events", nil, queryParams, 1)
}

// GetGuildScheduledEvent returns a guild scheduled event structure (discord.GuildScheduledEvent) for the given guild ID and scheduled event ID
func (c *Client) GetGuildScheduledEvent(guildID, scheduledEventID string, withUserCount bool) (*discord.GuildScheduledEvent, error) {
	queryParams := make(QueryParameters)
	if withUserCount {
		queryParams["with_user_count"] = "true"
	}
	return DoRequestAsStructure[discord.GuildScheduledEvent](c, "GET", GuildsEndpoint+"/"+guildID+"/scheduled-events/"+scheduledEventID, nil, nil, 1)
}

// GetGuildScheduledEventUsers returns a list of guild scheduled event user structures (discord.GuildScheduledEventUser)
func (c *Client) GetGuildScheduledEventUsers(guildID, scheduledEventID, before, after string, limit int, withMember bool) ([]discord.GuildScheduledEventUser, error) {
	queryParams := make(QueryParameters)
	if before != "" {
		queryParams["before"] = before
	}
	if after != "" {
		queryParams["after"] = after
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	if withMember {
		queryParams["with_member"] = "true"
	}
	return DoRequestAsList[discord.GuildScheduledEventUser](c, "GET", GuildsEndpoint+"/"+guildID+"/scheduled-events/"+scheduledEventID+"/users", nil, queryParams, 1)
}
