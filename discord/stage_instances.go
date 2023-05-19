package discord

type StageInstance struct {
	ID                    string                    `json:"id"`
	GuildID               string                    `json:"guild_id"`
	ChannelID             string                    `json:"channel_id"`
	Topic                 string                    `json:"topic"`
	PrivacyLevel          StageInstancePrivacyLevel `json:"privacy_level"`
	DiscoverableDisabled  bool                      `json:"discoverable_disabled"`
	GuildScheduledEventID string                    `json:"guild_scheduled_event_id"`
}

type StageInstancePrivacyLevel int

const (
	StageInstancePrivacyLevelPublic StageInstancePrivacyLevel = 1 + iota
	StageInstancePrivacyLevelGuildOnly
)
