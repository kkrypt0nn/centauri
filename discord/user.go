package discord

// User represents a Discord user or a Discord bot
// https://discord.com/developers/docs/resources/user#user-object-user-structure
type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	GlobalName string `json:"global_name,omitempty"`
	Avatar     string `json:"avatar"`
	// Deprecated: Will return "0" once a user has migrated to the new usernames
	Discriminator string      `json:"discriminator"`
	PublicFlags   UserFlags   `json:"public_flags,omitempty"`
	Flags         UserFlags   `json:"flags,omitempty"`
	Bot           bool        `json:"bot,omitempty"`
	Banner        string      `json:"banner,omitempty"`
	AccentColor   int         `json:"accent_color,omitempty"`
	Locale        Locale      `json:"locale,omitempty"`
	MFAEnabled    bool        `json:"mfa_enabled,omitempty"`
	PremiumType   PremiumType `json:"premium_type,omitempty"`
	Email         string      `json:"email,omitempty"`
	Verified      bool        `json:"verified,omitempty"`

	// These are undocumented fields, subject to change
	BannerColor       string            `json:"banner_color,omitempty"`
	DisplayName       string            `json:"display_name,omitempty"`
	AvatarDecoration  string            `json:"avatar_decoration,omitempty"`
	PurchasedFlags    PurchasedFlags    `json:"purchased_flags,omitempty"`
	PremiumUsageFlags PremiumUsageFlags `json:"premium_usage_flags,omitempty"`
	LinkedUsers       []interface{}     `json:"linked_users,omitempty"`
	NSFWAllowed       bool              `json:"nsfw_allowed,omitempty"`
	Biography         string            `json:"bio,omitempty"`
	Phone             string            `json:"phone,omitempty"`
}

// UserFlags are the flags a User may have
// https://discord.com/developers/docs/resources/user#user-object-user-flags
type UserFlags uint64

const UserFlagNone UserFlags = 0
const (
	UserFlagStaff UserFlags = 1 << iota
	UserFlagPartner
	UserFlagHypeSquadEvents
	UserFlagBugHunterLevel1
	UserFlagMfaSms                // Undocumented
	UserFlagPremiumPromoDismissed // Undocumented
	UserFlagHypeSquadBravery
	UserFlagHypeSquadBrilliance
	UserFlagHypeSquadBalance
	UserFlagEarlyNitroSupporter
	UserFlagTeamPseudoUser
	UserFlagInternalApplication     // Undocumented
	UserFlagSystem                  // Undocumented
	UserFlagHasUnreadUrgentMessages // Undocumented
	UserFlagBugHunterLevel2
	UserFlagUnderageDeleted // Undocumented
	UserFlagVerifiedBot
	UserFlagVerifiedBotDeveloper
	UserFlagCertifiedModerator
	UserFlagBotHttpInteractions
	UserFlagLikelySpammer  // Undocumented
	UserFlagDisablePremium // Undocumented
	UserFlagActiveDeveloper
)

// PremiumType represents the premium type a User has
type PremiumType int

const (
	PremiumTypeNone PremiumType = iota
	PremiumTypeNitroClassic
	PremiumTypeNitro
	PremiumTypeNitroBasic
)

// PremiumUsageFlags represents a user's usage flags of premium features
type PremiumUsageFlags int

const (
	PremiumUsageFlagPremiumDiscriminator PremiumUsageFlags = 1 << iota
	PremiumUsageFlagAnimatedAvatar
	PremiumUsageFlagProfileBanner
)

// PurchasedFlags represents a user's purchased state
type PurchasedFlags int

const (
	PurchasedFlagNitroClassic PurchasedFlags = 1 << iota
	PurchasedFlagNitro
	PurchasedFlagGuildBoost
)
