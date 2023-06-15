package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
)

// GetGuildTemplate returns a guild template structure (discord.GuildTemplate) for the given template code
func (c *Client) GetGuildTemplate(templateCode string) (*discord.GuildTemplate, error) {
	return DoRequestAsStructure[discord.GuildTemplate](c, "GET", endpoints.Template(templateCode), nil, nil, 1)
}

// CreateGuildFromTemplate creates a guild (discord.Guild) based on the given template code and returns its structure
func (c *Client) CreateGuildFromTemplate(templateCode string, guild discord.CreateGuildFromTemplate) (*discord.Guild, error) {
	return DoRequestAsStructure[discord.Guild](c, "POST", endpoints.Template(templateCode), guild, nil, 1)
}

// GetGuildTemplates returns a list of guild template structures (discord.GuildTemplate) for the given guild ID
func (c *Client) GetGuildTemplates(guildID discord.Snowflake) ([]discord.GuildTemplate, error) {
	return DoRequestAsList[discord.GuildTemplate](c, "GET", endpoints.GuildTemplates(guildID), nil, nil, 1)
}

// CreateGuildTemplate creates a guild template (discord.GuildTemplate) for the given guild ID and returns its structure
func (c *Client) CreateGuildTemplate(guildID discord.Snowflake, guildTemplate discord.CreateGuildTemplate) (*discord.GuildTemplate, error) {
	return DoRequestAsStructure[discord.GuildTemplate](c, "POST", endpoints.GuildTemplates(guildID), guildTemplate, nil, 1)
}

// SyncGuildTemplate syncs an existing guild template (discord.GuildTemplate) for the given guild ID and template code and returns its structure
func (c *Client) SyncGuildTemplate(guildID discord.Snowflake, templateCode string) (*discord.GuildTemplate, error) {
	return DoRequestAsStructure[discord.GuildTemplate](c, "PUT", endpoints.GuildTemplate(guildID, templateCode), nil, nil, 1)
}

// ModifyGuildTemplate modifies an existing guild template (discord.GuildTemplate) for the given guild ID and template code and returns its new structure
func (c *Client) ModifyGuildTemplate(guildID discord.Snowflake, guildTemplate discord.ModifyGuildTemplate) (*discord.GuildTemplate, error) {
	return DoRequestAsStructure[discord.GuildTemplate](c, "PATCH", endpoints.GuildTemplates(guildID), guildTemplate, nil, 1)
}

// DeleteGuildTemplate deletes an existing guild template (discord.GuildTemplate) for the given guild ID and template code and returns the structure of the deleted template
func (c *Client) DeleteGuildTemplate(guildID discord.Snowflake, templateCode string) (*discord.GuildTemplate, error) {
	return DoRequestAsStructure[discord.GuildTemplate](c, "DELETE", endpoints.GuildTemplate(guildID, templateCode), nil, nil, 1)
}
