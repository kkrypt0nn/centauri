package gateway

import (
	"github.com/kkrypt0nn/centauri/discord"
	"time"
)

// EventInterface is an interface for getting the interface of the event handler
type EventInterface interface {
	New() interface{}
}

// Hello is the data of the Hello event
// https://discord.com/developers/docs/topics/gateway-events#hello
type Hello struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// Ready is the data of the Ready event
// https://discord.com/developers/docs/topics/gateway-events#ready-ready-event-fields
type Ready struct {
	Version          int                        `json:"v"`
	User             *discord.User              `json:"user"`
	Guilds           []discord.UnavailableGuild `json:"guilds"`
	SessionID        string                     `json:"session_id"`
	ResumeGatewayURL string                     `json:"resume_gateway_url"`
	Shard            [2]int                     `json:"shard"`
	Application      discord.Application        `json:"application"`
}

// ApplicationCommandPermissionsUpdate is the data of the Application Command Permissions Update event
// https://discord.com/developers/docs/topics/gateway-events#application-command-permissions-update
type ApplicationCommandPermissionsUpdate struct {
	discord.GuildApplicationCommandPermissions
}

// AutoModerationRuleCreate is the data of the Auto Moderation Rule Create event
// https://discord.com/developers/docs/topics/gateway-events#auto-moderation-rule-create
type AutoModerationRuleCreate struct {
	discord.AutoModerationRule
}

// AutoModerationRuleUpdate is the data of the Auto Moderation Rule Update event
// https://discord.com/developers/docs/topics/gateway-events#auto-moderation-rule-update
type AutoModerationRuleUpdate struct {
	discord.AutoModerationRule
}

// AutoModerationRuleDelete is the data of the Auto Moderation Rule Delete event
// https://discord.com/developers/docs/topics/gateway-events#auto-moderation-rule-delete
type AutoModerationRuleDelete struct {
	discord.AutoModerationRule
}

// AutoModerationActionExecution is the data of the Auto Moderation Action Execution event
// https://discord.com/developers/docs/topics/gateway-events#auto-moderation-action-execution
type AutoModerationActionExecution struct {
	GuildID              discord.Snowflake                 `json:"guild_id"`
	Action               discord.AutoModerationAction      `json:"action"`
	RuleID               discord.Snowflake                 `json:"rule_id"`
	RuleTriggerType      discord.AutoModerationTriggerType `json:"rule_trigger_type"`
	UserID               discord.Snowflake                 `json:"user_id"`
	ChannelID            discord.Snowflake                 `json:"channel_id,omitempty"`
	MessageID            discord.Snowflake                 `json:"message_id,omitempty"`
	AlertSystemMessageID discord.Snowflake                 `json:"alert_system_message_id,omitempty"`
	Content              string                            `json:"content"`
	MatchedKeyword       string                            `json:"matched_keyword,omitempty"`
	MatchedContent       string                            `json:"matched_content,omitempty"`
}

// ChannelCreate is the data of the Channel Create event
// https://discord.com/developers/docs/topics/gateway-events#channel-create
type ChannelCreate struct {
	discord.Channel
}

// ChannelUpdate is the data of the Channel Update event
// https://discord.com/developers/docs/topics/gateway-events#channel-update
type ChannelUpdate struct {
	discord.Channel
}

// ChannelDelete is the data of the Channel Delete event
// https://discord.com/developers/docs/topics/gateway-events#channel-delete
type ChannelDelete struct {
	discord.Channel
}

// ChannelPinsUpdate is the data of the Channel Pins Update event
// https://discord.com/developers/docs/topics/gateway-events#channel-pins-update
type ChannelPinsUpdate struct {
	GuildID          discord.Snowflake `json:"guild_id,omitempty"`
	ChannelID        discord.Snowflake `json:"channel_id"`
	LastPinTimestamp *time.Time        `json:"last_pin_timestamp,omitempty"`
}

// ThreadCreate is the data of the Thread Create event
// https://discord.com/developers/docs/topics/gateway-events#thread-create
type ThreadCreate struct {
	discord.Channel
}

// ThreadUpdate is the data of the Thread Update event
// https://discord.com/developers/docs/topics/gateway-events#thread-update
type ThreadUpdate struct {
	discord.Channel
}

// ThreadDelete is the data of the Thread Delete event
// https://discord.com/developers/docs/topics/gateway-events#thread-delete
type ThreadDelete struct {
	discord.Channel
}

