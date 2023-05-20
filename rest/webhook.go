package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	WebhooksEndpoint = Endpoint + "webhooks"
)

// GetChannelWebhooks returns a list of channel webhook structures (discord.Webhook) given the channel ID
func (c *Client) GetChannelWebhooks(channelID string) ([]discord.Webhook, error) {
	return DoRequestAsList[discord.Webhook](c, "GET", ChannelsEndpoint+"/"+channelID+"/webhooks", nil, nil, 1)
}

// GetGuildWebhooks returns a list of guild webhook structures (discord.Webhook) given the guild ID
func (c *Client) GetGuildWebhooks(guildID string) ([]discord.Webhook, error) {
	return DoRequestAsList[discord.Webhook](c, "GET", GuildsEndpoint+"/"+guildID+"/webhooks", nil, nil, 1)
}

// GetWebhook returns a webhook structure (discord.Webhook) given the webhook ID
func (c *Client) GetWebhook(webhookID string) (*discord.Webhook, error) {
	return DoRequestAsStructure[discord.Webhook](c, "GET", WebhooksEndpoint+"/"+webhookID, nil, nil, 1)
}

// GetWebhookWithToken returns a webhook structure (discord.Webhook) given the webhook ID and webhook token (does not require authentication and returns no user in the webhook structure)
func (c *Client) GetWebhookWithToken(webhookID, webhookToken string) (*discord.Webhook, error) {
	return DoRequestAsStructure[discord.Webhook](c, "GET", WebhooksEndpoint+"/"+webhookID+"/"+webhookToken, nil, nil, 1)
}

// GetWebhookMessage returns a previously-sent webhook message structure (discord.Message) from the same token given the message ID
func (c *Client) GetWebhookMessage(webhookID, webhookToken, messageID, threadID string) (*discord.Message, error) {
	queryParams := make(QueryParameters)
	if threadID != "" {
		queryParams["thread_id"] = threadID
	}
	return DoRequestAsStructure[discord.Message](c, "GET", WebhooksEndpoint+"/"+webhookID+"/"+webhookToken+"/messages/"+messageID, nil, queryParams, 1)
}
