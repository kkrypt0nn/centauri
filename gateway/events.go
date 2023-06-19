package gateway

import (
	"encoding/json"
	"io"
)

// EventType represents the event type
type EventType string

const (
	EventTypeReady                               EventType = "READY"
	EventTypeResumed                             EventType = "RESUMED"
	EventTypeApplicationCommandPermissionsUpdate EventType = "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
	EventTypeAutoModerationRuleCreate            EventType = "AUTO_MODERATION_RULE_CREATE"
	EventTypeAutoModerationRuleUpdate            EventType = "AUTO_MODERATION_RULE_UPDATE"
	EventTypeAutoModerationRuleDelete            EventType = "AUTO_MODERATION_RULE_DELETE"
	EventTypeAutoModerationActionExecution       EventType = "AUTO_MODERATION_ACTION_EXECUTION"
	EventTypeChannelCreate                       EventType = "CHANNEL_CREATE"
	EventTypeChannelUpdate                       EventType = "CHANNEL_UPDATE"
	EventTypeChannelDelete                       EventType = "CHANNEL_DELETE"
	EventTypeChannelPinsUpdate                   EventType = "CHANNEL_PINS_UPDATE"
	EventTypeThreadCreate                        EventType = "THREAD_CREATE"
	EventTypeThreadUpdate                        EventType = "THREAD_UPDATE"
	EventTypeThreadDelete                        EventType = "THREAD_DELETE"
	EventTypeThreadListSync                      EventType = "THREAD_LIST_SYNC"
	EventTypeThreadMemberUpdate                  EventType = "THREAD_MEMBER_UPDATE"
	EventTypeThreadMembersUpdate                 EventType = "THREAD_MEMBERS_UPDATE"
	EventTypeGuildCreate                         EventType = "GUILD_CREATE"
	EventTypeGuildUpdate                         EventType = "GUILD_UPDATE"
	EventTypeGuildDelete                         EventType = "GUILD_DELETE"
	EventTypeGuildAuditLogEntryCreate            EventType = "GUILD_AUDIT_LOG_ENTRY_CREATE"
	EventTypeGuildBanAdd                         EventType = "GUILD_BAN_ADD"
	EventTypeGuildBanRemove                      EventType = "GUILD_BAN_REMOVE"
	EventTypeGuildEmojisUpdate                   EventType = "GUILD_EMOJIS_UPDATE"
	EventTypeGuildStickersUpdate                 EventType = "GUILD_STICKERS_UPDATE"
	EventTypeGuildIntegrationsUpdate             EventType = "GUILD_INTEGRATIONS_UPDATE"
	EventTypeGuildMemberAdd                      EventType = "GUILD_MEMBER_ADD"
	EventTypeGuildMemberRemove                   EventType = "GUILD_MEMBER_REMOVE"
	EventTypeGuildMemberUpdate                   EventType = "GUILD_MEMBER_UPDATE"
	EventTypeGuildMembersChunk                   EventType = "GUILD_MEMBERS_CHUNK"
	EventTypeGuildRoleCreate                     EventType = "GUILD_ROLE_CREATE"
	EventTypeGuildRoleUpdate                     EventType = "GUILD_ROLE_UPDATE"
	EventTypeGuildRoleDelete                     EventType = "GUILD_ROLE_DELETE"
	EventTypeGuildScheduledEventCreate           EventType = "GUILD_SCHEDULED_EVENT_CREATE"
	EventTypeGuildScheduledEventUpdate           EventType = "GUILD_SCHEDULED_EVENT_UPDATE"
	EventTypeGuildScheduledEventDelete           EventType = "GUILD_SCHEDULED_EVENT_DELETE"
	EventTypeGuildScheduledEventUserAdd          EventType = "GUILD_SCHEDULED_EVENT_USER_ADD"
	EventTypeGuildScheduledEventUserRemove       EventType = "GUILD_SCHEDULED_EVENT_USER_REMOVE"
	EventTypeIntegrationCreate                   EventType = "INTEGRATION_CREATE"
	EventTypeIntegrationUpdate                   EventType = "INTEGRATION_UPDATE"
	EventTypeIntegrationDelete                   EventType = "INTEGRATION_DELETE"
	EventTypeInteractionCreate                   EventType = "INTERACTION_CREATE"
	EventTypeInviteCreate                        EventType = "INVITE_CREATE"
	EventTypeInviteDelete                        EventType = "INVITE_DELETE"
	EventTypeMessageCreate                       EventType = "MESSAGE_CREATE"
	EventTypeMessageUpdate                       EventType = "MESSAGE_UPDATE"
	EventTypeMessageDelete                       EventType = "MESSAGE_DELETE"
	EventTypeMessageDeleteBulk                   EventType = "MESSAGE_DELETE_BULK"
	EventTypeMessageReactionAdd                  EventType = "MESSAGE_REACTION_ADD"
	EventTypeMessageReactionRemove               EventType = "MESSAGE_REACTION_REMOVE"
	EventTypeMessageReactionRemoveAll            EventType = "MESSAGE_REACTION_REMOVE_ALL"
	EventTypeMessageReactionRemoveEmoji          EventType = "MESSAGE_REACTION_REMOVE_EMOJI"
	EventTypePresenceUpdate                      EventType = "PRESENCE_UPDATE"
	EventTypeStageInstanceCreate                 EventType = "STAGE_INSTANCE_CREATE"
	EventTypeStageInstanceUpdate                 EventType = "STAGE_INSTANCE_UPDATE"
	EventTypeStageInstanceDelete                 EventType = "STAGE_INSTANCE_DELETE"
	EventTypeTypingStart                         EventType = "TYPING_START"
	EventTypeUserUpdate                          EventType = "USER_UPDATE"
	EventTypeVoiceStateUpdate                    EventType = "VOICE_STATE_UPDATE"
	EventTypeVoiceServerUpdate                   EventType = "VOICE_SERVER_UPDATE"
	EventTypeWebhooksUpdate                      EventType = "WEBHOOKS_UPDATE"
)

// Event is the structure of a gateway event
// https://discord.com/developers/docs/topics/gateway#gateway-events-example-gateway-event
type Event struct {
	OpCode   OpCode          `json:"op"`
	Sequence int64           `json:"s"`
	Type     string          `json:"t"`
	Data     json.RawMessage `json:"d"`

	Struct interface{} `json:"-"`
}

// NewEvent parses a websocket message and returns a new event
func NewEvent(message io.Reader) (*Event, error) {
	var event *Event
	err := json.NewDecoder(message).Decode(&event)
	if err != nil {
		return nil, err
	}
	return event, err
}
