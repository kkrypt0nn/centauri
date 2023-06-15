package rest

import "github.com/kkrypt0nn/centauri/discord"

// GetGlobalApplicationCommands returns a list of global application command structures (discord.ApplicationCommand) for the given application ID
func (c *Client) GetGlobalApplicationCommands(applicationID discord.Snowflake, withLocalizations bool) ([]discord.ApplicationCommand, error) {
	queryParams := make(QueryParameters)
	if withLocalizations {
		queryParams["with_localizations"] = "true"
	}
	return DoRequestAsList[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID.String()+"/commands", nil, queryParams, 1)
}

// CreateGlobalApplicationCommand creates a global application command (discord.ApplicationCommand) for the given application ID and returns its structure
func (c *Client) CreateGlobalApplicationCommand(applicationID discord.Snowflake, applicationCommand discord.CreateGlobalApplicationCommand) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "POST", ApplicationsEndpoint+"/"+applicationID.String()+"/commands", applicationCommand, nil, 1)
}

// GetGlobalApplicationCommand returns a global application command structure (discord.ApplicationCommand) for the given application and command IDs
func (c *Client) GetGlobalApplicationCommand(applicationID, commandID discord.Snowflake) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID.String()+"/commands/"+commandID.String(), nil, nil, 1)
}

// EditGlobalApplicationCommand edits and existing global application command (discord.ApplicationCommand) for the given application and command IDs and returns its new structure
func (c *Client) EditGlobalApplicationCommand(applicationID, commandID discord.Snowflake, applicationCommand discord.EditGlobalApplicationCommand) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "PATCH", ApplicationsEndpoint+"/"+applicationID.String()+"/commands/"+commandID.String(), applicationCommand, nil, 1)
}

// DeleteGlobalApplicationCommand deletes an existing global application command (discord.Guild) from the given application and command IDs
func (c *Client) DeleteGlobalApplicationCommand(applicationID, commandID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", ApplicationsEndpoint+"/"+applicationID.String()+"/commands/"+commandID.String(), nil, nil, 1)
}

// BulkOverwriteGlobalApplicationCommand overwrites all the existing global application commands (discord.ApplicationCommand) for the given application ID and returns a list of application command structures
func (c *Client) BulkOverwriteGlobalApplicationCommand(applicationID discord.Snowflake, applicationCommands []discord.ApplicationCommand) ([]discord.ApplicationCommand, error) {
	return DoRequestAsList[discord.ApplicationCommand](c, "PUT", ApplicationsEndpoint+"/"+applicationID.String()+"/commands", applicationCommands, nil, 1)
}

// GetGuildApplicationCommands returns a list of guild application command structures (discord.ApplicationCommand) for the given application ID that is available in the given guild ID
func (c *Client) GetGuildApplicationCommands(applicationID, guildID discord.Snowflake, withLocalizations bool) ([]discord.ApplicationCommand, error) {
	queryParams := make(QueryParameters)
	if withLocalizations {
		queryParams["with_localizations"] = "true"
	}
	return DoRequestAsList[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID.String()+"/guilds/"+guildID.String()+"/commands", nil, queryParams, 1)
}

// CreateGuildApplicationCommand creates a guild application command (discord.ApplicationCommand) for the given application and guild IDs and returns its structure
func (c *Client) CreateGuildApplicationCommand(applicationID, guildID discord.Snowflake, applicationCommand discord.CreateGuildApplicationCommand) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "POST", ApplicationsEndpoint+"/"+applicationID.String()+"/guilds/"+guildID.String()+"/commands", applicationCommand, nil, 1)
}

// GetGuildApplicationCommand returns a guild application command structure (discord.ApplicationCommand) for the given application, guild and command IDs that is available in the given guild ID
func (c *Client) GetGuildApplicationCommand(applicationID, guildID, commandID discord.Snowflake) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "GET", ApplicationsEndpoint+"/"+applicationID.String()+"/guilds/"+guildID.String()+"/commands/"+commandID.String(), nil, nil, 1)
}

// EditGuildApplicationCommand edits and existing guild application command (discord.ApplicationCommand) for the given application, guild and command IDs and returns its new structure
func (c *Client) EditGuildApplicationCommand(applicationID, guildID, commandID discord.Snowflake, applicationCommand discord.EditGuildApplicationCommand) (*discord.ApplicationCommand, error) {
	return DoRequestAsStructure[discord.ApplicationCommand](c, "PATCH", ApplicationsEndpoint+"/"+applicationID.String()+"/guilds/"+guildID.String()+"/commands/"+commandID.String(), applicationCommand, nil, 1)
}

// DeleteGuildApplicationCommand deletes an existing guild application command (discord.Guild) from the given application, guild and command IDs
func (c *Client) DeleteGuildApplicationCommand(applicationID, guildID, commandID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", ApplicationsEndpoint+"/"+applicationID.String()+"/guilds/"+guildID.String()+"/commands/"+commandID.String(), nil, nil, 1)
}

// BulkOverwriteGuildApplicationCommand overwrites all the existing guild application commands (discord.ApplicationCommand) for the given application and guild IDs and returns a list of application command structures
func (c *Client) BulkOverwriteGuildApplicationCommand(applicationID, guildID discord.Snowflake, applicationCommands []discord.ApplicationCommand) ([]discord.ApplicationCommand, error) {
	return DoRequestAsList[discord.ApplicationCommand](c, "PUT", ApplicationsEndpoint+"/"+applicationID.String()+"/guilds/"+guildID.String()+"/commands", applicationCommands, nil, 1)
}

// GetGuildApplicationCommandPermissions returns a list of guild application command permissions structures (discord.GuildApplicationCommandPermissions) for the given application and guild IDs that is available in the guild ID
func (c *Client) GetGuildApplicationCommandPermissions(applicationID, guildID discord.Snowflake) ([]discord.GuildApplicationCommandPermissions, error) {
	return DoRequestAsList[discord.GuildApplicationCommandPermissions](c, "GET", ApplicationsEndpoint+"/"+applicationID.String()+"/guilds/"+guildID.String()+"/commands/permissions", nil, nil, 1)
}

// GetApplicationCommandPermissions returns a guild application command permissions structure (discord.GuildApplicationCommandPermissions) for the given application, guild and command IDs that is available in the guild ID
func (c *Client) GetApplicationCommandPermissions(applicationID, guildID, commandID discord.Snowflake) ([]discord.GuildApplicationCommandPermissions, error) {
	return DoRequestAsList[discord.GuildApplicationCommandPermissions](c, "GET", ApplicationsEndpoint+"/"+applicationID.String()+"/guilds/"+guildID.String()+"/commands/"+commandID.String()+"/permissions", nil, nil, 1)
}
