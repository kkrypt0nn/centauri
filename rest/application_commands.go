package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
)

// GetGlobalApplicationCommands returns a list of global application command structures (discord.ApplicationCommand) for the given application ID
func (c *Client) GetGlobalApplicationCommands(applicationID string, withLocalizations bool) ([]discord.ApplicationCommand, error) {
	queryParams := make(QueryParameters)
	if withLocalizations {
		queryParams["with_localizations"] = "true"
	}
	return DoRequestAsList[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/commands", nil, queryParams, 1)
}

// GetGlobalApplicationCommand returns a global application command structure (discord.ApplicationCommand) for the given application and command IDs
func (c *Client) GetGlobalApplicationCommand(applicationID, commandID string) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/commands/"+commandID, nil, nil, 1)
}

// GetGuildApplicationCommands returns a list of guild application command structures (discord.ApplicationCommand) for the given application ID that is available in the given guild ID
func (c *Client) GetGuildApplicationCommands(applicationID, guildID string, withLocalizations bool) ([]discord.ApplicationCommand, error) {
	queryParams := make(QueryParameters)
	if withLocalizations {
		queryParams["with_localizations"] = "true"
	}
	return DoRequestAsList[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/guilds/"+guildID+"/commands", nil, queryParams, 1)
}

// GetGuildApplicationCommand returns a guild application command structure (discord.ApplicationCommand) for the given application, guild and command IDs that is available in the given guild ID
func (c *Client) GetGuildApplicationCommand(applicationID, guildID, commandID string) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/guilds/"+guildID+"/commands/"+commandID, nil, nil, 1)
}

// GetGuildApplicationCommandPermissions returns a list of guild application command permissions structures (discord.GuildApplicationCommandPermissions) for the given application and guild IDs that is available in the guild ID
func (c *Client) GetGuildApplicationCommandPermissions(applicationID, guildID string) ([]discord.GuildApplicationCommandPermissions, error) {
	return DoRequestAsList[discord.GuildApplicationCommandPermissions](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/guilds/"+guildID+"/commands/permissions", nil, nil, 1)
}

// GetApplicationCommandPermissions returns a guild application command permissions structure (discord.GuildApplicationCommandPermissions) for the given application, guild and command IDs that is available in the guild ID
func (c *Client) GetApplicationCommandPermissions(applicationID, guildID, commandID string) ([]discord.GuildApplicationCommandPermissions, error) {
	return DoRequestAsList[discord.GuildApplicationCommandPermissions](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/guilds/"+guildID+"/commands/"+commandID+"/permissions", nil, nil, 1)
}
