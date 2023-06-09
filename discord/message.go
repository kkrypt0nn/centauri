package discord

import (
	"encoding/json"
	"fmt"
	"github.com/kkrypt0nn/centauri/utils/flags"
	"io"
	"time"
)

// Message represents a message sent in a channel (discord.Channel) within Discord
// https://discord.com/developers/docs/resources/channel#message-object-message-structure
type Message struct {
	ID                   Snowflake             `json:"id"`
	ChannelID            Snowflake             `json:"channel_id"`
	Author               *User                 `json:"author"`
	Content              string                `json:"content"`
	Timestamp            time.Time             `json:"timestamp"`
	EditedTimestamp      *time.Time            `json:"edited_timestamp,omitempty"`
	TTS                  bool                  `json:"tts"`
	MentionEveryone      bool                  `json:"mention_everyone"`
	Mentions             []User                `json:"mentions"`
	MentionRoles         []Role                `json:"mention_roles"`
	MentionChannels      []Channel             `json:"mention_channels,omitempty"`
	Attachments          []Attachment          `json:"attachments"`
	Embeds               []Embed               `json:"embeds"`
	Reactions            []Reaction            `json:"reactions,omitempty"`
	Nonce                string                `json:"nonce,omitempty"`
	Pinned               bool                  `json:"pinned"`
	WebhookID            Snowflake             `json:"webhook_id,omitempty"`
	Type                 MessageType           `json:"type"`
	Activity             *MessageActivity      `json:"activity,omitempty"`
	Application          *Application          `json:"application,omitempty"`
	ApplicationID        Snowflake             `json:"application_id,omitempty"`
	MessageReference     *MessageReference     `json:"message_reference,omitempty"`
	Flags                MessageFlags          `json:"flags,omitempty"`
	ReferencedMessage    *Message              `json:"referenced_message,omitempty"`
	Interaction          *MessageInteraction   `json:"interaction,omitempty"`
	Thread               *Channel              `json:"thread,omitempty"`
	Components           []Component           `json:"components,omitempty"`
	StickerItems         []StickerItem         `json:"sticker_items,omitempty"`
	Stickers             []Sticker             `json:"stickers,omitempty"`
	Position             int                   `json:"position"`
	RoleSubscriptionData *RoleSubscriptionData `json:"role_subscription_data,omitempty"`

	// Usually that's set by Centauri
	GuildID string `json:"guild_id"`
}

// CreatedAt returns the creation time of the message (discord.Message)
func (m *Message) CreatedAt() time.Time {
	return m.ID.CreatedAt()
}

// URL returns the URL of the message (discord.Message)
func (m *Message) URL() string {
	guildID := m.GuildID
	if m.GuildID == "" {
		guildID = "@me"
	}
	return fmt.Sprintf("https://discord.com/channels/%s/%d/%d", guildID, m.ChannelID, m.ID)
}

func (m *Message) UnmarshalJSON(b []byte) error {
	type message Message
	var source struct {
		message
		UnidentifiedComponents []UnidentifiedComponent `json:"components"`
	}
	if err := json.Unmarshal(b, &source); err != nil {
		return fmt.Errorf("failed unmarshalling: %s", err)
	}

	*m = Message(source.message)

	if len(source.UnidentifiedComponents) == 0 {
		m.Components = nil
		return nil
	}

	m.Components = make([]Component, len(source.UnidentifiedComponents))
	for i, unidentifiedComponents := range source.UnidentifiedComponents {
		m.Components[i] = unidentifiedComponents.data
	}

	return nil
}

// AllowedMentions represents the allowed mentions for a message (discord.Message)
// https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mentions-structure
type AllowedMentions struct {
	Parse       []AllowedMentionType `json:"parse,omitempty"`
	Roles       ArraySnowflakes      `json:"roles,omitempty"`
	Users       ArraySnowflakes      `json:"users,omitempty"`
	RepliedUser *bool                `json:"replied_user,omitempty"`
}

// AllowedMentionType represents the allowed mention types for an allowed mentions structure (discord.AllowedMentions)
// https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mention-types
type AllowedMentionType string

const (
	AllowedMentionTypeRoleMentions     AllowedMentionType = "roles"
	AllowedMentionTypeUserMentions     AllowedMentionType = "users"
	AllowedMentionTypeEveryoneMentions AllowedMentionType = "everyone"
)

