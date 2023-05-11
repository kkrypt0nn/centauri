package rest

import "github.com/kkrypt0nn/centauri/discord"

// GetGuildAuditLog returns the discord.AuditLog of the given guild ID
func (c *Client) GetGuildAuditLog(guildID string) (*discord.AuditLog, error) {
	return DoRequestAs[discord.AuditLog](c, "GET", GuildsEndpoint+"/"+guildID+"/audit-logs", nil, 1)
}
