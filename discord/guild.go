package discord

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/utils/flags"
	"strings"
	"time"
)

// Guild represents an isolated collection of users and channels, and are often referred to as "servers" in the UI
// https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
type Guild struct {
	ID                          Snowflake                `json:"id"`
	Name                        string                   `json:"name"`
	Icon                        string                   `json:"icon"`
	IconHash                    string                   `json:"icon_hash"`
	Splash                      string                   `json:"splash"`
	DiscoverySplash             string                   `json:"discovery_splash"`
	Owner                       bool                     `json:"owner"`
	OwnerID                     Snowflake                `json:"owner_id"`
	Permissions                 uint64                   `json:"permissions"`
	Region                      string                   `json:"region"`
	AFKChannelID                Snowflake                `json:"afk_channel_id"`
	AFKTimeout                  int                      `json:"afk_timeout"`
	WidgetEnabled               bool                     `json:"widget_enabled"`
	WidgetChannelID             Snowflake                `json:"widget_channel_id"`
	VerificationLevel           GuildVerificationLevel   `json:"verification_level"`
	DefaultMessageNotifications MessageNotificationLevel `json:"default_message_notifications"`
	ExplicitContentFilter       ExplicitContentFilter    `json:"explicit_content_filter"`
	Roles                       []Role                   `json:"roles"`
	Emojis                      []Emoji                  `json:"emojis"`
	Features                    []GuildFeature           `json:"features"`
	MFALevel                    MFALevel                 `json:"mfa_level"`
	ApplicationID               Snowflake                `json:"application_id"`
	SystemChannelID             Snowflake                `json:"system_channel_id"`
	SystemChannelFlags          SystemChannelFlags       `json:"system_channel_flags"`
	RulesChannelID              Snowflake                `json:"rules_channel_id"`
	MaxPresences                int                      `json:"max_presences"`
	MaxMembers                  int                      `json:"max_members"`
	VanityURLCode               string                   `json:"vanity_url_code"`
	Description                 string                   `json:"description"`
	Banner                      string                   `json:"banner"`
	PremiumTier                 PremiumTier              `json:"premium_tier"`
	PremiumSubscriptionCount    int                      `json:"premium_subscription_count"`
	PreferredLocale             Locale                   `json:"preferred_locale"`
	PublicUpdatesChannelID      Snowflake                `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        int                      `json:"max_video_channel_users"`
	MaxStageVideoChannelUsers   int                      `json:"max_stage_video_channel_users"`
	ApproximateMemberCount      int                      `json:"approximate_member_count"`
	ApproximatePresenceCount    int                      `json:"approximate_presence_count"`
	WelcomeScreen               *WelcomeScreen           `json:"welcome_screen"`
	NSFWLevel                   NSFWLevel                `json:"nsfw_level"`
	Stickers                    []Sticker                `json:"stickers"`
	PremiumProgressBarEnabled   bool                     `json:"premium_progress_bar_enabled"`
	SafetyAlertsChannelID       Snowflake                `json:"safety_alerts_channel_id,omitempty"`
}

// CreatedAt returns the creation time of the guild (discord.Guild)
func (g *Guild) CreatedAt() time.Time {
	return g.ID.CreatedAt()
}

// IconURL returns the URL for the guild icon
func (g *Guild) IconURL(asFormat ImageFormat) string {
	if g.Icon != "" {
		if asFormat == "" {
			asFormat = "png"
			if strings.HasPrefix(g.Icon, "a_") {
				asFormat = "gif"
			}
		}

		suffix := fmt.Sprintf("%s.%s", g.Icon, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/icons/%d/%s", g.ID, suffix)
	}
	return ""
}

// SplashURL returns the URL for the guild splash
func (g *Guild) SplashURL(asFormat ImageFormat) string {
	if g.Splash != "" {
		if asFormat == "" {
			asFormat = "png"
		}

		suffix := fmt.Sprintf("%s.%s", g.Splash, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/splashes/%d/%s", g.ID, suffix)
	}
	return ""
}

// DiscoverySplashURL returns the URL for the guild discovery splash
func (g *Guild) DiscoverySplashURL(asFormat ImageFormat) string {
	if g.DiscoverySplash != "" {
		if asFormat == "" {
			asFormat = "png"
		}

		suffix := fmt.Sprintf("%s.%s", g.DiscoverySplash, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/discovery-splashes/%d/%s", g.ID, suffix)
	}
	return ""
}

// BannerURL returns the URL for the guild banner
func (g *Guild) BannerURL(asFormat ImageFormat) string {
	if g.Banner != "" {
		if asFormat == "" {
			asFormat = "png"
			if strings.HasPrefix(g.Banner, "a_") {
				asFormat = "gif"
			}
		}

		suffix := fmt.Sprintf("%s.%s", g.Banner, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/banners/%d/%s", g.ID, suffix)
	}
	return ""
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

// Compute creates a new channel flags structure (discord.SystemChannelFlags) from the given channel flags
func (f SystemChannelFlags) Compute(systemChannelFlags ...SystemChannelFlags) SystemChannelFlags {
	return flags.Compute(systemChannelFlags...)
}

// Add adds the given system channel flags (discord.SystemChannelFlags)
func (f SystemChannelFlags) Add(systemChannelFlags ...SystemChannelFlags) SystemChannelFlags {
	return flags.Add(f, systemChannelFlags...)
}

// Remove removes the given system channel flags (discord.SystemChannelFlags)
func (f SystemChannelFlags) Remove(systemChannelFlags ...SystemChannelFlags) SystemChannelFlags {
	return flags.Remove(f, systemChannelFlags...)
}

// Toggle toggles the given system channel flags (discord.SystemChannelFlags)
func (f SystemChannelFlags) Toggle(systemChannelFlags ...SystemChannelFlags) SystemChannelFlags {
	return flags.Toggle(f, systemChannelFlags...)
}

// Has checks if all the given system channel flags (discord.SystemChannelFlags) are set
func (f SystemChannelFlags) Has(systemChannelFlags ...SystemChannelFlags) bool {
	return flags.Has(f, systemChannelFlags...)
}

// HasAny checks if any of the given system channel flags (discord.SystemChannelFlags) is set
func (f SystemChannelFlags) HasAny(systemChannelFlags ...SystemChannelFlags) bool {
	return flags.HasAny(f, systemChannelFlags...)
}

// HasNot checks if all the given system channel flags (discord.SystemChannelFlags) are set
func (f SystemChannelFlags) HasNot(systemChannelFlags ...SystemChannelFlags) bool {
	return flags.HasNot(f, systemChannelFlags...)
}

// HasNotAny checks if any of the given system channel flags (discord.SystemChannelFlags) is not set
func (f SystemChannelFlags) HasNotAny(systemChannelFlags ...SystemChannelFlags) bool {
	return flags.HasNotAny(f, systemChannelFlags...)
}

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
	ChannelID   Snowflake `json:"channel_id"`
	Description string    `json:"description"`
	EmojiID     Snowflake `json:"emoji_id,omitempty"`
	EmojiName   string    `json:"emoji_name,omitempty"`
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
	ID                       Snowflake      `json:"id"`
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

// CreatedAt returns the creation time of the guild preview (discord.GuildPreview)
func (p *GuildPreview) CreatedAt() time.Time {
	return p.ID.CreatedAt()
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

	UserID         Snowflake `json:"-"`
	AuditLogReason string    `json:"-"`
}

// GuildPrune represents how many members (discord.Member) will get pruned when performing a prune
// https://discord.com/developers/docs/resources/guild#get-guild-prune-count
type GuildPrune struct {
	Pruned int `json:"pruned"`
}

// BeginGuildPrune represents the payload to send to Discord to begin a guild prune operation (discord.GuildPrune)
// https://discord.com/developers/docs/resources/guild#begin-guild-prune-json-params
type BeginGuildPrune struct {
	Days              *int            `json:"days,omitempty"`
	ComputePruneCount *bool           `json:"compute_prune_count,omitempty"`
	IncludeRoles      ArraySnowflakes `json:"include_roles,omitempty"`

	AuditLogReason string `json:"-"`
}

// WidgetSetting represents the channel (discord.Channel) used and whether the widget (discord.Widget) of the guild (discord.Guild) is enabled
// https://discord.com/developers/docs/resources/guild#guild-widget-settings-object-guild-widget-settings-structure
type WidgetSetting struct {
	Enabled   bool      `json:"enabled"`
	ChannelID Snowflake `json:"channel_id"`
}

// Widget represents the guild's (discord.Guild) widget
// https://discord.com/developers/docs/resources/guild#guild-widget-object-guild-widget-structure
type Widget struct {
	ID            Snowflake `json:"id"`
	Name          string    `json:"name"`
	InstantInvite string    `json:"instant_invite"`
	Channels      []Channel `json:"channels"`
	Members       []Member  `json:"members"`
	PresenceCount int       `json:"presence_count"`
}

// ModifyGuildWidget represents the payload to send to Discord to modify the guild widget (discord.Widget)
// https://discord.com/developers/docs/resources/guild#guild-widget-settings-object-guild-widget-settings-structure
type ModifyGuildWidget struct {
	Enabled   *bool      `json:"enabled,omitempty"`
	ChannelID *Snowflake `json:"channel_id,omitempty"`

	AuditLogReason string `json:"-"`
}

// Onboarding represents the onboarding flow for a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#guild-onboarding-object-guild-onboarding-structure
type Onboarding struct {
	GuildID           Snowflake          `json:"guild_id"`
	Prompts           []OnboardingPrompt `json:"prompts"`
	DefaultChannelIDs ArraySnowflakes    `json:"default_channel_ids"`
	Enabled           bool               `json:"enabled"`
}

// OnboardingPrompt represents a prompt shown during onboarding (discord.Onboarding) and in customize community
// https://discord.com/developers/docs/resources/guild#guild-onboarding-object-onboarding-prompt-structure
type OnboardingPrompt struct {
	ID           Snowflake                `json:"id"`
	Type         OnboardingPromptType     `json:"type"`
	Options      []OnboardingPromptOption `json:"options"`
	Title        string                   `json:"title"`
	SingleSelect bool                     `json:"single_select"`
	Required     bool                     `json:"required"`
	InOnboarding bool                     `json:"in_onboarding"`
}

// CreatedAt returns the creation time of the onboarding prompt (discord.OnboardingPrompt)
func (p *OnboardingPrompt) CreatedAt() time.Time {
	return p.ID.CreatedAt()
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
	ID          Snowflake       `json:"id"`
	ChannelIDs  ArraySnowflakes `json:"channel_ids"`
	RoleIDs     ArraySnowflakes `json:"role_ids"`
	Emoji       *Emoji          `json:"emoji"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
}

