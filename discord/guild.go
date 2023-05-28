package discord

// Guild represents an isolated collection of users and channels, and are often referred to as "servers" in the UI
// https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
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

// CreateGuild represents the payload to send to Discord to create a new guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#create-guild-json-params
type CreateGuild struct {
	Name                        string                    `json:"name"`
	Region                      *string                   `json:"region,omitempty"`
	Icon                        *string                   `json:"icon,omitempty"`
	VerificationLevel           *GuildVerificationLevel   `json:"verification_level,omitempty"`
	DefaultMessageNotifications *MessageNotificationLevel `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       *ExplicitContentFilter    `json:"explicit_content_filter,omitempty"`
	Roles                       []Role                    `json:"roles,omitempty"`
	Channels                    []Channel                 `json:"channels,omitempty"`
	AFKChannelID                *string                   `json:"afk_channel_id,omitempty"`
	AFKTimeout                  *int                      `json:"afk_timeout,omitempty"`
	SystemChannelID             *string                   `json:"system_channel_id,omitempty"`
	SystemChannelFlags          *SystemChannelFlags       `json:"system_channel_flags,omitempty"`
}

// ModifyGuild represents the payload to send to Discord to modify an existing guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#modify-guild-json-params
type ModifyGuild struct {
	Name                        *string                   `json:"name,omitempty"`
	Region                      *string                   `json:"region,omitempty"`
	VerificationLevel           *GuildVerificationLevel   `json:"verification_level,omitempty"`
	DefaultMessageNotifications *MessageNotificationLevel `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       *ExplicitContentFilter    `json:"explicit_content_filter,omitempty"`
	AFKChannelID                *string                   `json:"afk_channel_id,omitempty"`
	AFKTimeout                  *int                      `json:"afk_timeout,omitempty"`
	Icon                        *string                   `json:"icon,omitempty"`
	OwnerID                     *string                   `json:"owner_id,omitempty"`
	Splash                      *string                   `json:"splash,omitempty"`
	DiscoverySplash             *string                   `json:"discovery_splash,omitempty"`
	Banner                      *string                   `json:"banner,omitempty"`
	SystemChannelID             *string                   `json:"system_channel_id,omitempty"`
	SystemChannelFlags          *SystemChannelFlags       `json:"system_channel_flags,omitempty"`
	RulesChannelID              *string                   `json:"rules_channel_id,omitempty"`
	PublicUpdatesChannelID      *string                   `json:"public_updates_channel_id,omitempty"`
	PreferredLocale             *Locale                   `json:"preferred_locale,omitempty"`
	Features                    []GuildFeature            `json:"features,omitempty"`
	Description                 *string                   `json:"description,omitempty"`
	PremiumProgressBarEnabled   *bool                     `json:"premium_progress_bar_enabled,omitempty"`
	SafetyAlertsChannelID       *string                   `json:"safety_alerts_channel_id,omitempty"`

	AuditLogReason string `json:"-"`
}

// ModifyGuildMFALevel represents the payload to send to Discord to modify the MFA level for a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#modify-guild-mfa-level-json-params
type ModifyGuildMFALevel struct {
	Level MFALevel `json:"level"`

	AuditLogReason string `json:"-"`
}

// GuildVerificationLevel represents the verification level used in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#guild-object-verification-level
type GuildVerificationLevel int

const (
	GuildVerificationLevelNone GuildVerificationLevel = iota
	GuildVerificationLevelLow
	GuildVerificationLevelMedium
	GuildVerificationLevelHigh
	GuildVerificationLevelVeryHigh
)

// MessageNotificationLevel represents the default message notification level in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
type MessageNotificationLevel int

const (
	MessageNotificationLevelAllMessages MessageNotificationLevel = iota
	MessageNotificationLevelOnlyMentions
)

// ExplicitContentFilter represents the explicit content filter in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#guild-object-explicit-content-filter-level
type ExplicitContentFilter int

const (
	ExplicitContentFilterDisabled ExplicitContentFilter = iota
	ExplicitContentFilterMembersWithoutRoles
	ExplicitContentFilterAllMembers
)

// MFALevel represents the required MFA level in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#guild-object-mfa-level
type MFALevel int

const (
	MFALevelNone MFALevel = iota
	MFALevelElevated
)

// SystemChannelFlags represents the channel flags in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
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

// PremiumTier represents the premium tier, also known as server boost, in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#guild-object-premium-tier
type PremiumTier int

const (
	PremiumTierNone PremiumTier = iota
	PremiumTier1
	PremiumTier2
	PremiumTier3
)

// WelcomeScreen represents the welcome screen of a community guild (discord.GuildFeatureCommunity), that is shown to new members
// https://discord.com/developers/docs/resources/guild#welcome-screen-object-welcome-screen-structure
type WelcomeScreen struct {
	Description     string                 `json:"description"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

// ModifyGuildWelcomeScreen represents the payload to send to Discord to modify and existing welcome screen (discord.WelcomeScreen)
// https://discord.com/developers/docs/resources/guild#modify-guild-welcome-screen-json-params
type ModifyGuildWelcomeScreen struct {
	Enabled         *bool                  `json:"enabled,omitempty"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels,omitempty"`
	Description     *string                `json:"description,omitempty"`

	AuditLogReason string `json:"-"`
}

