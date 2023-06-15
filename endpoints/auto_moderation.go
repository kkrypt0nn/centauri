package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

func AutoModerationRules(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/auto-moderation/rules", GuildsEndpoint, guildID)
}

func AutoModerationRule(guildID, ruleID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/auto-moderation/rules/%d", GuildsEndpoint, guildID, ruleID)
}