// Attachment represents an attached file in a message (discord.Message)
// https://discord.com/developers/docs/resources/channel#attachment-object-attachment-structure
type Attachment struct {
	ID           Snowflake `json:"id"`
	Filename     string    `json:"filename"`
	Description  string    `json:"description,omitempty"`
	ContentType  string    `json:"content_type,omitempty"`
	Size         int       `json:"size,omitempty"`
	URL          string    `json:"url,omitempty"`
	ProxyURL     string    `json:"proxy_url,omitempty"`
	Height       int       `json:"height,omitempty"`
	Width        int       `json:"width,omitempty"`
	Ephemeral    bool      `json:"ephemeral,omitempty"`
	DurationSecs float64   `json:"duration_secs,omitempty"`
	Waveform     string    `json:"waveform,omitempty"`
}

// CreatedAt returns the creation time of the attachment (discord.Attachment)
func (a *Attachment) CreatedAt() time.Time {
	return a.ID.CreatedAt()
}

// AttachmentSend represents an attachment that is attached to a message to send, these are the only fields necessary when sending an attachment
type AttachmentSend struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// MessageType represents the type of message (discord.Message)
// https://discord.com/developers/docs/resources/channel#message-object-message-types
type MessageType int

const (
	MessageTypeDefault MessageType = iota
	MessageTypeRecipientAdd
	MessageTypeRecipientRemove
	MessageTypeCall
	MessageTypeChannelNameChange
	MessageTypeChannelIconChange
	MessageTypeChannelPinnedMessage
	MessageTypeUserJoin
	MessageTypeGuildBoost
	MessageTypeGuildBoostTier1
	MessageTypeGuildBoostTier2
	MessageTypeGuildBoostTier3
	MessageTypeChannelFollowAdd
	_
	MessageTypeGuildDiscoveryDisqualified
	MessageTypeGuildDiscoveryRequalified
	MessageTypeGuildDiscoveryGracePeriodInitialWarning
	MessageTypeGuildDiscoveryGracePeriodFinalWarning
	MessageTypeThreadCreated
	MessageTypeReply
	MessageTypeChatInputCommand
	MessageTypeThreadStarterMessage
	MessageTypeGuildInviteReminder
	MessageTypeContextMenuCommand
	MessageTypeAutoModerationAction
	MessageTypeRoleSubscriptionPurchase
	MessageTypeInteractionPremiumUpsell
	MessageTypeStageStart
	MessageTypeStageEnd
	MessageTypeStageSpeaker
	_
	MessageTypeStageTopic
	MessageTypeGuildApplicationPremiumSubscription
)

// MessageActivity represents the message activity, sent with Rich Presence-related chat embeds (discord.Embed)
// https://discord.com/developers/docs/resources/channel#message-object-message-activity-structure
type MessageActivity struct {
	Type    MessageActivityType `json:"type"`
	PartyID string              `json:"party_id,omitempty"`
}

// MessageActivityType represents teh type of message activity (discord.MessageActivity)
type MessageActivityType int

const (
	MessageActivityTypeJoin MessageActivityType = 1 + iota
	MessageActivityTypeSpectate
	MessageActivityTypeListen
	MessageActivityTypeJoinRequest
)

// MessageReference represents the referenced message (discord.Message) during a cross-post, channel follow add, pin, or reply message
// https://discord.com/developers/docs/resources/channel#message-reference-object-message-reference-structure
type MessageReference struct {
	MessageID       Snowflake `json:"message_id"`
	ChannelID       Snowflake `json:"channel_id,omitempty"`
	GuildID         Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists bool      `json:"fail_if_not_exists,omitempty"`
}

// MessageFlags represents the flags for a message (discord.Message)
// https://discord.com/developers/docs/resources/channel#message-object-message-flags
type MessageFlags uint64

const MessageFlagsNone MessageFlags = 0
const (
	MessageFlagCrossposted MessageFlags = 1 << iota
	MessageFlagIsCrosspost
	MessageFlagSuppressEmbeds
	MessageFlagSourceMessageDeleted
	MessageFlagUrgent
	MessageFlagHasThread
	MessageFlagEphemeral
	MessageFlagLoading
	MessageFlagFailedToMentionSomeRolesInThread
	_
	_
	_
	MessageFlagSuppressNotifications
	MessageFlagIsVoiceMessage
)