// CreatedAt returns the creation time of the onboarding prompt option (discord.OnboardingPromptOption)
func (o *OnboardingPromptOption) CreatedAt() time.Time {
	return o.ID.CreatedAt()
}

// PartialGuild represents a guild (discord.Guild) with only a few fields
// https://discord.com/developers/docs/resources/user#get-current-user-guilds-example-partial-guild
type PartialGuild struct {
	ID          Snowflake      `json:"id"`
	Name        string         `json:"name"`
	Icon        string         `json:"icon"`
	Owner       bool           `json:"owner"`
	Permissions uint64         `json:"permissions"`
	Features    []GuildFeature `json:"features"`
}

// CreatedAt returns the creation time of the partial guild (discord.PartialGuild)
func (g *PartialGuild) CreatedAt() time.Time {
	return g.ID.CreatedAt()
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
	GuildFeatureRaidAlertsDisabled                    GuildFeature = "RAID_ALERTS_DISABLED"
	GuildFeatureRoleIcons                             GuildFeature = "ROLE_ICONS"
	GuildFeatureRoleSubscriptionsAvailableForPurchase GuildFeature = "ROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE"
	GuildFeatureRoleSubscriptionsEnabled              GuildFeature = "ROLE_SUBSCRIPTIONS_ENABLED"
	GuildFeatureTicketedEventsEnabled                 GuildFeature = "TICKETED_EVENTS_ENABLED"
	GuildFeatureVanityURL                             GuildFeature = "VANITY_URL"
	GuildFeatureVerified                              GuildFeature = "VERIFIED"
	GuildFeatureVIPRegions                            GuildFeature = "VIP_REGIONS"
	GuildFeatureWelcomeScreenEnabled                  GuildFeature = "WELCOME_SCREEN_ENABLED"
)

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
	AFKChannelID                *Snowflake                `json:"afk_channel_id,omitempty"`
	AFKTimeout                  *int                      `json:"afk_timeout,omitempty"`
	SystemChannelID             *Snowflake                `json:"system_channel_id,omitempty"`
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
	AFKChannelID                *Snowflake                `json:"afk_channel_id,omitempty"`
	AFKTimeout                  *int                      `json:"afk_timeout,omitempty"`
	Icon                        *string                   `json:"icon,omitempty"`
	OwnerID                     *Snowflake                `json:"owner_id,omitempty"`
	Splash                      *string                   `json:"splash,omitempty"`
	DiscoverySplash             *string                   `json:"discovery_splash,omitempty"`
	Banner                      *string                   `json:"banner,omitempty"`
	SystemChannelID             *Snowflake                `json:"system_channel_id,omitempty"`
	SystemChannelFlags          *SystemChannelFlags       `json:"system_channel_flags,omitempty"`
	RulesChannelID              *Snowflake                `json:"rules_channel_id,omitempty"`
	PublicUpdatesChannelID      *Snowflake                `json:"public_updates_channel_id,omitempty"`
	PreferredLocale             *Locale                   `json:"preferred_locale,omitempty"`
	Features                    []GuildFeature            `json:"features,omitempty"`
	Description                 *string                   `json:"description,omitempty"`
	PremiumProgressBarEnabled   *bool                     `json:"premium_progress_bar_enabled,omitempty"`
	SafetyAlertsChannelID       *Snowflake                `json:"safety_alerts_channel_id,omitempty"`

	AuditLogReason string `json:"-"`
}

// ModifyGuildMFALevel represents the payload to send to Discord to modify the MFA level for a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#modify-guild-mfa-level-json-params
type ModifyGuildMFALevel struct {
	Level MFALevel `json:"level"`

	AuditLogReason string `json:"-"`
}
