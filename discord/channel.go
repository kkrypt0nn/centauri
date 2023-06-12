package discord

import "time"

// Channel represents a guild or DM channel within Discord
// https://discord.com/developers/docs/resources/channel#channel-object-channel-structure
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
	LastPinTimestamp              *time.Time             `json:"last_pin_timestamp,omitempty"`
	RTCRegion                     string                 `json:"rtc_region,omitempty"`
	VideoQualityMode              VideoQualityMode       `json:"video_quality_mode,omitempty"`
	MessageCount                  int                    `json:"message_count,omitempty"`
	MemberCount                   int                    `json:"member_count,omitempty"`
	ThreadMetadata                ThreadMetadata         `json:"thread_metadata,omitempty"`
	Member                        ThreadMember           `json:"member,omitempty"`
	DefaultAutoArchiveDuration    int                    `json:"default_auto_archive_duration,omitempty"`
	Permissions                   Permissions            `json:"permissions,string,omitempty"`
	Flags                         ChannelFlags           `json:"flags,omitempty"`
	TotalMessageSent              int                    `json:"total_message_sent,omitempty"`
	AvailableTags                 []Tag                  `json:"available_tags,omitempty"`
	AppliedTags                   []string               `json:"applied_tags,omitempty"`
	DefaultReactionEmoji          DefaultReaction        `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser int                    `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              DefaultSortOrderType   `json:"default_sort_order,omitempty"`
	DefaultForumLayout            DefaultForumLayoutView `json:"default_forum_layout,omitempty"`
}

// ThreadType represents the type of thread
// https://discord.com/developers/docs/resources/channel#channel-object-channel-types
type ThreadType int

const (
	ThreadTypePublic ThreadType = 11 + iota
	ThreadTypePrivate
)

// FollowedChannel represents a channel that is being followed in a target channel
// https://discord.com/developers/docs/resources/channel#followed-channel-object-followed-channel-structure
type FollowedChannel struct {
	ChannelID string `json:"channel_id"`
	WebhookID string `json:"webhook_id"`
}

// ChannelType represents the type of channel (discord.Channel)
// https://discord.com/developers/docs/resources/channel#channel-object-channel-types
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

// Overwrite represents explicit permission overwrites for members (discord.Member) and roles (discord.Role)
// https://discord.com/developers/docs/resources/channel#overwrite-object
type Overwrite struct {
	ID    string        `json:"id"`
	Type  OverwriteType `json:"type"`
	Allow uint64        `json:"allow,string"`
	Deny  uint64        `json:"deny,string"`
}

// OverwriteType represents the type of overwrite (discord.Overwrite)
type OverwriteType int

const (
	OverwriteTypeRole OverwriteType = iota
	OverwriteTypeMember
)

// VideoQualityMode represents the camera video quality mode of the voice channel (discord.Channel)
// https://discord.com/developers/docs/resources/channel#channel-object-video-quality-modes
type VideoQualityMode int

const (
	VideoQualityModeAuto VideoQualityMode = 1 + iota
	VideoQualityModeFull
)

// ThreadMetadata represent additional information for thread channels (discord.Channel)
// https://discord.com/developers/docs/resources/channel#thread-metadata-object-thread-metadata-structure
type ThreadMetadata struct {
	Archived            bool       `json:"archived"`
	AutoArchiveDuration int        `json:"auto_archive_duration"`
	ArchiveTimestamp    *time.Time `json:"archive_timestamp"`
	Locked              bool       `json:"locked"`
	Invitable           bool       `json:"invitable,omitempty"`
	CreateTimestamp     time.Time  `json:"create_timestamp,omitempty"`
}

// ThreadMember represents a member that has joined a thread (discord.Channel & discord.ThreadMetadata)
// https://discord.com/developers/docs/resources/channel#thread-member-object-thread-member-structure
type ThreadMember struct {
	ID            string     `json:"id,omitempty"`
	UserID        string     `json:"user_id,omitempty"`
	JoinTimestamp *time.Time `json:"join_timestamp"`
	Flags         uint64     `json:"flags"`
	Member        *Member    `json:"member,omitempty"`
}

// ChannelFlags represents the channel (discord.Channel) flags
// https://discord.com/developers/docs/resources/channel#channel-object-channel-flags
type ChannelFlags uint64