// ComputeMessageFlags creates a new message flags structure (discord.MessageFlags) from the given message flags
func ComputeMessageFlags(messageFlags ...MessageFlags) MessageFlags {
	return flags.Compute(messageFlags...)
}

// Add adds the given message flags (discord.MessageFlags) to the current message flags
func (f MessageFlags) Add(messageFlags ...MessageFlags) MessageFlags {
	return flags.Add(f, messageFlags...)
}

// Remove removes the given message flags (discord.MessageFlags)
func (f MessageFlags) Remove(messageFlags ...MessageFlags) MessageFlags {
	return flags.Remove(f, messageFlags...)
}

// Toggle toggles the given message flags (discord.MessageFlags)
func (f MessageFlags) Toggle(messageFlags ...MessageFlags) MessageFlags {
	return flags.Toggle(f, messageFlags...)
}

// Has checks if all the given message flags (discord.MessageFlags) are set
func (f MessageFlags) Has(messageFlags ...MessageFlags) bool {
	return flags.Has(f, messageFlags...)
}

// HasAny checks if any of the given message flags (discord.MessageFlags) is set
func (f MessageFlags) HasAny(messageFlags ...MessageFlags) bool {
	return flags.HasAny(f, messageFlags...)
}

// HasNot checks if all the given message flags (discord.MessageFlags) are not set
func (f MessageFlags) HasNot(messageFlags ...MessageFlags) bool {
	return flags.HasNot(f, messageFlags...)
}

// HasNotAny checks if any of the given message flags (discord.MessageFlags) is not set
func (f MessageFlags) HasNotAny(messageFlags ...MessageFlags) bool {
	return flags.HasNotAny(f, messageFlags...)
}

// MessageInteraction represents an interaction structure that is sent if the message is a response to an interaction
// https://discord.com/developers/docs/interactions/receiving-and-responding#message-interaction-object-message-interaction-structure
type MessageInteraction struct {
	ID     Snowflake       `json:"id"`
	Type   InteractionType `json:"type"`
	Name   string          `json:"name"`
	User   *User           `json:"user"`
	Member *Member         `json:"member,omitempty"`
}

// CreatedAt returns the creation time of the message interaction (discord.MessageInteraction)
func (i *MessageInteraction) CreatedAt() time.Time {
	return i.ID.CreatedAt()
}

// File represents a file that can be used to send to Discord in a message
type File struct {
	Name        string
	Description string
	Reader      io.Reader
}

// CreateMessage represents the payload to send to Discord to post a message (discord.Message) in a channel (discord.Channel)
// https://discord.com/developers/docs/resources/channel#create-message-jsonform-params
type CreateMessage struct {
	Content          *string           `json:"content,omitempty"`
	Nonce            *string           `json:"nonce,omitempty"`
	TTS              *bool             `json:"tts,omitempty"`
	Embeds           []Embed           `json:"embeds,omitempty"`
	AllowedMentions  *AllowedMentions  `json:"allowed_mentions,omitempty"`
	MessageReference *MessageReference `json:"message_reference,omitempty"`
	Components       []Component       `json:"components,omitempty"`
	StickerIDs       ArraySnowflakes   `json:"sticker_ids,omitempty"`
	Flags            *MessageFlags     `json:"flags,omitempty"`
	Attachments      []AttachmentSend  `json:"attachments,omitempty"`

	Files []File `json:"-"`
}

// EditMessage represents the payload to send to Discord to edit an existing message (discord.Message) in a channel (discord.Channel)
// https://discord.com/developers/docs/resources/channel#edit-message-jsonform-params
type EditMessage struct {
	Content         *string          `json:"content,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	Flags           *MessageFlags    `json:"flags,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Attachments     []AttachmentSend `json:"attachments,omitempty"`

	Files []File `json:"-"`
}

// BulkDeleteMessages represents the payload to send to Discord to perform a bulk delete of multiple messages (discord.Message)
// https://discord.com/developers/docs/resources/channel#bulk-delete-messages-json-params
type BulkDeleteMessages struct {
	Messages ArraySnowflakes `json:"messages"`

	AuditLogReason string `json:"-"`
}
