package discord

// User represents a Discord user or a Discord bot
// https://discord.com/developers/docs/resources/user#user-object
type User struct {
	// ID is the user's id
	ID string `json:"id"`
	// Username is the user's username, not unique across the platform
	Username string `json:"username"`
	// Avatar is the user's avatar hash
	Avatar string `json:"avatar"`
	// Discriminator is the user's 4-digit discord-tag
	Discriminator string `json:"discriminator"`
	// PublicFlags are the public UserFlags on a user's account
	PublicFlags UserFlags `json:"public_flags,omitempty"`
	// Flags are the UserFlags on a user's account
	Flags UserFlags `json:"flags,omitempty"`
	// Bot shows whether the user belongs to an OAuth2 application
	Bot bool `json:"bot,omitempty"`
	// Banner is the user's banner hash
	Banner string `json:"banner,omitempty"`
	// AccentColor is the user's banner color encoded as an integer representation of hexadecimal color code
	AccentColor int `json:"accent_color,omitempty"`
	// Locale is the user's chosen language option
	Locale string `json:"locale,omitempty"`
	// MFAEnabled shows whether the user has two factor enabled on their account
	MFAEnabled bool `json:"mfa_enabled,omitempty"`
	// PremiumType is the type of Nitro subscription on a user's account
	PremiumType int `json:"premium_type,omitempty"`
	// Email is the user's email
	Email string `json:"email,omitempty"`
	// Verified shows whether the email on this account has been verified
	Verified bool `json:"verified,omitempty"`

	// These are undocumented fields, subject to change

	// BannerColor is the user's banner color as hexadecimal value
	BannerColor string `json:"banner_color,omitempty"`
	GlobalName  string `json:"global_name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	// AvatarDecoration is the user's avatar decoration
	AvatarDecoration  string        `json:"avatar_decoration,omitempty"`
	PurchasedFlags    int           `json:"purchased_flags,omitempty"`
	PremiumUsageFlags int           `json:"premium_usage_flags,omitempty"`
	LinkedUsers       []interface{} `json:"linked_users,omitempty"`
	// NSFWAllowed shows whether the user allows NSFW content in their settings
	NSFWAllowed bool `json:"nsfw_allowed,omitempty"`
	// Biography is the user's biography or "About me"
	Biography string `json:"bio,omitempty"`
	// Phone is the user's phone number
	Phone string `json:"phone,omitempty"`
}

// UserFlags are the flags a discord.User may have
// https://discord.com/developers/docs/resources/user#user-object-user-flags
type UserFlags uint64

// NoFlags is when a discord.User has no flags on their profile
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
	// LikelySpammer represents a discord.User that has been flagged by Discord for likely being a spammer
	LikelySpammer // Undocumented, subject to change
	_
	// ActiveDeveloper represents Active Developers (https://support-dev.discord.com/hc/en-us/articles/10113997751447)
	ActiveDeveloper
)
