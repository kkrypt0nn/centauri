package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
)

// CreateWebhook creates a new webhook (discord.Webhook) for the given channel ID and returns its structure
func (c *Client) CreateWebhook(channelID discord.Snowflake, webhook discord.CreateWebhook) (*discord.Webhook, error) {
	return DoRequestAsStructure[discord.Webhook](c, "POST", endpoints.ChannelWebhooks(channelID), webhook, nil, 1, WithReason(webhook.AuditLogReason))
}

// GetChannelWebhooks returns a list of channel webhook structures (discord.Webhook) given the channel ID
func (c *Client) GetChannelWebhooks(channelID discord.Snowflake) ([]discord.Webhook, error) {
	return DoRequestAsList[discord.Webhook](c, "GET", endpoints.ChannelWebhooks(channelID), nil, nil, 1)
}

// GetGuildWebhooks returns a list of guild webhook structures (discord.Webhook) given the guild ID
func (c *Client) GetGuildWebhooks(guildID discord.Snowflake) ([]discord.Webhook, error) {
	return DoRequestAsList[discord.Webhook](c, "GET", endpoints.GuildWebhooks(guildID), nil, nil, 1)
}

// GetWebhook returns a webhook structure (discord.Webhook) given the webhook ID
func (c *Client) GetWebhook(webhookID discord.Snowflake) (*discord.Webhook, error) {
	return DoRequestAsStructure[discord.Webhook](c, "GET", endpoints.Webhook(webhookID), nil, nil, 1)
}

// GetWebhookWithToken returns a webhook structure (discord.Webhook) given the webhook ID and webhook token (does not require authentication and returns no user in the webhook structure)
func (c *Client) GetWebhookWithToken(webhookID discord.Snowflake, webhookToken string) (*discord.Webhook, error) {
	return DoRequestAsStructure[discord.Webhook](c, "GET", endpoints.WebhookWithToken(webhookID, webhookToken), nil, nil, 1)
}

// ModifyWebhook modifies an existing webhook (discord.Webhook) for the given webhook ID and returns its new structure
func (c *Client) ModifyWebhook(webhookID discord.Snowflake, webhook discord.ModifyWebhook) (*discord.Webhook, error) {
	return DoRequestAsStructure[discord.Webhook](c, "PATCH", endpoints.Webhook(webhookID), webhook, nil, 1, WithReason(webhook.AuditLogReason))
}

// ModifyWebhookWithToken modifies an existing webhook (discord.Webhook) for the given webhook ID and webhook token and returns its new structure (does not require authentication and returns no user in the webhook structure)
func (c *Client) ModifyWebhookWithToken(webhookID discord.Snowflake, webhookToken string, webhook discord.ModifyWebhookWithToken) (*discord.Webhook, error) {
	return DoRequestAsStructure[discord.Webhook](c, "PATCH", endpoints.WebhookWithToken(webhookID, webhookToken), webhook, nil, 1, WithReason(webhook.AuditLogReason))
}

// DeleteWebhook deletes an existing webhook (discord.Webhook) from the given webhook ID
func (c *Client) DeleteWebhook(webhookID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", endpoints.Webhook(webhookID), nil, nil, 1)
}

// DeleteWebhookWithToken deletes an existing webhook (discord.Webhook) from the given webhook ID and webhook token (does not require authentication)
func (c *Client) DeleteWebhookWithToken(webhookID discord.Snowflake, webhookToken string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.WebhookWithToken(webhookID, webhookToken), nil, nil, 1)
}

// ExecuteWebhook executes a webhook (discord.Webhook) to send some message with it for the given webhook ID and webhook token (can also be executed in a given thread ID)
func (c *Client) ExecuteWebhook(webhookID discord.Snowflake, webhookToken string, threadID discord.Snowflake, message discord.ExecuteWebhook) error {
	queryParams := make(QueryParameters)
	if threadID != 0 {
		queryParams["thread_id"] = threadID.String()
	}
	if len(message.Files) >= 1 {
		contentType, body, err := CreateMultipartBodyWithJSON(message, message.Files)
		if err != nil {
			return err
		}
		return DoEmptyRequestWithFiles(c, "POST", endpoints.WebhookWithToken(webhookID, webhookToken), body, queryParams, 1, WithHeader("Content-Type", contentType))
	} else {
		return DoEmptyRequest(c, "POST", endpoints.WebhookWithToken(webhookID, webhookToken), message, queryParams, 1)
	}
}

// GetWebhookMessage returns a previously-sent webhook message structure (discord.Message) from the same token given the message ID
func (c *Client) GetWebhookMessage(webhookID discord.Snowflake, webhookToken string, messageID, threadID discord.Snowflake) (*discord.Message, error) {
	queryParams := make(QueryParameters)
	if threadID != 0 {
		queryParams["thread_id"] = threadID.String()
	}
	return DoRequestAsStructure[discord.Message](c, "GET", endpoints.WebhookMessage(webhookID, webhookToken, messageID), nil, queryParams, 1)
}

// EditWebhookMessage edits a message (discord.Message) sent by a webhook (discord.Webhook) for the given webhook and message IDs and webhook token and returns its new structure
func (c *Client) EditWebhookMessage(webhookID, messageID discord.Snowflake, webhookToken string, threadID discord.Snowflake, message discord.EditWebhookMessage) (*discord.Webhook, error) {
	queryParams := make(QueryParameters)
	if threadID != 0 {
		queryParams["thread_id"] = threadID.String()
	}
	if len(message.Files) >= 1 {
		contentType, body, err := CreateMultipartBodyWithJSON(message, message.Files)
		if err != nil {
			return nil, err
		}
		return DoRequestAsStructure[discord.Webhook](c, "PATCH", endpoints.WebhookMessage(webhookID, webhookToken, messageID), body, queryParams, 1, WithHeader("Content-Type", contentType))
	} else {
		return DoRequestAsStructure[discord.Webhook](c, "PATCH", endpoints.WebhookMessage(webhookID, webhookToken, messageID), message, queryParams, 1)
	}
}

// DeleteWebhookMessage deletes a message (discord.Message) sent by a webhook (discord.Webhook) for the given webhook and message IDs and webhook token
func (c *Client) DeleteWebhookMessage(webhookID, messageID discord.Snowflake, webhookToken string, threadID discord.Snowflake) error {
	queryParams := make(QueryParameters)
	if threadID != 0 {
		queryParams["thread_id"] = threadID.String()
	}
	return DoEmptyRequest(c, "DELETE", endpoints.WebhookMessage(webhookID, webhookToken, messageID), nil, queryParams, 1)
}
