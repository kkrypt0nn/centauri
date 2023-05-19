package discord

import (
	"encoding/json"
	"fmt"
	"time"
)

// Message represents a message sent in a channel (discord.Channel) within Discord
// https://discord.com/developers/docs/resources/channel#message-object-message-structure
type Message struct {
	ID                   string                `json:"id"`
	ChannelID            string                `json:"channel_id"`
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
	WebhookID            string                `json:"webhook_id,omitempty"`
	Type                 MessageType           `json:"type"`
	Activity             *MessageActivity      `json:"activity,omitempty"`
	Application          *Application          `json:"application,omitempty"`
	ApplicationID        string                `json:"application_id,omitempty"`
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

// Attachment represents an attached file in a message (discord.Message)
// https://discord.com/developers/docs/resources/channel#attachment-object-attachment-structure
type Attachment struct {
	ID           string  `json:"id"`
	Filename     string  `json:"filename"`
	Description  string  `json:"description,omitempty"`
	ContentType  string  `json:"content_type,omitempty"`
	Size         int     `json:"size,omitempty"`
	URL          string  `json:"url,omitempty"`
	ProxyURL     string  `json:"proxy_url,omitempty"`
	Height       int     `json:"height,omitempty"`
	Width        int     `json:"width,omitempty"`
	Ephemeral    bool    `json:"ephemeral,omitempty"`
	DurationSecs float64 `json:"duration_secs,omitempty"`
	Waveform     string  `json:"waveform,omitempty"`
}

// Embed represents embedded content in a message (discord.Message)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-structure
type Embed struct {
	Title       string         `json:"title,omitempty"`
	Type        EmbedType      `json:"type,omitempty"`
	Description string         `json:"description,omitempty"`
	URL         string         `json:"url,omitempty"`
	Timestamp   time.Time      `json:"timestamp,omitempty"`
	Color       int            `json:"color,omitempty"`
	Footer      *EmbedFooter   `json:"footer,omitempty"`
	Image       *EmbedResource `json:"image,omitempty"`
	Thumbnail   *EmbedResource `json:"thumbnail,omitempty"`
	Video       *EmbedResource `json:"video,omitempty"`
	Provider    *EmbedProvider `json:"provider,omitempty"`
	Author      *EmbedAuthor   `json:"author,omitempty"`
	Fields      []EmbedField   `json:"fields,omitempty"`
}

// EmbedType represents the type of embedded content (discord.Embed)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-types
type EmbedType string

const (
	EmbedTypeRich    EmbedType = "rich"
	EmbedTypeImage   EmbedType = "image"
	EmbedTypVideo    EmbedType = "video"
	EmbedTypeGIFV    EmbedType = "gifv"
	EmbedTypeArticle EmbedType = "article"
	EmbedTypeLink    EmbedType = "link"
)

// EmbedFooter represents footer information in embedded content (discord.Embed)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-footer-structure
type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// EmbedResource represents resource information in embedded content (discord.Embed)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-image-structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-thumbnail-structure
// https://discord.com/developers/docs/resources/channel#embed-object-embed-video-structure
type EmbedResource struct {
	Text         string `json:"text"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
	Height       int    `json:"height,omitempty"`
	Width        int    `json:"width,omitempty"`
}

// EmbedProvider represents the provider information in embedded content (discord.Embed)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-provider-structure
type EmbedProvider struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

// EmbedAuthor represents the author information in embedded content (discord.Embed)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-author-structure
type EmbedAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// EmbedField represents the field information in embedded content (discord.Embed)
// https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
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
	MessageID       string `json:"message_id"`
	ChannelID       string `json:"channel_id,omitempty"`
	GuildID         string `json:"guild_id,omitempty"`
	FailIfNotExists bool   `json:"fail_if_not_exists,omitempty"`
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

// MessageInteraction represents an interaction structure that is sent if the message is a response to an interaction
// https://discord.com/developers/docs/interactions/receiving-and-responding#message-interaction-object-message-interaction-structure
type MessageInteraction struct {
	ID     string          `json:"id"`
	Type   InteractionType `json:"type"`
	Name   string          `json:"name"`
	User   *User           `json:"user"`
	Member *Member         `json:"member,omitempty"`
}
