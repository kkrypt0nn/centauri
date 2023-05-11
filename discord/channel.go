package discord

import "time"

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
