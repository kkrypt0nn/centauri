package discord

import "time"

// Webhook represents a low-effort way to post messages to channels (discord.Channel). A bot authentication is not required for using them
// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-structure
type Webhook struct {
	ID            Snowflake     `json:"id"`
	Type          WebhookType   `json:"type"`
	GuildID       Snowflake     `json:"guild_id,omitempty"`
	ChannelID     Snowflake     `json:"channel_id"`
	User          *User         `json:"user,omitempty"`
	Name          string        `json:"name"`
	Avatar        string        `json:"avatar"`
	Token         string        `json:"token,omitempty"`
	ApplicationID Snowflake     `json:"application_id"`
	SourceGuild   *PartialGuild `json:"source_guild,omitempty"`
	SourceChannel *Channel      `json:"source_channel,omitempty"`
	URL           string        `json:"url,omitempty"`
}

// CreatedAt returns the creation date of the webhook (discord.Webhook)
func (w *Webhook) CreatedAt() time.Time {
	return w.ID.CreatedAt()
}

// WebhookType represents the type of webhook (discord.Webhook)
// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-types
type WebhookType int

const (
	WebhookTypeIncoming WebhookType = 1 + iota
	WebhookTypeChannelFollower
	WebhookTypeApplication
)

// CreateWebhook represents the payload to send to Discord to create a new webhook (discord.Webhook)
// https://discord.com/developers/docs/resources/webhook#create-webhook-json-params
type CreateWebhook struct {
	Name   string  `json:"name"`
	Avatar *string `json:"avatar,omitempty"`

	AuditLogReason string `json:"-"`
}

// ModifyWebhook represents the payload to send to Discord to modify an existing webhook (discord.Webhook)
// https://discord.com/developers/docs/resources/webhook#modify-webhook-json-params
type ModifyWebhook struct {
	Name      *string    `json:"name,omitempty"`
	Avatar    *string    `json:"avatar,omitempty"`
	ChannelID *Snowflake `json:"channel_id,omitempty"`

	AuditLogReason string `json:"-"`
}

// ModifyWebhookWithToken represents the payload to send to Discord to modify an existing webhook (discord.Webhook) with token only
// https://discord.com/developers/docs/resources/webhook#modify-webhook-json-params
type ModifyWebhookWithToken struct {
	Name   *string `json:"name,omitempty"`
	Avatar *string `json:"avatar,omitempty"`

	AuditLogReason string `json:"-"`
}

// ExecuteWebhook represents the payload to send to Discord to execute a webhook (discord.Webhook) aka send a message
// https://discord.com/developers/docs/resources/webhook#execute-webhook-jsonform-params
type ExecuteWebhook struct {
	Content         *string          `json:"content,omitempty"`
	Username        *string          `json:"username,omitempty"`
	AvatarURL       *string          `json:"avatar_url,omitempty"`
	TTS             *bool            `json:"tts,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Components      []Component      `json:"components,omitempty"`
	Flags           *MessageFlags    `json:"flags,omitempty"`
	ThreadName      *string          `json:"thread_name,omitempty"`
	Attachments     []AttachmentSend `json:"attachments,omitempty"`

	Files []File `json:"-"`
}

// EditWebhookMessage represents the payload to send to Discord to edit a message (discord.Message) sent by a webhook (discord.Webhook)
// https://discord.com/developers/docs/resources/webhook#edit-webhook-message-jsonform-params
type EditWebhookMessage struct {
	Content         *string          `json:"content,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Components      []Component      `json:"components,omitempty"`
	Attachments     []AttachmentSend `json:"attachments,omitempty"`

	Files []File `json:"-"`
}
