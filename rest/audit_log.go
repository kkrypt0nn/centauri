package rest

import "github.com/kkrypt0nn/centauri/discord"

// GetGuildAuditLog returns an audit log structure (discord.AuditLog) for given guild ID
func (c *Client) GetGuildAuditLog(guildID string) (*discord.AuditLog, error) {
	return DoRequestAsStructure[discord.AuditLog](c, "GET", GuildsEndpoint+"/"+guildID+"/audit-logs", nil, nil, 1)
}
