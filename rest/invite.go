package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
)

const (
	InvitesEndpoint = Endpoint + "invites"
)

// GetInvite returns an invitation structure (discord.Invite) for the given invite code
func (c *Client) GetInvite(inviteCode string, withCounts, withExpiration bool, guildScheduledEventID string) (*discord.Invite, error) {
	queryParams := make(QueryParameters)
	if withCounts {
		queryParams["with_counts"] = "true"
	}
	if withExpiration {
		queryParams["with_expiration"] = "true"
	}
	if guildScheduledEventID != "" {
		queryParams["guild_scheduled_event_id"] = guildScheduledEventID
	}
	return DoRequestAs[discord.Invite](c, "GET", InvitesEndpoint+"/"+inviteCode, queryParams, 1)
}