const ChannelFlagsNone ChannelFlags = 0
const (
	_ ChannelFlags = 1 << iota
	ChannelFlagPinned
	_
	_
	ChannelFlagRequireTag
)

// Tag represents a tag that is able to be applied to a thread in a forum channel (discord.ChannelTypeGuildForum)
// https://discord.com/developers/docs/resources/channel#forum-tag-object-forum-tag-structure
type Tag struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Moderated bool   `json:"moderated"`
	EmojiID   string `json:"emoji_id,omitempty"`
	EmojiName string `json:"emoji_name,omitempty"`
}

// DefaultReaction represents the default reaction that is used when creating a new post in a forum channel (discord.ChannelTypeGuildForum)
// https://discord.com/developers/docs/resources/channel#default-reaction-object-default-reaction-structure
type DefaultReaction struct {
	EmojiID   string `json:"emoji_id,omitempty"`
	EmojiName string `json:"emoji_name,omitempty"`
}

// DefaultSortOrderType represents the default sorting method for forum channels (discord.ChannelTypeGuildForum)
// https://discord.com/developers/docs/resources/channel#channel-object-sort-order-types
type DefaultSortOrderType int

const (
	DefaultSortOrderTypeLatestActivity DefaultSortOrderType = iota
	DefaultSortOrderTypeCreationDate
)

// DefaultForumLayoutView represents the default layout view for forum channels (discord.ChannelTypeGuildForum)
// https://discord.com/developers/docs/resources/channel#channel-object-forum-layout-types
type DefaultForumLayoutView int

const (
	DefaultForumLayoutViewNotSet DefaultForumLayoutView = iota
	DefaultForumLayoutViewListView
	DefaultForumLayoutViewGalleryView
)

// ArchivedThreads represents the list of archived threads and members within those
// https://discord.com/developers/docs/resources/channel#list-public-archived-threads-response-body
type ArchivedThreads struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
	HasMore bool           `json:"has_more"`
}

// ActiveThreads represents the list of active threads and members within those
// https://discord.com/developers/docs/resources/guild#list-active-guild-threads-response-body
type ActiveThreads struct {
	Threads []Channel      `json:"threads"`
	Members []ThreadMember `json:"members"`
}

