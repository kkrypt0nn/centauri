package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	WebhooksEndpoint = Endpoint + "webhooks"
)

// CreateWebhook creates a new webhook (discord.Webhook) for the given channel ID and returns its structure
func (c *Client) CreateWebhook(channelID string, webhook discord.CreateWebhook) (*discord.Webhook, error) {
	return DoRequestAsStructure[discord.Webhook](c, "POST", ChannelsEndpoint+"/"+channelID+"/webhooks", webhook, nil, 1, WithReason(webhook.AuditLogReason))
}

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

// ModifyWebhook modifies an existing webhook (discord.Webhook) for the given webhook ID and returns its new structure
func (c *Client) ModifyWebhook(webhookID string, webhook discord.ModifyWebhook) (*discord.Webhook, error) {
	return DoRequestAsStructure[discord.Webhook](c, "PATCH", WebhooksEndpoint+"/"+webhookID, webhook, nil, 1, WithReason(webhook.AuditLogReason))
}

// ModifyWebhookWithToken modifies an existing webhook (discord.Webhook) for the given webhook ID and webhook token and returns its new structure (does not require authentication and returns no user in the webhook structure)
func (c *Client) ModifyWebhookWithToken(webhookID, webhookToken string, webhook discord.ModifyWebhookWithToken) (*discord.Webhook, error) {
	return DoRequestAsStructure[discord.Webhook](c, "PATCH", ChannelsEndpoint+webhookID+"/"+webhookToken, webhook, nil, 1, WithReason(webhook.AuditLogReason))
}

// DeleteWebhook deletes an existing webhook (discord.Webhook) from the given webhook ID
func (c *Client) DeleteWebhook(webhookID string) error {
	return DoEmptyRequest(c, "DELETE", WebhooksEndpoint+"/"+webhookID, nil, nil, 1)
}

// DeleteWebhookWithToken deletes an existing webhook (discord.Webhook) from the given webhook ID and webhook token (does not require authentication)
func (c *Client) DeleteWebhookWithToken(webhookID, webhookToken string) error {
	return DoEmptyRequest(c, "DELETE", WebhooksEndpoint+"/"+webhookID+"/"+webhookToken, nil, nil, 1)
}

// ExecuteWebhook executes a webhook (discord.Webhook) to send some message with it for the given webhook ID and webhook token (can also be executed in a given thread ID)
func (c *Client) ExecuteWebhook(webhookID, webhookToken, threadID string, message discord.ExecuteWebhook) error {
	queryParams := make(QueryParameters)
	if threadID != "" {
		queryParams["thread_id"] = threadID
	}
	if len(message.Files) >= 1 {
		contentType, body, err := CreateMultipartBodyWithJSON(message, message.Files)
		if err != nil {
			return err
		}
		return DoEmptyRequestWithFiles(c, "POST", WebhooksEndpoint+"/"+webhookID+"/"+webhookToken, body, queryParams, 1, WithHeader("Content-Type", contentType))
	} else {
		return DoEmptyRequest(c, "POST", WebhooksEndpoint+"/"+webhookID+"/"+webhookToken, message, queryParams, 1)
	}
}

// GetWebhookMessage returns a previously-sent webhook message structure (discord.Message) from the same token given the message ID
func (c *Client) GetWebhookMessage(webhookID, webhookToken, messageID, threadID string) (*discord.Message, error) {
	queryParams := make(QueryParameters)
	if threadID != "" {
		queryParams["thread_id"] = threadID
	}
	return DoRequestAsStructure[discord.Message](c, "GET", WebhooksEndpoint+"/"+webhookID+"/"+webhookToken+"/messages/"+messageID, nil, queryParams, 1)
}
