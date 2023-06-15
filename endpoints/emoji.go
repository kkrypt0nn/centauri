package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

func GuildEmojis(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/emojis", GuildsEndpoint, guildID)
}

func GuildEmoji(guildID, emojiID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/emojis/%d", GuildsEndpoint, guildID, emojiID)
}