// ThreadListSync is the data of the Thread List Sync event
// https://discord.com/developers/docs/topics/gateway-events#thread-list-sync
type ThreadListSync struct {
	GuildID    discord.Snowflake       `json:"guild_id"`
	ChannelIDs discord.ArraySnowflakes `json:"channel_ids,omitempty"`
	Threads    []discord.Channel       `json:"threads"`
	Members    []discord.ThreadMember  `json:"members"`
}

// ThreadMemberUpdate is the data of the Thread Member Update event
// https://discord.com/developers/docs/topics/gateway-events#thread-member-update
type ThreadMemberUpdate struct {
	discord.ThreadMember
	GuildID discord.Snowflake `json:"guild_id"`
}

// ThreadMembersUpdate is the data of the Thread Members Update event
// https://discord.com/developers/docs/topics/gateway-events#thread-members-update
type ThreadMembersUpdate struct {
	ID               discord.Snowflake       `json:"id"`
	GuildID          discord.Snowflake       `json:"guild_id"`
	MemberCount      int                     `json:"member_count"`
	AddedMembers     []discord.ThreadMember  `json:"added_members,omitempty"`
	RemovedMemberIDs discord.ArraySnowflakes `json:"removed_member_ids,omitempty"`
}

// GuildCreate is the data of the Guild Create event
// https://discord.com/developers/docs/topics/gateway-events#guild-create
type GuildCreate struct {
	discord.Guild
	JoinedAt    time.Time            `json:"joined_at"`
	Large       bool                 `json:"large"`
	Unavailable bool                 `json:"unavailable,omitempty"`
	MemberCount int                  `json:"member_count"`
	VoiceStates []discord.VoiceState `json:"voice_states"`
	Members     []discord.Member     `json:"members"`
	Channels    []discord.Channel    `json:"channels"`
	Threads     []discord.Channel    `json:"threads"`
	// TODO: Presences
	StageInstances       []discord.StageInstance       `json:"stage_instances"`
	GuildScheduledEvents []discord.GuildScheduledEvent `json:"guild_scheduled_events"`
}

// GuildUpdate is the data of the Guild Update event
// https://discord.com/developers/docs/topics/gateway-events#guild-update
type GuildUpdate struct {
	discord.Guild
}

// GuildDelete is the data of the Guild Delete event
// https://discord.com/developers/docs/topics/gateway-events#guild-delete
type GuildDelete struct {
	discord.UnavailableGuild
}

// GuildAuditLogEntryCreate is the data of the Guild Audit Log Entry Create event
// https://discord.com/developers/docs/topics/gateway-events#guild-audit-log-entry-create
type GuildAuditLogEntryCreate struct {
	discord.AuditLogEntry
}

// GenericGuildBan is the common data of the Guild Ban Add/Remove events
// https://discord.com/developers/docs/topics/gateway-events#guild-ban-add
// https://discord.com/developers/docs/topics/gateway-events#guild-ban-remove
type GenericGuildBan struct {
	GuildID discord.Snowflake `json:"guild_id"`
	User    discord.User      `json:"user"`
}

// GuildBanAdd is the data of the Guild Ban Add event
// https://discord.com/developers/docs/topics/gateway-events#guild-ban-add
type GuildBanAdd struct {
	GenericGuildBan
}

// GuildBanRemove is the data of the Guild Ban Remove event
// https://discord.com/developers/docs/topics/gateway-events#guild-ban-remove
type GuildBanRemove struct {
	GenericGuildBan
}

// GuildEmojisUpdate is the data of the Guild Emojis Update event
// https://discord.com/developers/docs/topics/gateway-events#guild-emojis-update
type GuildEmojisUpdate struct {
	GuildID discord.Snowflake `json:"guild_id"`
	Emojis  []discord.Emoji   `json:"emojis"`
}

// GuildStickersUpdate is the data of the Guild Stickers Update event
// https://discord.com/developers/docs/topics/gateway-events#guild-stickers-update
type GuildStickersUpdate struct {
	GuildID  discord.Snowflake `json:"guild_id"`
	Stickers []discord.Sticker `json:"stickers"`
}

// GuildIntegrationsUpdate is the data of the Guild Integrations Update event
// https://discord.com/developers/docs/topics/gateway-events#guild-integrations-update
type GuildIntegrationsUpdate struct {
	GuildID discord.Snowflake `json:"guild_id"`
}