// WelcomeScreenChannel represents a channel shown in the welcome screen
// https://discord.com/developers/docs/resources/guild#welcome-screen-object-welcome-screen-channel-structure
type WelcomeScreenChannel struct {
	ChannelID   string `json:"channel_id"`
	Description string `json:"description"`
	EmojiID     string `json:"emoji_id"`
	EmojiName   string `json:"emoji_name"`
}

// NSFWLevel represents the NSFW level of a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#guild-object-guild-nsfw-level
type NSFWLevel int

const (
	NSFWLevelDefault NSFWLevel = iota
	NSFWLevelExplicit
	NSFWLevelSafe
	NSFWLevelAgeRestricted
)

// GuildPreview represents a small preview of a guild (discord.Guild) with partial information
// https://discord.com/developers/docs/resources/guild#guild-preview-object-guild-preview-structure
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

// GuildBan represents a banned user (discord.User) in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#ban-object-ban-structure
type GuildBan struct {
	Reason string `json:"reason"`
	User   *User  `json:"user"`
}

// CreateGuildBan represents the payload to send to Discord to create a guild ban (discord.GuildBan)
// https://discord.com/developers/docs/resources/guild#create-guild-ban-json-params
type CreateGuildBan struct {
	DeleteMessageSeconds *int `json:"delete_message_seconds,omitempty"`

	UserID         string `json:"-"`
	AuditLogReason string `json:"-"`
}

// GuildPrune represents how many members (discord.Member) will get pruned when performing a prune
// https://discord.com/developers/docs/resources/guild#get-guild-prune-count
type GuildPrune struct {
	Pruned int `json:"pruned"`
}

// BeginGuildPrune represents the payload to send to Discord to begin a guild prune operation (discord.GuildPrune)
// https://discord.com/developers/docs/resources/guild#begin-guild-prune-json-params
type BeginGuildPrune struct {
	Days              *int     `json:"days,omitempty"`
	ComputePruneCount *bool    `json:"compute_prune_count,omitempty"`
	IncludeRoles      []string `json:"include_roles,omitempty"`

	AuditLogReason string `json:"-"`
}

// WidgetSetting represents the channel (discord.Channel) used and whether the widget (discord.Widget) of the guild (discord.Guild) is enabled
// https://discord.com/developers/docs/resources/guild#guild-widget-settings-object-guild-widget-settings-structure
type WidgetSetting struct {
	Enabled   bool   `json:"enabled"`
	ChannelID string `json:"channel_id"`
}

// Widget represents the guild's (discord.Guild) widget
// https://discord.com/developers/docs/resources/guild#guild-widget-object-guild-widget-structure
type Widget struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	InstantInvite string    `json:"instant_invite"`
	Channels      []Channel `json:"channels"`
	Members       []Member  `json:"members"`
	PresenceCount int       `json:"presence_count"`
}

// ModifyGuildWidget represents the payload to send to Discord to modify the guild widget (discord.Widget)
// https://discord.com/developers/docs/resources/guild#guild-widget-settings-object-guild-widget-settings-structure
type ModifyGuildWidget struct {
	Enabled   *bool   `json:"enabled,omitempty"`
	ChannelID *string `json:"channel_id,omitempty"`

	AuditLogReason string `json:"-"`
}

// Onboarding represents the onboarding flow for a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#guild-onboarding-object-guild-onboarding-structure
type Onboarding struct {
	GuildID           string             `json:"guild_id"`
	Prompts           []OnboardingPrompt `json:"prompts"`
	DefaultChannelIDs []string           `json:"default_channel_ids"`
	Enabled           bool               `json:"enabled"`
}

// OnboardingPrompt represents a prompt shown during onboarding (discord.Onboarding) and in customize community
// https://discord.com/developers/docs/resources/guild#guild-onboarding-object-onboarding-prompt-structure
type OnboardingPrompt struct {
	ID           string                   `json:"id"`
	Type         OnboardingPromptType     `json:"type"`
	Options      []OnboardingPromptOption `json:"options"`
	Title        string                   `json:"title"`
	SingleSelect bool                     `json:"single_select"`
	Required     bool                     `json:"required"`
	InOnboarding bool                     `json:"in_onboarding"`
}

// OnboardingPromptType represents the type of prompt for the onboarding prompt (discord.OnboardingPrompt)
// https://discord.com/developers/docs/resources/guild#guild-onboarding-object-prompt-types
type OnboardingPromptType int

const (
	OnboardingPromptTypeMultipleChoice = iota
	OnboardingPromptTypeDropdown
)

// OnboardingPromptOption represents the options available within the prompt (discord.OnboardingPrompt)
// https://discord.com/developers/docs/resources/guild#guild-onboarding-object-prompt-option-structure
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
