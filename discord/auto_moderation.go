package discord

// AutoModerationRule is a rule that trigger based on some criteria
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-auto-moderation-rule-structure
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

// AutoModerationEventType represents in what event context a rule (discord.AutoModerationRule) should be checked
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-event-types
type AutoModerationEventType int

const (
	AutoModerationEventTypeMessageSend AutoModerationEventType = 1 + iota
)

// AutoModerationTriggerType represents the type of content which can trigger the rule (discord.AutoModerationRule)
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-trigger-types
type AutoModerationTriggerType int

const (
	AutoModerationTriggerTypeKeyword AutoModerationTriggerType = 1 + iota
	_
	AutoModerationTriggerTypeSpam
	AutoModerationTriggerTypeKeywordPreset
	AutoModerationTriggerTypeMentionSpam
)

// AutoModerationTriggerMetadata represents additional data used to determine whether a rule (discord.AutoModerationRule) should be triggered
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-trigger-metadata
type AutoModerationTriggerMetadata struct {
	KeywordFilter     []string                          `json:"keyword_filter,omitempty"`
	RegexPatterns     []string                          `json:"regex_patterns,omitempty"`
	Presets           []AutoModerationKeywordPresetType `json:"presets,omitempty"`
	AllowList         []string                          `json:"allow_list,omitempty"`
	MentionTotalLimit int                               `json:"mention_total_limit,omitempty"`
}

// AutoModerationKeywordPresetType represents the internally pre-defined word sets which will be searched for in content
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-keyword-preset-types
type AutoModerationKeywordPresetType int

const (
	AutoModerationKeywordPresetTypeProfanity AutoModerationKeywordPresetType = 1 + iota
	AutoModerationKeywordPresetTypeSexualContent
	AutoModerationKeywordPresetTypeSlurs
)

// AutoModerationAction represents the action which will execute whenever a rule (discord.AutoModerationRule) is triggered
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-action-object-auto-moderation-action-structure
type AutoModerationAction struct {
	Type     AutoModerationActionType      `json:"type"`
	Metadata *AutoModerationActionMetadata `json:"metadata,omitempty"`
}

// AutoModerationActionType represents the type of action (discord.AutoModerationAction) that will be performed
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-action-object-action-types
type AutoModerationActionType int

const (
	AutoModerationActionTypeBlockMessage AutoModerationActionType = 1 + iota
	AutoModerationActionTypeSendAlertMessage
	AutoModerationActionTypeTimeout
)

// AutoModerationActionMetadata represents additional data used when an action (discord.AutoModerationAction) is executed. Different fields are relevant based on the value of action type (discord.AutoModerationActionType)
// https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-action-object-action-metadata
type AutoModerationActionMetadata struct {
	ChannelID       string `json:"channel_id,omitempty"`
	DurationSeconds int    `json:"duration_seconds,omitempty"`
	CustomMessage   string `json:"custom_message,omitempty"`
}

// CreateAutoModerationRule represents the payload to send to Discord to create a new auto moderation rule (discord.AutoModerationRule)
// https://discord.com/developers/docs/resources/auto-moderation#create-auto-moderation-rule-json-params
type CreateAutoModerationRule struct {
	Name            string                         `json:"name"`
	EventType       AutoModerationEventType        `json:"event_type"`
	TriggerType     AutoModerationTriggerType      `json:"trigger_type"`
	TriggerMetadata *AutoModerationTriggerMetadata `json:"trigger_metadata,omitempty"`
	Actions         []AutoModerationAction         `json:"actions"`
	Enabled         *bool                          `json:"enabled,omitempty"`
	ExemptRoles     []string                       `json:"exempt_roles,omitempty"`
	ExemptChannels  []string                       `json:"exempt_channels,omitempty"`

	AuditLogReason string `json:"-"`
}
