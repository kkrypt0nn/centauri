package discord

import (
	"encoding/json"
	"fmt"
)

type AuditLog struct {
	ApplicationCommands  []ApplicationCommand  `json:"application_commands"`
	AuditLogEntries      []AuditLogEntry       `json:"audit_log_entries"`
	AutoModerationRules  []AutoModerationRule  `json:"auto_moderation_rules"`
	GuildScheduledEvents []GuildScheduledEvent `json:"guild_scheduled_events"`
	Integrations         []Integration         `json:"integrations"`
	Threads              []Channel             `json:"threads"`
	Users                []User                `json:"users"`
	Webhooks             []Webhook             `json:"webhooks"`
}

type AuditLogEntry struct {
	TargetID   string                   `json:"target_id"`
	Changes    []AuditLogChange         `json:"changes,omitempty"`
	UserID     string                   `json:"user_id,omitempty"`
	ID         string                   `json:"id"`
	ActionType AuditLogActionType       `json:"action_type,omitempty"`
	Options    AuditLogEntryInformation `json:"options,omitempty"`
	Reason     string                   `json:"reason,omitempty"`
}

type AuditLogChange struct {
	NewValue json.RawMessage `json:"new_value,omitempty"`
	OldValue json.RawMessage `json:"old_value,omitempty"`
	Key      string          `json:"key"`
}

func (c *AuditLogChange) ParseValuesTo(old, new interface{}) error {
	if err := json.Unmarshal(c.OldValue, &old); err != nil {
		return fmt.Errorf("failed to unmarshal old value, %s", err)
	}
	if err := json.Unmarshal(c.NewValue, &new); err != nil {
		return fmt.Errorf("failed to unmarshal new value, %s", err)
	}
	return nil
}

type AuditLogActionType int

const (
	AuditLogActionTypeGuildUpdate AuditLogActionType = 1 + iota
	_
	_
	_
	_
	_
	_
	_
	_
	AuditLogActionTypeChannelCreate
	AuditLogActionTypeChannelUpdate
	AuditLogActionTypeChannelDelete
	AuditLogActionTypeChannelOverwriteCreate
	AuditLogActionTypeChannelOverwriteUpdate
	AuditLogActionTypeChannelOverwriteDelete
	_
	_
	_
	_
	AuditLogActionTypeMemberKick
	AuditLogActionTypeMemberPrune
	AuditLogActionTypeMemberBanAdd
	AuditLogActionTypeMemberBanRemove
	AuditLogActionTypeMemberUpdate
	AuditLogActionTypeMemberRoleUpdate
	AuditLogActionTypeMemberMove
	AuditLogActionTypeMemberDisconnect
	AuditLogActionTypeBotAdd
	_
	AuditLogActionTypeRoleCreate
	AuditLogActionTypeRoleUpdate
	AuditLogActionTypeRoleDelete
	_
	_
	_
	_
	_
	_
	_
	AuditLogActionTypeInviteCreate
	AuditLogActionTypeInviteUpdate
	AuditLogActionTypeInviteDelete
	_
	_
	_
	_
	_
	_
	_
	AuditLogActionTypeWebhookCreate
	AuditLogActionTypeWebhookUpdate
	AuditLogActionTypeWebhookDelete
	_
	_
	_
	_
	_
	_
	_
	AuditLogActionTypeEmojiCreate
	AuditLogActionTypeEmojiUpdate
	AuditLogActionTypeEmojiDelete
	_
	_
	_
	_
	_
	_
	_
	_
	_
	AuditLogActionTypeMessageDelete
	AuditLogActionTypeMessageBulkDelete
	AuditLogActionTypeMessagePin
	AuditLogActionTypeMessageUnpin
	_
	_
	_
	_
	AuditLogActionTypeIntegrationCreate
	AuditLogActionTypeIntegrationUpdate
	AuditLogActionTypeIntegrationDelete
	AuditLogActionTypeStanceInstanceCreate
	AuditLogActionTypeStageInstanceUpdate
	AuditLogActionTypeStageInstanceDelete
	_
	_
	_
	_
	AuditLogActionTypeStickerCreate
	AuditLogActionTypeStickerUpdate
	AuditLogActionTypeStickerDelete
	_
	_
	_
	_
	_
	_
	_
	AuditLogActionTypeGuildScheduledEventCreate
	AuditLogActionTypeGuildScheduledEventUpdate
	AuditLogActionTypeGuildScheduledEventDelete
	_
	_
	_
	_
	_
	_
	_
	AuditLogActionTypeThreadCreate
	AuditLogActionTypeThreadUpdate
	AuditLogActionTypeThreadDelete
	_
	_
	_
	_
	_
	_
	_
	_
	AuditLogActionTypeApplicationCommandPermissionUpdate
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	AuditLogActionTypeAutoModerationRuleCreate
	AuditLogActionTypeAutoModerationRuleUpdate
	AuditLogActionTypeAutoModerationRuleDelete
	AuditLogActionTypeAutoModerationBlockMessage
	AuditLogActionTypeAutoModerationFlagToChannel
	AuditLogActionTypeAutoModerationUserCommunicationDisabled
)

type AuditLogEntryInformation struct {
	ApplicationID                 string `json:"application_id,omitempty"`
	AutoModerationRuleName        string `json:"auto_moderation_rule_name,omitempty"`
	AutoModerationRuleTriggerType string `json:"auto_moderation_rule_trigger_type,omitempty"`
	ChannelID                     string `json:"channel_id,omitempty"`
	Count                         string `json:"count,omitempty"`
	DeleteMemberDays              string `json:"delete_member_days,omitempty"`
	ID                            string `json:"id,omitempty"`
	MembersRemoved                string `json:"members_removed,omitempty"`
	MessageID                     string `json:"message_id,omitempty"`
	RoleName                      string `json:"role_name,omitempty"`
	Type                          string `json:"type,omitempty"`
}