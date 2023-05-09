package discord

// User represents a Discord user or a Discord bot
// https://discord.com/developers/docs/resources/user#user-object
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
	Locale        string      `json:"locale,omitempty"`
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

const NoFlags UserFlags = 0
const (
	Staff UserFlags = 1 << iota
	Partner
	HypeSquadEvents
	BugHunterLevel1
	MfaSms                // Undocumented
	PremiumPromoDismissed // Undocumented
	HypeSquadBravery
	HypeSquadBrilliance
	HypeSquadBalance
	EarlyNitroSupporter
	TeamPseudoUser
	InternalApplication     // Undocumented
	System                  // Undocumented
	HasUnreadUrgentMessages // Undocumented
	BugHunterLevel2
	UnderageDeleted // Undocumented
	VerifiedBot
	VerifiedBotDeveloper
	CertifiedModerator
	BotHttpInteractions
	LikelySpammer  // Undocumented
	DisablePremium // Undocumented
	ActiveDeveloper
)

// PremiumType represents the premium type a User has
type PremiumType int

const (
	NoPremium PremiumType = iota
	NitroClassic
	Nitro
	NitroBasic
)

// PremiumUsageFlags represents a user's usage flags of premium features
type PremiumUsageFlags int

const (
	PremiumDiscriminator PremiumUsageFlags = 1 << iota
	AnimatedAvatar
	ProfileBanner
)

// PurchasedFlags represents a user's purchased state
type PurchasedFlags int

const (
	NitroClassicPurchased PurchasedFlags = 1 << iota
	NitroPurchased
	GuildBoostPurchased
)
