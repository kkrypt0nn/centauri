package discord

// User represents a Discord user or a Discord bot
// https://discord.com/developers/docs/resources/user#user-object
type User struct {
	ID            string    `json:"id"`
	Username      string    `json:"username"`
	Avatar        string    `json:"avatar"`
	Discriminator string    `json:"discriminator"`
	PublicFlags   UserFlags `json:"public_flags,omitempty"`
	Flags         UserFlags `json:"flags,omitempty"`
	Bot           bool      `json:"bot,omitempty"`
	Banner        string    `json:"banner,omitempty"`
	BannerColor   string    `json:"banner_color,omitempty"`
	AccentColor   int       `json:"accent_color,omitempty"`
	Locale        string    `json:"locale,omitempty"`
	MFAEnabled    bool      `json:"mfa_enabled,omitempty"`
	PremiumType   int       `json:"premium_type,omitempty"`
	Email         string    `json:"email,omitempty"`
	Verified      bool      `json:"verified,omitempty"`

	// These are undocumented fields, subject to change
	GlobalName        string        `json:"global_name,omitempty"`
	DisplayName       string        `json:"display_name,omitempty"`
	AvatarDecoration  string        `json:"avatar_decoration,omitempty"`
	PurchasedFlags    int           `json:"purchased_flags,omitempty"`
	PremiumUsageFlags int           `json:"premium_usage_flags,omitempty"`
	LinkedUsers       []interface{} `json:"linked_users,omitempty"`
	NSFWAllowed       bool          `json:"nsfw_allowed,omitempty"`
	Biography         string        `json:"bio,omitempty"`
	Phone             string        `json:"phone,omitempty"`
}

// UserFlags are the flags a user, or bot, may have
// https://discord.com/developers/docs/resources/user#user-object-user-flags
type UserFlags uint64

// NoFlags is when a user has no flags on their profile
const NoFlags UserFlags = 0

const (
	// Staff represents Discord Employees
	Staff UserFlags = 1 << iota
	// Partner represents Partnered Server Owners
	Partner
	// HypeSquadEvents represents HypeSquad Events Members
	HypeSquadEvents
	// BugHunterLevel1 represents Bug Hunters Level 1
	BugHunterLevel1
	_
	_
	// HypeSquadBravery represents House Bravery Members
	HypeSquadBravery
	// HypeSquadBrilliance represents House Brilliance Members
	HypeSquadBrilliance
	// HypeSquadBalance represents House Balance Members
	HypeSquadBalance
	// EarlyNitroSupporter represents Early Nitro Supporters
	EarlyNitroSupporter
	// TeamPseudoUser represents that the User is a team (https://discord.com/developers/docs/topics/teams)
	TeamPseudoUser
	_
	_
	_
	// BugHunterLevel2 represents Bug Hunters Level 2
	BugHunterLevel2
	_
	// VerifiedBot represents Verified Bots
	VerifiedBot
	// VerifiedBotDeveloper represents Early Verified Bot Developers
	VerifiedBotDeveloper
	// CertifiedModerator represents Moderator Programs Alumnis
	CertifiedModerator
	// BotHttpInteractions represents a bot that only HTTP interactions (https://discord.com/developers/docs/interactions/receiving-and-responding#receiving-an-interaction) and is shown in the online member list
	BotHttpInteractions
	// LikelySpammer represents a user that has been flagged by Discord for being likely a spammer
	LikelySpammer // Undocumented, subject to change
	_
	// ActiveDeveloper represents Active Developers (https://support-dev.discord.com/hc/en-us/articles/10113997751447)
	ActiveDeveloper
)
