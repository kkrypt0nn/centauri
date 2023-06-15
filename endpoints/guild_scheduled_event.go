package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

func GuildScheduledEvents(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/scheduled-events", GuildsEndpoint, guildID)
}

func GuildScheduledEvent(guildID, scheduledEventID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/scheduled-events/%d", GuildsEndpoint, guildID, scheduledEventID)
}

func GuildScheduledEventUsers(guildID, scheduledEventID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/scheduled-events/%d/users", GuildsEndpoint, guildID, scheduledEventID)
}
