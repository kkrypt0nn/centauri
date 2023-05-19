package discord

// Webhook represents a low-effort way to post messages to channels (discord.Channel). A bot authentication is not required for using them
// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-structure
type Webhook struct {
	ID            string        `json:"id"`
	Type          WebhookType   `json:"type"`
	GuildID       string        `json:"guild_id,omitempty"`
	ChannelID     string        `json:"channel_id"`
	User          *User         `json:"user,omitempty"`
	Name          string        `json:"name"`
	Avatar        string        `json:"avatar"`
	Token         string        `json:"token,omitempty"`
	ApplicationID string        `json:"application_id"`
	SourceGuild   *PartialGuild `json:"source_guild,omitempty"`
	SourceChannel *Channel      `json:"source_channel,omitempty"`
	URL           string        `json:"url,omitempty"`
}

// WebhookType represents the type of webhook (discord.Webhook)
// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-types
type WebhookType int

const (
	WebhookTypeIncoming WebhookType = 1 + iota
	WebhookTypeChannelFollower
	WebhookTypeApplication
)
