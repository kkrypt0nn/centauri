package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

func GlobalApplicationCommands(applicationID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/commands", ApplicationsEndpoint, applicationID)
}

func GlobalApplicationCommand(applicationID, commandID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/commands/%d", ApplicationsEndpoint, applicationID, commandID)
}

func GuildApplicationCommands(applicationID, guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/guilds/%d/commands", ApplicationsEndpoint, applicationID, guildID)
}

func GuildApplicationCommand(applicationID, guildID, commandID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/guilds/%d/commands/%d", ApplicationsEndpoint, applicationID, guildID, commandID)
}

func GuildApplicationCommandPermissions(applicationID, guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/guilds/%d/commands/permissions", ApplicationsEndpoint, applicationID, guildID)
}

func ApplicationCommandPermissions(applicationID, guildID, commandID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/guilds/%d/commands/%d/permissions", ApplicationsEndpoint, applicationID, guildID, commandID)
}