// GuildMemberAdd is the data of the Guild Member Add event
// https://discord.com/developers/docs/topics/gateway-events#guild-member-add
type GuildMemberAdd struct {
	discord.Member
}

// GuildMemberRemove is the data of the Guild Member Remove event
// https://discord.com/developers/docs/topics/gateway-events#guild-member-remove
type GuildMemberRemove struct {
	GuildID discord.Snowflake `json:"guild_id"`
	User    discord.User      `json:"user"`
}

// GuildMemberUpdate is the data of the Guild Member Update event
// https://discord.com/developers/docs/topics/gateway-events#guild-member-update
type GuildMemberUpdate struct {
	GuildID                    discord.Snowflake       `json:"guild_id"`
	Roles                      discord.ArraySnowflakes `json:"roles"`
	User                       discord.User            `json:"user"`
	Nick                       string                  `json:"nick,omitempty"`
	Avatar                     string                  `json:"avatar,omitempty"`
	JoinedAt                   *time.Time              `json:"joined_at,omitempty"`
	PremiumSince               *time.Time              `json:"premium_since,omitempty"`
	Deaf                       bool                    `json:"deaf"`
	Mute                       bool                    `json:"mute"`
	Pending                    bool                    `json:"pending"`
	CommunicationDisabledUntil *time.Time              `json:"communication_disabled_until,omitempty"`
}

// GuildMembersChunk is the data of the Guild Members Chunk event
// https://discord.com/developers/docs/topics/gateway-events#guild-members-chunk
type GuildMembersChunk struct {
	GuildID    discord.Snowflake       `json:"guild_id"`
	Members    []discord.Member        `json:"members"`
	ChunkIndex int                     `json:"chunk_index"`
	ChunkCount int                     `json:"chunk_count"`
	NotFound   discord.ArraySnowflakes `json:"not_found"`
	// TODO: Presences
	Nonce string `json:"nonce"`
}

// GenericGuildRole is the common data of the Guild Role Create/Update events
// https://discord.com/developers/docs/topics/gateway-events#guild-role-create
// https://discord.com/developers/docs/topics/gateway-events#guild-role-update
type GenericGuildRole struct {
	GuildID discord.Snowflake `json:"guild_id"`
	Role    discord.Role      `json:"role"`
}

// GuildRoleCreate is the data of the Guild Role Create event
// https://discord.com/developers/docs/topics/gateway-events#guild-role-create
type GuildRoleCreate struct {
	GenericGuildRole
}

// GuildRoleUpdate is the data of the Guild Role Update event
// https://discord.com/developers/docs/topics/gateway-events#guild-role-update
type GuildRoleUpdate struct {
	GenericGuildRole
}

// GuildRoleDelete is the data of the Guild Role Delete event
// https://discord.com/developers/docs/topics/gateway-events#guild-role-delete
type GuildRoleDelete struct {
	GuildID discord.Snowflake `json:"guild_id"`
	RoleID  discord.Snowflake `json:"role_id"`
}

// GuildScheduledEventCreate is the data of the Guild Scheduled Event Create event
// https://discord.com/developers/docs/topics/gateway-events#guild-scheduled-event-create
type GuildScheduledEventCreate struct {
	discord.GuildScheduledEvent
}

// GuildScheduledEventUpdate is the data of the Guild Scheduled Event Update event
// https://discord.com/developers/docs/topics/gateway-events#guild-scheduled-event-update
type GuildScheduledEventUpdate struct {
	discord.GuildScheduledEvent
}

// GuildScheduledEventDelete is the data of the Guild Scheduled Event Delete event
// https://discord.com/developers/docs/topics/gateway-events#guild-scheduled-event-delete
type GuildScheduledEventDelete struct {
	discord.GuildScheduledEvent
}

// GenericGuildScheduledEventUser is the common data of the Guild Scheduled Event User Add/Remove events
// https://discord.com/developers/docs/topics/gateway-events#guild-scheduled-event-user-add
// https://discord.com/developers/docs/topics/gateway-events#guild-scheduled-event-user-remove
type GenericGuildScheduledEventUser struct {
	GuildScheduledEventID discord.Snowflake `json:"guild_scheduled_event_id"`
	UserID                discord.Snowflake `json:"user_id"`
	GuildID               discord.Snowflake `json:"guild_id"`
}

