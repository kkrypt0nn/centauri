package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
)

// GetGlobalApplicationCommands returns a list of global discord.ApplicationCommand
func (c *Client) GetGlobalApplicationCommands(applicationID string, withLocalizations bool) ([]discord.ApplicationCommand, error) {
	queryParams := make(QueryParameters)
	if withLocalizations {
		queryParams["with_localizations"] = "true"
	}
	return DoRequestAsList[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/commands", queryParams, 1)
}

// GetGlobalApplicationCommand returns a global discord.ApplicationCommand
func (c *Client) GetGlobalApplicationCommand(applicationID, commandID string) (*discord.ApplicationCommand, error) {
	return DoRequestAs[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/commands/"+commandID, nil, 1)
}

// GetGuildApplicationCommands returns a list of discord.ApplicationCommand from a specific guild
func (c *Client) GetGuildApplicationCommands(applicationID, guildID string, withLocalizations bool) ([]discord.ApplicationCommand, error) {
	queryParams := make(QueryParameters)
	if withLocalizations {
		queryParams["with_localizations"] = "true"
	}
	return DoRequestAsList[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/guilds/"+guildID+"/commands", queryParams, 1)
}

// GetGuildApplicationCommand returns a discord.ApplicationCommand from a specific guild
func (c *Client) GetGuildApplicationCommand(applicationID, guildID, commandID string) (*discord.ApplicationCommand, error) {
	return DoRequestAs[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/guilds/"+guildID+"/commands/"+commandID, nil, 1)
}

// GetGuildApplicationCommandPermissions returns a list of discord.GuildApplicationCommandPermissions from a guild
func (c *Client) GetGuildApplicationCommandPermissions(applicationID, guildID string) ([]discord.GuildApplicationCommandPermissions, error) {
	return DoRequestAsList[discord.GuildApplicationCommandPermissions](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/guilds/"+guildID+"/commands/permissions", nil, 1)
}

// GetApplicationCommandPermissions returns a discord.GuildApplicationCommandPermissions for a specific command from a guild
func (c *Client) GetApplicationCommandPermissions(applicationID, guildID, commandID string) ([]discord.GuildApplicationCommandPermissions, error) {
	return DoRequestAsList[discord.GuildApplicationCommandPermissions](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/guilds/"+guildID+"/commands/"+commandID+"/permissions", nil, 1)
}
