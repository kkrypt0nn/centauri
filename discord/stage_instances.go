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

// CreateStageInstance represents the payload to send to Discord to create a new stage instance (discord.StageInstance)
// https://discord.com/developers/docs/resources/stage-instance#create-stage-instance-json-params
type CreateStageInstance struct {
	ChannelID             string                     `json:"channel_id"`
	Topic                 string                     `json:"topic"`
	PrivacyLevel          *StageInstancePrivacyLevel `json:"privacy_level,omitempty"`
	SendStartNotification *bool                      `json:"send_start_notification,omitempty"`

	AuditLogReason string `json:"-"`
}

// ModifyStageInstance represents the payload to send to Discord to modify an existing stage instance (discord.StageInstance)
// https://discord.com/developers/docs/resources/stage-instance#modify-stage-instance-json-params
type ModifyStageInstance struct {
	Topic        *string                    `json:"topic,omitempty"`
	PrivacyLevel *StageInstancePrivacyLevel `json:"privacy_level,omitempty"`

	AuditLogReason string `json:"-"`
}
