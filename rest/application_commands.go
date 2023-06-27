package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
)

// GetGlobalApplicationCommands returns a list of global application command structures (discord.ApplicationCommand)
func (c *Client) GetGlobalApplicationCommands(withLocalizations bool) ([]discord.ApplicationCommand, error) {
	queryParams := make(QueryParameters)
	if withLocalizations {
		queryParams["with_localizations"] = "true"
	}
	return DoRequestAsList[discord.ApplicationCommand](c, "GET", endpoints.GlobalApplicationCommands(c.selfUser.ID), nil, queryParams, 1)
}

// CreateGlobalApplicationCommand creates a global application command (discord.ApplicationCommand) and returns its structure
func (c *Client) CreateGlobalApplicationCommand(applicationCommand discord.CreateGlobalApplicationCommand) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "POST", endpoints.GlobalApplicationCommands(c.selfUser.ID), applicationCommand, nil, 1)
}

// GetGlobalApplicationCommand returns a global application command structure (discord.ApplicationCommand) for the given command ID
func (c *Client) GetGlobalApplicationCommand(commandID discord.Snowflake) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "GET", endpoints.GlobalApplicationCommand(c.selfUser.ID, commandID), nil, nil, 1)
}

// EditGlobalApplicationCommand edits and existing global application command (discord.ApplicationCommand) for the given command ID and returns its new structure
func (c *Client) EditGlobalApplicationCommand(commandID discord.Snowflake, applicationCommand discord.EditGlobalApplicationCommand) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "PATCH", endpoints.GlobalApplicationCommand(c.selfUser.ID, commandID), applicationCommand, nil, 1)
}

// DeleteGlobalApplicationCommand deletes an existing global application command (discord.Guild) from the given command ID
func (c *Client) DeleteGlobalApplicationCommand(commandID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", endpoints.GlobalApplicationCommand(c.selfUser.ID, commandID), nil, nil, 1)
}

// BulkOverwriteGlobalApplicationCommand overwrites all the existing global application commands (discord.ApplicationCommand) and returns a list of application command structures
func (c *Client) BulkOverwriteGlobalApplicationCommand(applicationCommands []discord.CreateGlobalApplicationCommand) ([]discord.ApplicationCommand, error) {
	return DoRequestAsList[discord.ApplicationCommand](c, "PUT", endpoints.GlobalApplicationCommands(c.selfUser.ID), applicationCommands, nil, 1)
}

// GetGuildApplicationCommands returns a list of guild application command structures (discord.ApplicationCommand) that are available in the given guild ID
func (c *Client) GetGuildApplicationCommands(guildID discord.Snowflake, withLocalizations bool) ([]discord.ApplicationCommand, error) {
	queryParams := make(QueryParameters)
	if withLocalizations {
		queryParams["with_localizations"] = "true"
	}
	return DoRequestAsList[discord.ApplicationCommand](c, "GET", endpoints.GuildApplicationCommands(c.selfUser.ID, guildID), nil, queryParams, 1)
}

// CreateGuildApplicationCommand creates a guild application command (discord.ApplicationCommand) for the given guild ID and returns its structure
func (c *Client) CreateGuildApplicationCommand(guildID discord.Snowflake, applicationCommand discord.CreateGuildApplicationCommand) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "POST", endpoints.GuildApplicationCommands(c.selfUser.ID, guildID), applicationCommand, nil, 1)
}

// GetGuildApplicationCommand returns a guild application command structure (discord.ApplicationCommand) for the given guild and command IDs that is available in the given guild ID
func (c *Client) GetGuildApplicationCommand(guildID, commandID discord.Snowflake) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "GET", endpoints.GuildApplicationCommand(c.selfUser.ID, guildID, commandID), nil, nil, 1)
}

// EditGuildApplicationCommand edits and existing guild application command (discord.ApplicationCommand) for the given guild and command IDs and returns its new structure
func (c *Client) EditGuildApplicationCommand(guildID, commandID discord.Snowflake, applicationCommand discord.EditGuildApplicationCommand) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "PATCH", endpoints.GuildApplicationCommand(c.selfUser.ID, guildID, commandID), applicationCommand, nil, 1)
}

// DeleteGuildApplicationCommand deletes an existing guild application command (discord.Guild) from the given guild and command IDs
func (c *Client) DeleteGuildApplicationCommand(guildID, commandID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", endpoints.GuildApplicationCommand(c.selfUser.ID, guildID, commandID), nil, nil, 1)
}

// BulkOverwriteGuildApplicationCommand overwrites all the existing guild application commands (discord.ApplicationCommand) for the given guild ID and returns a list of application command structures
func (c *Client) BulkOverwriteGuildApplicationCommand(guildID discord.Snowflake, applicationCommands []discord.CreateGuildApplicationCommand) ([]discord.ApplicationCommand, error) {
	return DoRequestAsList[discord.ApplicationCommand](c, "PUT", endpoints.GuildApplicationCommands(c.selfUser.ID, guildID), applicationCommands, nil, 1)
}

// GetGuildApplicationCommandPermissions returns a list of guild application command permissions structures (discord.GuildApplicationCommandPermissions) for the given guild ID that is available in the guild ID
func (c *Client) GetGuildApplicationCommandPermissions(guildID discord.Snowflake) ([]discord.GuildApplicationCommandPermissions, error) {
	return DoRequestAsList[discord.GuildApplicationCommandPermissions](c, "GET", endpoints.GuildApplicationCommandPermissions(c.selfUser.ID, guildID), nil, nil, 1)
}

// GetApplicationCommandPermissions returns a guild application command permissions structure (discord.GuildApplicationCommandPermissions) for the given guild and command IDs that is available in the guild ID
func (c *Client) GetApplicationCommandPermissions(guildID, commandID discord.Snowflake) ([]discord.GuildApplicationCommandPermissions, error) {
	return DoRequestAsList[discord.GuildApplicationCommandPermissions](c, "GET", endpoints.ApplicationCommandPermissions(c.selfUser.ID, guildID, commandID), nil, nil, 1)
}
