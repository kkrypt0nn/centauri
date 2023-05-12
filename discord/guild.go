package discord

import (
	"github.com/kkrypt0nn/centauri/oauth2"
	"time"
)

type Guild struct {
	ID                          string                   `json:"id"`
	Name                        string                   `json:"name"`
	Icon                        string                   `json:"icon"`
	IconHash                    string                   `json:"icon_hash"`
	Splash                      string                   `json:"splash"`
	DiscoverySplash             string                   `json:"discovery_splash"`
	Owner                       bool                     `json:"owner"`
	OwnerID                     string                   `json:"owner_id"`
	Permissions                 uint64                   `json:"permissions,string"`
	Region                      string                   `json:"region"`
	AFKChannelID                string                   `json:"afk_channel_id"`
	AFKTimeout                  int                      `json:"afk_timeout"`
	WidgetEnabled               bool                     `json:"widget_enabled"`
	WidgetChannelID             string                   `json:"widget_channel_id"`
	VerificationLevel           GuildVerificationLevel   `json:"verification_level"`
	DefaultMessageNotifications MessageNotificationLevel `json:"default_message_notifications"`
	ExplicitContentFilter       ExplicitContentFilter    `json:"explicit_content_filter"`
	Roles                       []Role                   `json:"roles"`
	Emojis                      []Emoji                  `json:"emojis"`
	Features                    []GuildFeature           `json:"features"`
	MFALevel                    MFALevel                 `json:"mfa_level"`
	ApplicationID               string                   `json:"application_id"`
	SystemChannelID             string                   `json:"system_channel_id"`
	SystemChannelFlags          SystemChannelFlags       `json:"system_channel_flags"`
	RulesChannelID              string                   `json:"rules_channel_id"`
	MaxPresences                int                      `json:"max_presences"`
	MaxMembers                  int                      `json:"max_members"`
	VanityURLCode               string                   `json:"vanity_url_code"`
	Description                 string                   `json:"description"`
	Banner                      string                   `json:"banner"`
	PremiumTier                 PremiumTier              `json:"premium_tier"`
	PremiumSubscriptionCount    int                      `json:"premium_subscription_count"`
	PreferredLocale             Locale                   `json:"preferred_locale"`
	PublicUpdatesChannelID      string                   `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        int                      `json:"max_video_channel_users"`
	MaxStageVideoChannelUsers   int                      `json:"max_stage_video_channel_users"`
	ApproximateMemberCount      int                      `json:"approximate_member_count"`
	ApproximatePresenceCount    int                      `json:"approximate_presence_count"`
	WelcomeScreen               *WelcomeScreen           `json:"welcome_screen"`
	NSFWLevel                   NSFWLevel                `json:"nsfw_level"`
	Stickers                    []Sticker                `json:"stickers"`
	PremiumProgressBarEnabled   bool                     `json:"premium_progress_bar_enabled"`
}

type GuildVerificationLevel int

const (
	GuildVerificationLevelNone GuildVerificationLevel = iota
	GuildVerificationLevelLow
	GuildVerificationLevelMedium
	GuildVerificationLevelHigh
	GuildVerificationLevelVeryHigh
)

type MessageNotificationLevel int

const (
	MessageNotificationLevelAllMessages MessageNotificationLevel = iota
	MessageNotificationLevelOnlyMentions
)

type ExplicitContentFilter int

const (
	ExplicitContentFilterDisabled ExplicitContentFilter = iota
	ExplicitContentFilterMembersWithoutRoles
	ExplicitContentFilterAllMembers
)

type MFALevel int

const (
	MFALevelNone MFALevel = iota
	MFALevelElevated
)

type SystemChannelFlags uint64

const SystemChannelFlagsNone SystemChannelFlags = 0
const (
	SystemChannelFlagsSuppressJoinNotifications SystemChannelFlags = 1 << iota
	SystemChannelFlagsSuppressPremiumSubscriptions
	SystemChannelFlagsSuppressGuildReminderNotification
	SystemChannelFlagsSuppressJoinNotificationReplies
	SystemChannelFlagsSuppressRoleSubscriptionPurchaseNotifications
	SystemChannelFlagsSuppressRoleSubscriptionPurchaseNotificationsReplies
)

type PremiumTier int

const (
	PremiumTierNone PremiumTier = iota
	PremiumTier1
	PremiumTier2
	PremiumTier3
)

type WelcomeScreen struct {
	Description     string                 `json:"description"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	ChannelID   string `json:"channel_id"`
	Description string `json:"description"`
	EmojiID     string `json:"emoji_id"`
	EmojiName   string `json:"emoji_name"`
}

