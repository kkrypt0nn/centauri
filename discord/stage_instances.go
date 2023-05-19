package discord

// StageInstance represents a live stage
// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-stage-instance-structure
type StageInstance struct {
	ID                    string                    `json:"id"`
	GuildID               string                    `json:"guild_id"`
	ChannelID             string                    `json:"channel_id"`
	Topic                 string                    `json:"topic"`
	PrivacyLevel          StageInstancePrivacyLevel `json:"privacy_level"`
	DiscoverableDisabled  bool                      `json:"discoverable_disabled"`
	GuildScheduledEventID string                    `json:"guild_scheduled_event_id"`
}

// StageInstancePrivacyLevel represents the privacy level of the stage instance (discord.StageInstance)
// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-privacy-level
type StageInstancePrivacyLevel int

const (
	StageInstancePrivacyLevelPublic StageInstancePrivacyLevel = 1 + iota
	StageInstancePrivacyLevelGuildOnly
)