// GuildScheduledEventUserAdd is the data of the Guild Scheduled Event User Add event
// https://discord.com/developers/docs/topics/gateway-events#guild-scheduled-event-user-add
type GuildScheduledEventUserAdd struct {
	GenericGuildScheduledEventUser
}

// GuildScheduledEventUserRemove is the data of the Guild Scheduled Event User Remove event
// https://discord.com/developers/docs/topics/gateway-events#guild-scheduled-event-user-remove
type GuildScheduledEventUserRemove struct {
	GenericGuildScheduledEventUser
}

// GenericIntegration is the common data of the Integration Create/Update events
// https://discord.com/developers/docs/topics/gateway-events#integration-create
// https://discord.com/developers/docs/topics/gateway-events#integration-update
type GenericIntegration struct {
	discord.Integration
	GuildID discord.Snowflake `json:"guild_id"`
}

// IntegrationCreate is the data of the Integration Create event
// https://discord.com/developers/docs/topics/gateway-events#integration-create
type IntegrationCreate struct {
	GenericIntegration
}

// IntegrationUpdate is the data of the Integration Update event
// https://discord.com/developers/docs/topics/gateway-events#integration-update
type IntegrationUpdate struct {
	GenericIntegration
}

// IntegrationDelete is the data of the Integration Delete event
// https://discord.com/developers/docs/topics/gateway-events#integration-delete
type IntegrationDelete struct {
	ID            discord.Snowflake `json:"id"`
	GuildID       discord.Snowflake `json:"guild_id"`
	ApplicationID discord.Snowflake `json:"application_id"`
}

// InteractionCreate is the data of the Interaction Create event
// https://discord.com/developers/docs/topics/gateway-events#interaction-create
type InteractionCreate struct {
	discord.Interaction
}

// InviteCreate is the data of the Invite Create event
// https://discord.com/developers/docs/topics/gateway-events#invite-create
type InviteCreate struct {
	ChannelID         discord.Snowflake        `json:"channel_id"`
	Code              string                   `json:"code"`
	CreatedAt         time.Time                `json:"created_at"`
	GuildID           discord.Snowflake        `json:"guild_id,omitempty"`
	Inviter           discord.User             `json:"inviter"`
	MaxAge            int                      `json:"max_age"`
	MaxUses           int                      `json:"max_uses"`
	TargetType        discord.InviteTargetType `json:"target_type"`
	TargetUser        *discord.User            `json:"target_user,omitempty"`
	TargetApplication *discord.Application     `json:"target_application,omitempty"`
	Temporary         bool                     `json:"temporary"`
	Uses              int                      `json:"uses"`
}

// InviteDelete is the data of the Invite Delete event
// https://discord.com/developers/docs/topics/gateway-events#invite-delete
type InviteDelete struct {
	ChannelID discord.Snowflake `json:"channel_id"`
	GuildID   discord.Snowflake `json:"guild_id,omitempty"`
	Code      string            `json:"code"`
}

// GenericMessage is the common data of the Message Create/Update events
// https://discord.com/developers/docs/topics/gateway-events#message-create
// https://discord.com/developers/docs/topics/gateway-events#message-update
type GenericMessage struct {
	discord.Message
	Member *discord.Member `json:"member,omitempty"`
}

// MessageCreate is the data of the Message Create event
// https://discord.com/developers/docs/topics/gateway-events#message-create
type MessageCreate struct {
	GenericMessage
}

// MessageUpdate is the data of the Message Update event
// https://discord.com/developers/docs/topics/gateway-events#message-update
type MessageUpdate struct {
	GenericMessage
}

// MessageDelete is the data of the Message Delete event
// https://discord.com/developers/docs/topics/gateway-events#message-delete
type MessageDelete struct {
	ID        discord.Snowflake `json:"id"`
	ChannelID discord.Snowflake `json:"channel_id"`
	GuildID   discord.Snowflake `json:"guild_id,omitempty"`
}

// MessageDeleteBulk is the data of the Message Delete Bulk event
// https://discord.com/developers/docs/topics/gateway-events#message-delete-bulk
type MessageDeleteBulk struct {
	IDs       discord.ArraySnowflakes `json:"ids"`
	ChannelID discord.Snowflake       `json:"channel_id"`
	GuildID   discord.Snowflake       `json:"guild_id,omitempty"`
}

