package discord

type AutoModerationRule struct {
	ID              string                        `json:"id"`
	GuildID         string                        `json:"guild_id"`
	Name            string                        `json:"name"`
	CreatorID       string                        `json:"creator_id"`
	EventType       AutoModerationEventType       `json:"event_type"`
	TriggerType     AutoModerationTriggerType     `json:"trigger_type"`
	TriggerMetadata AutoModerationTriggerMetadata `json:"trigger_metadata"`
	Actions         []AutoModerationAction        `json:"actions"`
	Enabled         bool                          `json:"enabled"`
	ExemptRoles     []string                      `json:"exempt_roles"`
	ExemptChannels  []string                      `json:"exempt_channels"`
}

type AutoModerationEventType int

const (
	AutoModerationEventTypeMessage AutoModerationEventType = 1 + iota
)

type AutoModerationTriggerType int

const (
	AutoModerationTriggerTypeKeyword AutoModerationTriggerType = 1 + iota
	_
	AutoModerationTriggerTypeSpam
	AutoModerationTriggerTypeKeywordPreset
	AutoModerationTriggerTypeMentionSpam
)

type AutoModerationTriggerMetadata struct {
	KeywordFilter     []string                          `json:"keyword_filter"`
	RegexPatterns     []string                          `json:"regex_patterns"`
	Presets           []AutoModerationKeywordPresetType `json:"presets"`
	AllowList         []string                          `json:"allow_list"`
	MentionTotalLimit int                               `json:"mention_total_limit"`
}

type AutoModerationKeywordPresetType int

const (
	AutoModerationKeywordPresetTypeProfanity AutoModerationKeywordPresetType = 1 + iota
	AutoModerationKeywordPresetTypeSexualContent
	AutoModerationKeywordPresetTypeSlurs
)

type AutoModerationAction struct {
	Type     AutoModerationActionType     `json:"type"`
	Metadata AutoModerationActionMetadata `json:"metadata,omitempty"`
}

type AutoModerationActionType int

const (
	AutoModerationActionTypeBlockMessage AutoModerationActionType = 1 + iota
	AutoModerationActionTypeSendAlertMessage
	AutoModerationActionTypeTimeout
)

type AutoModerationActionMetadata struct {
	ChannelID       string `json:"channel_id,omitempty"`
	DurationSeconds int    `json:"duration_seconds,omitempty"`
	CustomMessage   string `json:"custom_message,omitempty"`
}