type NSFWLevel int

const (
	NSFWLevelDefault NSFWLevel = iota
	NSFWLevelExplicit
	NSFWLevelSafe
	NSFWLevelAgeRestricted
)

type GuildPreview struct {
	ID                       string         `json:"id"`
	Name                     string         `json:"name"`
	Icon                     string         `json:"icon"`
	Splash                   string         `json:"splash"`
	DiscoverySplash          string         `json:"discovery_splash"`
	Emojis                   []Emoji        `json:"emojis"`
	Features                 []GuildFeature `json:"features"`
	ApproximateMemberCount   int            `json:"approximate_member_count"`
	ApproximatePresenceCount int            `json:"approximate_presence_count"`
	Description              string         `json:"description"`
	Stickers                 []Sticker      `json:"stickers"`
}

type GuildBan struct {
	Reason string `json:"reason"`
	User   *User  `json:"user"`
}

type GuildPrune struct {
	Pruned int `json:"pruned"`
}

type WidgetSetting struct {
	Enabled  bool   `json:"enabled"`
	ChanelID string `json:"chanel_id"`
}

type Widget struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	InstantInvite string    `json:"instant_invite"`
	Channels      []Channel `json:"channels"`
	Members       []Member  `json:"members"`
	PresenceCount int       `json:"presence_count"`
}

type Onboarding struct {
	GuildID           string             `json:"guild_id"`
	Prompts           []OnboardingPrompt `json:"prompts"`
	DefaultChannelIDs []string           `json:"default_channel_ids"`
	Enabled           bool               `json:"enabled"`
}

type OnboardingPrompt struct {
	ID           string                   `json:"id"`
	Type         OnboardingPromptType     `json:"type"`
	Options      []OnboardingPromptOption `json:"options"`
	Title        string                   `json:"title"`
	SingleSelect bool                     `json:"single_select"`
	Required     bool                     `json:"required"`
	InOnboarding bool                     `json:"in_onboarding"`
}

type OnboardingPromptType int

const (
	OnboardingPromptTypeMultipleChoice = iota
	OnboardingPromptTypeDropdown
)

