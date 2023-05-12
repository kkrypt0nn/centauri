package discord

import (
	"encoding/json"
	"fmt"
	"time"
)

type Channel struct {
	ID                            string                 `json:"id"`
	Type                          ChannelType            `json:"type"`
	GuildID                       string                 `json:"guild_id,omitempty"`
	Position                      int                    `json:"position,omitempty"`
	PermissionOverwrites          []Overwrite            `json:"permission_overwrites"`
	Name                          string                 `json:"name,omitempty"`
	Topic                         string                 `json:"topic,omitempty"`
	NSFW                          bool                   `json:"nsfw,omitempty"`
	LastMessageID                 string                 `json:"last_message_id,omitempty"`
	Bitrate                       int                    `json:"bitrate,omitempty"`
	UserLimit                     int                    `json:"user_limit,omitempty"`
	RateLimitPerUser              int                    `json:"rate_limit_per_user,omitempty"`
	Recipients                    []User                 `json:"recipients,omitempty"`
	Icon                          string                 `json:"icon,omitempty"`
	OwnerID                       string                 `json:"owner_id,omitempty"`
	ApplicationID                 string                 `json:"application_id,omitempty"`
	Managed                       bool                   `json:"managed,omitempty"`
	ParentID                      string                 `json:"parent_id,omitempty"`
	LastPinTimestamp              time.Time              `json:"last_pin_timestamp,omitempty"`
	RTCRegion                     string                 `json:"rtc_region,omitempty"`
	VideoQualityMode              VideoQualityMode       `json:"video_quality_mode,omitempty"`
	MessageCount                  int                    `json:"message_count,omitempty"`
	MemberCount                   int                    `json:"member_count,omitempty"`
	ThreadMetadata                ThreadMetadata         `json:"thread_metadata,omitempty"`
	Member                        ThreadMember           `json:"member,omitempty"`
	DefaultAutoArchiveDuration    int                    `json:"default_auto_archive_duration,omitempty"`
	Permissions                   uint64                 `json:"permissions,string,omitempty"`
	Flags                         ChannelFlags           `json:"flags,omitempty"`
	TotalMessageSent              int                    `json:"total_message_sent,omitempty"`
	AvailableTags                 []Tag                  `json:"available_tags,omitempty"`
	AppliedTags                   []string               `json:"applied_tags,omitempty"`
	DefaultReactionEmoji          DefaultReaction        `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser int                    `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              DefaultSortOrderType   `json:"default_sort_order,omitempty"`
	DefaultForumLayout            DefaultForumLayoutView `json:"default_forum_layout,omitempty"`
}

type ChannelType int

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildAnnouncement
	ChannelTypeGuildStore // Undocumented
	_
	_
	_
	ChannelTypeAnnouncementThread
	ChannelTypePublicThread
	ChannelTypePrivateThread
	ChannelTypeGuildStageVoice
	ChannelTypeGuildDirectory
	ChannelTypeGuildForum
)

type Overwrite struct {
	ID    string        `json:"id"`
	Type  OverwriteType `json:"type"`
	Allow uint64        `json:"allow,string"`
	Deny  uint64        `json:"deny,string"`
}

type OverwriteType int

const (
	OverwriteTypeRole OverwriteType = iota
	OverwriteTypeMember
)

type VideoQualityMode int

const (
	VideoQualityModeAuto VideoQualityMode = 1 + iota
	VideoQualityModeFull
)

type ThreadMetadata struct {
	Archived            bool      `json:"archived"`
	AutoArchiveDuration int       `json:"auto_archive_duration"`
	ArchiveTimestamp    time.Time `json:"archive_timestamp"`
	Locked              bool      `json:"locked"`
	Invitable           bool      `json:"invitable,omitempty"`
	CreateTimestamp     time.Time `json:"create_timestamp,omitempty"`
}

type ThreadMember struct {
	ID            string    `json:"id,omitempty"`
	UserID        string    `json:"user_id,omitempty"`
	JoinTimestamp time.Time `json:"join_timestamp"`
	Flags         uint64    `json:"flags"`
	Member        *Member   `json:"member,omitempty"`
}

type ChannelFlags uint64

const ChannelFlagsNone ChannelFlags = 0
const (
	_ ChannelFlags = 1 << iota
	ChannelFlagPinned
	_
	_
	ChannelFlagRequireTag
)

type Tag struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Moderated bool   `json:"moderated"`
	EmojiID   string `json:"emoji_id,omitempty"`
	EmojiName string `json:"emoji_name,omitempty"`
}

type DefaultReaction struct {
	EmojiID   string `json:"emoji_id,omitempty"`
	EmojiName string `json:"emoji_name,omitempty"`
}

type DefaultSortOrderType int

const (
	DefaultSortOrderTypeLatestActivity DefaultSortOrderType = iota
	DefaultSortOrderTypeCreationDate
)

type DefaultForumLayoutView int

const (
	DefaultForumLayoutViewNotSet DefaultForumLayoutView = iota
	DefaultForumLayoutViewListView
	DefaultForumLayoutViewGalleryView
)

type Message struct {
	ID                   string                `json:"id"`
	ChannelID            string                `json:"channel_id"`
	Author               *User                 `json:"author"`
	Content              string                `json:"content"`
	Timestamp            time.Time             `json:"timestamp"`
	EditedTimestamp      time.Time             `json:"edited_timestamp,omitempty"`
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

type EmbedType string

const (
	EmbedTypeRich    EmbedType = "rich"
	EmbedTypeImage   EmbedType = "image"
	EmbedTypVideo    EmbedType = "video"
	EmbedTypeGIFV    EmbedType = "gifv"
	EmbedTypeArticle EmbedType = "article"
	EmbedTypeLink    EmbedType = "link"
)

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedResource struct {
	Text         string `json:"text"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
	Height       int    `json:"height,omitempty"`
	Width        int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type Reaction struct {
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emoji `json:"emoji"`
}

type Emoji struct {
	ID            string   `json:"id"`
	Name          string   `json:"name,omitempty"`
	Roles         []string `json:"roles,omitempty"`
	User          *User    `json:"user,omitempty"`
	RequireColons bool     `json:"require_colons,omitempty"`
	Managed       bool     `json:"managed,omitempty"`
	Animated      bool     `json:"animated,omitempty"`
	Available     bool     `json:"available,omitempty"`
}

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

type MessageActivity struct {
	Type    MessageActivityType `json:"type"`
	PartyID string              `json:"party_id,omitempty"`
}

type MessageActivityType int

const (
	MessageActivityTypeJoin MessageActivityType = 1 + iota
	MessageActivityTypeSpectate
	MessageActivityTypeListen
	MessageActivityTypeJoinRequest
)

type MessageReference struct {
	MessageID       string `json:"message_id"`
	ChannelID       string `json:"channel_id,omitempty"`
	GuildID         string `json:"guild_id,omitempty"`
	FailIfNotExists bool   `json:"fail_if_not_exists,omitempty"`
}

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

type MessageInteraction struct {
	ID     string          `json:"id"`
	Type   InteractionType `json:"type"`
	Name   string          `json:"name"`
	User   *User           `json:"user"`
	Member *Member         `json:"member,omitempty"`
}

type RoleSubscriptionData struct {
	RoleSubscriptionListingID string `json:"role_subscription_listing_id"`
	TierName                  string `json:"tier_name"`
	TotalMonthsSubscribed     int    `json:"total_months_subscribed"`
	IsRenewal                 bool   `json:"is_renewal"`
}

type ArchivedThreads struct {
	Threads []Channel `json:"threads"`
	Members []Member  `json:"members"`
	HasMore bool      `json:"has_more"`
}
