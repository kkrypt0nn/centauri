package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

const (
	WebhooksEndpoint = Endpoint + "webhooks"
)

func ChannelWebhooks(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/webhooks", ChannelsEndpoint, channelID)
}

func GuildWebhooks(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/webhooks", GuildsEndpoint, guildID)
}

func Webhook(webhookID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d", WebhooksEndpoint, webhookID)
}

func WebhookWithToken(webhookID discord.Snowflake, webhookToken string) string {
	return fmt.Sprintf("%s/%d/%s", WebhooksEndpoint, webhookID, webhookToken)
}

func WebhookMessage(webhookID discord.Snowflake, webhookToken string, messageID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/%s/messages/%d", WebhooksEndpoint, webhookID, webhookToken, messageID)
}
