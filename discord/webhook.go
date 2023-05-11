package discord

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

type WebhookType int

const (
	WebhookTypeIncoming WebhookType = 1 + iota
	WebhookTypeChannelFollower
	WebhookTypeApplication
)
