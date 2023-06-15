package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

func Template(templateCode string) string {
	return fmt.Sprintf("%s/templates/%s", GuildsEndpoint, templateCode)
}

func GuildTemplates(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/templates", GuildsEndpoint, guildID)
}

func GuildTemplate(guildID discord.Snowflake, templateCode string) string {
	return fmt.Sprintf("%s/%d/templates/%s", GuildsEndpoint, guildID, templateCode)
}
