package discord

import (
	"time"
)

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

// ChannelType represents the type of channel (discord.Channel)
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
