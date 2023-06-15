package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	InvitesEndpoint = Endpoint + "invites"
)

// GetInvite returns an invitation structure (discord.Invite) for the given invite code
func (c *Client) GetInvite(inviteCode string, withCounts, withExpiration bool, guildScheduledEventID discord.Snowflake) (*discord.Invite, error) {
	queryParams := make(QueryParameters)
	if withCounts {
		queryParams["with_counts"] = "true"
	}
	if withExpiration {
		queryParams["with_expiration"] = "true"
	}
	if guildScheduledEventID != 0 {
		queryParams["guild_scheduled_event_id"] = guildScheduledEventID.String()
	}
	return DoRequestAsStructure[discord.Invite](c, "GET", InvitesEndpoint+"/"+inviteCode, nil, queryParams, 1)
}

// DeleteInvite deletes an existing invitation (discord.Invite) for the given invite code
func (c *Client) DeleteInvite(inviteCode string) (*discord.Invite, error) {
	return DoRequestAsStructure[discord.Invite](c, "DELETE", InvitesEndpoint+"/"+inviteCode, nil, nil, 1)
}
