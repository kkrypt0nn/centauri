package rest

import (
	"strconv"

	"github.com/kkrypt0nn/centauri/discord"
)

// ListGuildScheduledEvents returns a list of guild scheduled event structures (discord.GuildScheduledEvent) for the given guild ID
func (c *Client) ListGuildScheduledEvents(guildID string, withUserCount bool) ([]discord.GuildScheduledEvent, error) {
	queryParams := make(QueryParameters)
	if withUserCount {
		queryParams["with_user_count"] = "true"
	}
	return DoRequestAsList[discord.GuildScheduledEvent](c, "GET", GuildsEndpoint+"/"+guildID+"/scheduled-events", nil, queryParams, 1)
}

// CreateGuildScheduledEvent creates a guild scheduled event (discord.GuildScheduledEvent) for the given guild ID and returns its structure
func (c *Client) CreateGuildScheduledEvent(guildID string, scheduledEvent discord.CreateGuildScheduledEvent) (*discord.GuildScheduledEvent, error) {
	return DoRequestAsStructure[discord.GuildScheduledEvent](c, "POST", GuildsEndpoint+"/"+guildID+"/scheduled-events", scheduledEvent, nil, 1, WithReason(scheduledEvent.AuditLogReason))
}

// GetGuildScheduledEvent returns a guild scheduled event structure (discord.GuildScheduledEvent) for the given guild ID and scheduled event ID
func (c *Client) GetGuildScheduledEvent(guildID, scheduledEventID string, withUserCount bool) (*discord.GuildScheduledEvent, error) {
	queryParams := make(QueryParameters)
	if withUserCount {
		queryParams["with_user_count"] = "true"
	}
	return DoRequestAsStructure[discord.GuildScheduledEvent](c, "GET", GuildsEndpoint+"/"+guildID+"/scheduled-events/"+scheduledEventID, nil, nil, 1)
}

// ModifyGuildScheduledEvent modifies an existing guild scheduled event (discord.GuildScheduledEvent) for the given guild and scheduled event IDs and returns its new structure
func (c *Client) ModifyGuildScheduledEvent(guildID, scheduledEventID string, scheduledEvent discord.ModifyGuildScheduledEvent) (*discord.GuildScheduledEvent, error) {
	return DoRequestAsStructure[discord.GuildScheduledEvent](c, "PATCH", GuildsEndpoint+"/"+guildID+"/scheduled-events/"+scheduledEventID, scheduledEvent, nil, 1, WithReason(scheduledEvent.AuditLogReason))
}

// DeleteGuildScheduledEvent deletes an existing guild scheduled event (discord.GuildScheduledEvent) for the given guild and scheduled event IDs
func (c *Client) DeleteGuildScheduledEvent(guildID, scheduledEventID string) error {
	return DoEmptyRequest(c, "DELETE", GuildsEndpoint+"/"+guildID+"/scheduled-events/"+scheduledEventID, nil, nil, 1)
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