type OnboardingPromptOption struct {
	ID          string   `json:"id"`
	ChannelIDs  []string `json:"channel_ids"`
	RoleIDs     []string `json:"role_ids"`
	Emoji       *Emoji   `json:"emoji"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
}

// PartialGuild represents a discord.Guild with only a few fields
// https://discord.com/developers/docs/resources/user#get-current-user-guilds-example-partial-guild
type PartialGuild struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Icon        string         `json:"icon"`
	Owner       bool           `json:"owner"`
	Permissions uint64         `json:"permissions,string"`
	Features    []GuildFeature `json:"features"`
}

// GuildFeature represents the different features a guild may have activated
// https://discord.com/developers/docs/resources/guild#guild-object-guild-features
type GuildFeature string

const (
	GuildFeatureAnimatedBanner                        GuildFeature = "ANIMATED_BANNER"
	GuildFeatureAnimatedIcon                          GuildFeature = "ANIMATED_ICON"
	GuildFeatureApplicationCommandPermissionsV2       GuildFeature = "APPLICATION_COMMAND_PERMISSIONS_V2"
	GuildFeatureAutoModeration                        GuildFeature = "AUTO_MODERATION"
	GuildFeatureBanner                                GuildFeature = "BANNER"
	GuildFeatureCommunity                             GuildFeature = "COMMUNITY"
	GuildFeatureCreatorMonetizableProvisional         GuildFeature = "CREATOR_MONETIZABLE_PROVISIONAL"
	GuildFeatureCreatorStorePage                      GuildFeature = "CREATOR_STORE_PAGE"
	GuildFeatureDeveloperSupportServer                GuildFeature = "DEVELOPER_SUPPORT_SERVER"
	GuildFeatureDiscoverable                          GuildFeature = "DISCOVERABLE"
	GuildFeatureFeaturable                            GuildFeature = "FEATURABLE"
	GuildFeatureInvitesDisabled                       GuildFeature = "INVITES_DISABLED"
	GuildFeatureInviteSplash                          GuildFeature = "INVITE_SPLASH"
	GuildFeatureMemberVerificationGateEnabled         GuildFeature = "MEMBER_VERIFICATION_GATE_ENABLED"
	GuildFeatureMoreStickers                          GuildFeature = "MORE_STICKERS"
	GuildFeatureNews                                  GuildFeature = "NEWS"
	GuildFeaturePartnered                             GuildFeature = "PARTNERED"
	GuildFeaturePreviewEnabled                        GuildFeature = "PREVIEW_ENABLED"
	GuildFeatureRoleIcons                             GuildFeature = "ROLE_ICONS"
	GuildFeatureRoleSubscriptionsAvailableForPurchase GuildFeature = "ROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE"
	GuildFeatureRoleSubscriptionsEnabled              GuildFeature = "ROLE_SUBSCRIPTIONS_ENABLED"
	GuildFeatureTicketedEventsEnabled                 GuildFeature = "TICKETED_EVENTS_ENABLED"
	GuildFeatureVanityURL                             GuildFeature = "VANITY_URL"
	GuildFeatureVerified                              GuildFeature = "VERIFIED"
	GuildFeatureVIPRegions                            GuildFeature = "VIP_REGIONS"
	GuildFeatureWelcomeScreenEnabled                  GuildFeature = "WELCOME_SCREEN_ENABLED"
)

type Integration struct {
	ID                string                    `json:"id"`
	Name              string                    `json:"name"`
	Type              string                    `json:"type"`
	Enabled           bool                      `json:"enabled"`
	Syncing           bool                      `json:"syncing,omitempty"`
	RoleID            string                    `json:"role_id,omitempty"`
	EnableEmoticons   bool                      `json:"enable_emoticons,omitempty"`
	ExpireBehavior    IntegrationExpireBehavior `json:"expire_behavior,omitempty"`
	ExpireGracePeriod int                       `json:"expire_grace_period,omitempty"`
	User              *User                     `json:"user,omitempty"`
	Account           IntegrationAccount        `json:"account"`
	SyncedAt          *time.Time                `json:"synced_at,omitempty"`
	SubscriberCount   int                       `json:"subscriber_count,omitempty"`
	Revoked           bool                      `json:"revoked,omitempty"`
	Application       IntegrationApplication    `json:"application,omitempty"`
	Scopes            []oauth2.Scope            `json:"scopes,omitempty"`
}

type IntegrationExpireBehavior int

const (
	IntegrationExpireBehaviorRemoveRole IntegrationExpireBehavior = iota
	IntegrationExpireBehaviorKick
)

type IntegrationAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type IntegrationApplication struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Bot         *User  `json:"bot,omitempty"`
}

type Member struct {
	User                       *User       `json:"user,omitempty"`
	Nick                       string      `json:"nick,omitempty"`
	Avatar                     string      `json:"avatar,omitempty"`
	Roles                      []string    `json:"roles"`
	JoinedAt                   time.Time   `json:"joined_at"`
	PremiumSince               *time.Time  `json:"premium_since,omitempty"`
	Deaf                       bool        `json:"deaf"`
	Mute                       bool        `json:"mute"`
	Flags                      MemberFlags `json:"flags"`
	Pending                    bool        `json:"pending,omitempty"`
	Permissions                uint64      `json:"permissions,string,omitempty"`
	CommunicationDisabledUntil *time.Time  `json:"communication_disabled_until,omitempty"`
}

type MemberFlags uint64

const MemberFlagsNone MemberFlags = 0
const (
	MemberFlagDidRejoin MemberFlags = 1 << iota
	MemberFlagCompletedOnboarding
	MemberFlagBypassesVerification
	MemberFlagStartedOnboarding
)

type Role struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Color        int       `json:"color"`
	Hoist        bool      `json:"hoist"`
	Icon         string    `json:"icon,omitempty"`
	UnicodeEmoji string    `json:"unicode_emoji"`
	Position     int       `json:"position"`
	Permissions  uint64    `json:"permissions,string"`
	Managed      bool      `json:"managed"`
	Mentionable  bool      `json:"mentionable"`
	Tags         *RoleTags `json:"tags,omitempty"`
}

type RoleTags struct {
	BotID                 string `json:"bot_id"`
	IntegrationID         string `json:"integration_id"`
	PremiumSubscriber     bool   `json:"premium_subscriber"`
	SubscriptionListingID string `json:"subscription_listing_id"`
	AvailableForPurchase  bool   `json:"available_for_purchase"`
	GuildConnections      bool   `json:"guild_connections"`
}