// CreateGuildChannel represents the payload to send to Discord to create a new channel (discord.Channel) in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#create-guild-channel-json-params
type CreateGuildChannel struct {
	Name                       string                `json:"name"`
	Type                       *ChannelType          `json:"type,omitempty"`
	Topic                      *string               `json:"topic,omitempty"`
	Bitrate                    *int                  `json:"bitrate,omitempty"`
	UserLimit                  *int                  `json:"user_limit,omitempty"`
	RateLimitPerUser           *int                  `json:"rate_limit_per_user,omitempty"`
	Position                   *int                  `json:"position,omitempty"`
	PermissionOverwrites       []Overwrite           `json:"permission_overwrites"`
	ParentID                   *string               `json:"parent_id,omitempty"`
	NSFW                       *bool                 `json:"nsfw,omitempty"`
	RTCRegion                  *string               `json:"rtc_region,omitempty"`
	VideoQualityMode           *VideoQualityMode     `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration *int                  `json:"default_auto_archive_duration,omitempty"`
	DefaultReaction            *DefaultReaction      `json:"default_reaction,omitempty"`
	AvailableTags              []Tag                 `json:"available_tags,omitempty"`
	DefaultSortOrder           *DefaultSortOrderType `json:"default_sort_order,omitempty"`

	AuditLogReason string `json:"-"`
}

// ModifyGuildChannelPosition represents the payload to send to Discord to modify the position of an existing channel (discord.Channel) in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#modify-guild-channel-positions-json-params
type ModifyGuildChannelPosition struct {
	ID              string  `json:"id"`
	Position        *int    `json:"position,omitempty"`
	LockPermissions *bool   `json:"lock_permissions,omitempty"`
	ParentID        *string `json:"parent_id,omitempty"`
}

// ModifyChannel represents the payload to send to Discord to modify an existing channel (discord.Channel)
// https://discord.com/developers/docs/resources/channel#modify-channel-json-params-group-dm
// https://discord.com/developers/docs/resources/channel#modify-channel-json-params-guild-channel
// https://discord.com/developers/docs/resources/channel#modify-channel-json-params-thread
type ModifyChannel struct {
	Name                          *string                 `json:"name,omitempty"`
	Type                          *ChannelType            `json:"type,omitempty"`
	Position                      *int                    `json:"position,omitempty"`
	Topic                         *string                 `json:"topic,omitempty"`
	NSFW                          *bool                   `json:"nsfw,omitempty"`
	RateLimitPerUser              *int                    `json:"rate_limit_per_user,omitempty"`
	Bitrate                       *int                    `json:"bitrate,omitempty"`
	UserLimit                     *int                    `json:"user_limit,omitempty"`
	PermissionOverwrites          []Overwrite             `json:"permission_overwrites,omitempty"`
	ParentID                      *string                 `json:"parent_id,omitempty"`
	RTCRegion                     *string                 `json:"rtc_region,omitempty"`
	VideoQualityMode              *VideoQualityMode       `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration    *int                    `json:"default_auto_archive_duration,omitempty"`
	Flags                         *ChannelFlags           `json:"flags,omitempty"`
	AvailableTags                 []Tag                   `json:"available_tags,omitempty"`
	DefaultReaction               *DefaultReaction        `json:"default_reaction,omitempty"`
	DefaultThreadRateLimitPerUser *int                    `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              *DefaultSortOrderType   `json:"default_sort_order,omitempty"`
	DefaultForumLayoutView        *DefaultForumLayoutView `json:"default_forum_layout_view,omitempty"`

	// Group DMS only
	Icon *string `json:"icon,omitempty"`

	// Thread only
	Archived            *bool    `json:"archived,omitempty"`
	AutoArchiveDuration *int     `json:"auto_archive_duration,omitempty"`
	Locked              *bool    `json:"locked,omitempty"`
	Invitable           *bool    `json:"invitable,omitempty"`
	AppliedTags         []string `json:"applied_tags,omitempty"`

	AuditLogReason string `json:"-"`
}

// EditChannelPermissions represents the payload to send to Discord to edit the channel permission overwrites
// https://discord.com/developers/docs/resources/channel#edit-channel-permissions-json-params
type EditChannelPermissions struct {
	Type  OverwriteType `json:"type"`
	Allow *string       `json:"allow,omitempty"`
	Deny  *string       `json:"deny,omitempty"`

	AuditLogReason string `json:"-"`
}

// StartThreadFromMessage represents the payload to send to Discord to start a thread from an existing message (discord.Message)
// https://discord.com/developers/docs/resources/channel#start-thread-from-message-json-params
type StartThreadFromMessage struct {
	Name                string `json:"name"`
	AutoArchiveDuration *int   `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser    *int   `json:"rate_limit_per_user,omitempty"`

	AuditLogReason string `json:"-"`
}

// StartThreadWithoutMessage represents the payload to send to Discord to start a thread
// https://discord.com/developers/docs/resources/channel#start-thread-without-message-json-params
type StartThreadWithoutMessage struct {
	Name                string      `json:"name"`
	AutoArchiveDuration *int        `json:"auto_archive_duration,omitempty"`
	Type                *ThreadType `json:"thread_type,omitempty"`
	Invitable           *bool       `json:"invitable,omitempty"`
	RateLimitPerUser    *int        `json:"rate_limit_per_user,omitempty"`

	AuditLogReason string `json:"-"`
}

// StartThreadInForumChannel represents the payload to send to Discord to start a thread in a forum channel (discord.Channel)
// https://discord.com/developers/docs/resources/channel#start-thread-in-forum-channel-jsonform-params
type StartThreadInForumChannel struct {
	Name                string             `json:"name"`
	AutoArchiveDuration *int               `json:"auto_archive_duration,omitempty"`
	RateLimitPerUser    *int               `json:"rate_limit_per_user,omitempty"`
	Message             ForumThreadMessage `json:"message"`
	AppliedTags         []Tag              `json:"applied_tags,omitempty"`

	AuditLogReason string `json:"-"`
}

// ForumThreadMessage represents a forum thread message when creating a new thread
// https://discord.com/developers/docs/resources/channel#start-thread-in-forum-channel-forum-thread-message-params-object
type ForumThreadMessage struct {
	Content         *string          `json:"content,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Components      []Component      `json:"components,omitempty"`
	StickerIDs      []string         `json:"sticker_ids,omitempty"`
	Flags           *MessageFlags    `json:"flags,omitempty"`
	Attachments     []Attachment     `json:"attachments,omitempty"`

	Files []File `json:"-"`
}
