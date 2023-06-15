package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

func GuildAuditLog(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/audit-logs", GuildsEndpoint, guildID)
}
