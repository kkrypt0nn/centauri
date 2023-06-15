package discord

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/utils/flags"
	"strconv"
	"strings"
	"time"
)

// User represents a Discord user or a Discord bot
// https://discord.com/developers/docs/resources/user#user-object-user-structure
type User struct {
	ID         Snowflake `json:"id"`
	Username   string    `json:"username"`
	GlobalName string    `json:"global_name,omitempty"`
	Avatar     string    `json:"avatar"`
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

// CreatedAt returns the creation time of the user (discord.User)
func (u *User) CreatedAt() time.Time {
	return u.ID.CreatedAt()
}

// BannerURL returns the banner URL of the user (discord.User)
func (u *User) BannerURL(asFormat ImageFormat) string {
	if u.Banner != "" {
		if asFormat == "" {
			asFormat = "png"
			if strings.HasPrefix(u.Banner, "a_") {
				asFormat = "gif"
			}
		}

		suffix := fmt.Sprintf("%s.%s", u.Banner, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/banners/%d/%s", u.ID, suffix)
	}
	return ""
}

// DefaultAvatarURL returns the default avatar URL of the user (discord.User)
func (u *User) DefaultAvatarURL() string {
	var index int
	if u.Discriminator == "0" {
		// User has migrated to the new username system
		index = int((u.ID >> 22) % 6)
	} else {
		// User is still on the old username and discriminator system
		discriminatorInt, err := strconv.Atoi(u.Discriminator)
		if err != nil {
			return err.Error()
		}
		index = discriminatorInt % 5
	}
	return fmt.Sprintf("https://cdn.discordapp.com/embed/avatars/%d.png", index)
}

// AvatarURL returns the avatar URL of the user (discord.User)
func (u *User) AvatarURL(asFormat ImageFormat) string {
	if u.Avatar != "" {
		if asFormat == "" {
			asFormat = "png"
			if strings.HasPrefix(u.Avatar, "a_") {
				asFormat = "gif"
			}
		}

		suffix := fmt.Sprintf("%s.%s", u.Avatar, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/avatars/%d/%s", u.ID, suffix)
	}
	return u.DefaultAvatarURL()
}

// EffectiveName returns either the global name or the username of the user (discord.User) which is displayed in the client
func (u *User) EffectiveName() string {
	if u.GlobalName != "" {
		return u.GlobalName
	}
	return u.Username
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

// Compute creates a new user flags structure (discord.UserFlags) from the given user flags
func (f UserFlags) Compute(userFlags ...UserFlags) UserFlags {
	return flags.Compute(userFlags...)
}

// Add adds the given user flags (discord.UserFlags)
func (f UserFlags) Add(userFlags ...UserFlags) UserFlags {
	return flags.Add(f, userFlags...)
}

// Remove removes the given user flags (discord.UserFlags)
func (f UserFlags) Remove(userFlags ...UserFlags) UserFlags {
	return flags.Remove(f, userFlags...)
}

// Toggle toggles the given user flags (discord.UserFlags)
func (f UserFlags) Toggle(userFlags ...UserFlags) UserFlags {
	return flags.Toggle(f, userFlags...)
}

// Has checks if all the given user flags (discord.UserFlags) are set
func (f UserFlags) Has(userFlags ...UserFlags) bool {
	return flags.Has(f, userFlags...)
}

// HasAny checks if any of the given user flags (discord.UserFlags) is set
func (f UserFlags) HasAny(userFlags ...UserFlags) bool {
	return flags.HasAny(f, userFlags...)
}

// HasNot checks if all the given user flags (discord.UserFlags) are not set
func (f UserFlags) HasNot(userFlags ...UserFlags) bool {
	return flags.HasNot(f, userFlags...)
}

// HasNotAny checks if any of the given user flags (discord.UserFlags) is not set
func (f UserFlags) HasNotAny(userFlags ...UserFlags) bool {
	return flags.HasNotAny(f, userFlags...)
}

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

// ModifyCurrentUser represents the payload to send to Discord to modify the current user (discord.User)
// https://discord.com/developers/docs/resources/user#modify-current-user-json-params
type ModifyCurrentUser struct {
	Username *string `json:"username,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}