// GenericMessageReaction is the common data of the Message Reaction Add/Remove/Remove All/Remove Emoji events
// https://discord.com/developers/docs/topics/gateway-events#message-reaction-add
// https://discord.com/developers/docs/topics/gateway-events#message-reaction-remove
// https://discord.com/developers/docs/topics/gateway-events#message-reaction-remove-all
// https://discord.com/developers/docs/topics/gateway-events#message-reaction-remove-emoji
type GenericMessageReaction struct {
	ChannelID discord.Snowflake `json:"channel_id"`
	MessageID discord.Snowflake `json:"message_id"`
	GuildID   discord.Snowflake `json:"guild_id,omitempty"`
}

// MessageReactionAdd is the data of the Message Reaction Add event
// https://discord.com/developers/docs/topics/gateway-events#message-reaction-add
type MessageReactionAdd struct {
	GenericMessageReaction
	UserID discord.Snowflake `json:"user_id"`
	Member *discord.Member   `json:"member,omitempty"`
	Emoji  discord.Emoji     `json:"emoji"`
}

// MessageReactionRemove is the data of the Message Reaction Remove event
// https://discord.com/developers/docs/topics/gateway-events#message-reaction-remove
type MessageReactionRemove struct {
	GenericMessageReaction
	UserID discord.Snowflake `json:"user_id"`
	Emoji  discord.Emoji     `json:"emoji"`
}

// MessageReactionRemoveAll is the data of the Message Reaction Remove All event
// https://discord.com/developers/docs/topics/gateway-events#message-reaction-remove-all
type MessageReactionRemoveAll struct {
	GenericMessageReaction
}

// MessageReactionRemoveEmoji is the data of the Message Reaction Remove Emoji event
// https://discord.com/developers/docs/topics/gateway-events#message-reaction-remove-emoji
type MessageReactionRemoveEmoji struct {
	GenericMessageReaction
	Emoji discord.Emoji `json:"emoji"`
}

// PresenceUpdate is the data of the Presence Update event
// https://discord.com/developers/docs/topics/gateway-events#presence-update
type PresenceUpdate struct {
	User         discord.User         `json:"user"`
	GuildID      discord.Snowflake    `json:"guild_id"`
	Status       string               `json:"status"`
	Activities   []discord.Activity   `json:"activities"`
	ClientStatus discord.ClientStatus `json:"client_status"`
}

// StageInstanceCreate is the data of the Stage Instance Create event
// https://discord.com/developers/docs/topics/gateway-events#stage-instance-create
type StageInstanceCreate struct {
	discord.StageInstance
}

// StageInstanceUpdate is the data of the Stage Instance Update event
// https://discord.com/developers/docs/topics/gateway-events#stage-instance-update
type StageInstanceUpdate struct {
	discord.StageInstance
}

// StageInstanceDelete is the data of the Stage Instance Delete event
// https://discord.com/developers/docs/topics/gateway-events#stage-instance-delete
type StageInstanceDelete struct {
	discord.StageInstance
}

// TypingStart is the data of the Typing Start event
// https://discord.com/developers/docs/topics/gateway-events#typing-start
type TypingStart struct {
	ChannelID discord.Snowflake `json:"channel_id"`
	GuildID   discord.Snowflake `json:"guild_id,omitempty"`
	UserID    discord.Snowflake `json:"user_id"`
	Timestamp int               `json:"timestamp"`
	Member    *discord.Member   `json:"member,omitempty"`
}

// UserUpdate is the data of the User Update event
// https://discord.com/developers/docs/topics/gateway-events#user-update
type UserUpdate struct {
	discord.User
}

// VoiceStateUpdate is the data of the Voice State Update event
// https://discord.com/developers/docs/topics/gateway-events#voice-state-update
type VoiceStateUpdate struct {
	discord.VoiceState
}

// VoiceServerUpdate is the data of the Voice Server Update event
// https://discord.com/developers/docs/topics/gateway-events#voice-server-update
type VoiceServerUpdate struct {
	Token    string            `json:"token"`
	GuildID  discord.Snowflake `json:"guild_id"`
	Endpoint string            `json:"endpoint"`
}

// WebhooksUpdate is the data of the Webhooks Update event
// https://discord.com/developers/docs/topics/gateway-events#webhooks-update
type WebhooksUpdate struct {
	GuildID   discord.Snowflake `json:"guild_id"`
	ChannelID discord.Snowflake `json:"channel_id"`
}
