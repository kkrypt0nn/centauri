package discord

import (
	"github.com/kkrypt0nn/centauri/oauth2"
	"time"
)

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
	SyncedAt          time.Time                 `json:"synced_at,omitempty"`
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
	Icon        string `json:"icon,omitempty"`
	Description string `json:"description"`
	Bot         *User  `json:"bot,omitempty"`
}

type Member struct {
	User                       *User       `json:"user,omitempty"`
	Nick                       string      `json:"nick,omitempty"`
	Avatar                     string      `json:"avatar,omitempty"`
	Roles                      []string    `json:"roles"`
	JoinedAt                   time.Time   `json:"joined_at"`
	PremiumSince               time.Time   `json:"premium_since,omitempty"`
	Deaf                       bool        `json:"deaf"`
	Mute                       bool        `json:"mute"`
	Flags                      MemberFlags `json:"flags"`
	Pending                    bool        `json:"pending,omitempty"`
	Permissions                uint64      `json:"permissions,string,omitempty"`
	CommunicationDisabledUntil time.Time   `json:"communication_disabled_until,omitempty"`
}

type MemberFlags uint64

const MemberFlagsNone MemberFlags = 0
const (
	MemberFlagDidRejoin MemberFlags = 1 << iota
	MemberFlagCompletedOnboarding
	MemberFlagBypassesVerification
	MemberFlagStartedOnboarding
)
