package discord

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/utils/flags"
	"strings"
	"time"
)

// Member represents a user (discord.User) that is a member/that has joined a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#guild-member-object-guild-member-structure
type Member struct {
	User                       *User           `json:"user,omitempty"`
	Nick                       string          `json:"nick,omitempty"`
	Avatar                     string          `json:"avatar,omitempty"`
	Roles                      ArraySnowflakes `json:"roles,omitempty"`
	JoinedAt                   time.Time       `json:"joined_at"`
	PremiumSince               *time.Time      `json:"premium_since,omitempty"`
	Deaf                       bool            `json:"deaf"`
	Mute                       bool            `json:"mute"`
	Flags                      MemberFlags     `json:"flags"`
	Pending                    bool            `json:"pending,omitempty"`
	Permissions                Permissions     `json:"permissions,omitempty"`
	CommunicationDisabledUntil *time.Time      `json:"communication_disabled_until,omitempty"`

	// Usually that's set by Centauri
	GuildID Snowflake `json:"guild_id,omitempty"`
}

// CreatedAt returns the creation time of the member (discord.Member)
func (m *Member) CreatedAt() time.Time {
	return m.User.ID.CreatedAt()
}

// GuildAvatarURL returns the avatar URL specific to the guild (discord.Guild)
func (m *Member) GuildAvatarURL(asFormat ImageFormat) string {
	if m.Avatar != "" {
		if asFormat == "" {
			asFormat = "png"
			if strings.HasPrefix(m.Avatar, "a_") {
				asFormat = "gif"
			}
		}

		suffix := fmt.Sprintf("%s.%s", m.Avatar, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/guilds/%d/users/%d/avatars/%s", m.GuildID, m.User.ID, suffix)
	}
	return ""
}

// EffectiveAvatarURL returns the avatar URL specific to the guild (discord.Guild) if set, otherwise the user's default avatar
func (m *Member) EffectiveAvatarURL(asFormat ImageFormat) string {
	if guildAvatarUrl := m.GuildAvatarURL(asFormat); guildAvatarUrl != "" {
		return guildAvatarUrl
	}
	return m.User.AvatarURL(asFormat)
}

// EffectiveName returns either the nickname of the member (discord.Member) or the username/global name of the user (discord.User) which is displayed in the client
func (m *Member) EffectiveName() string {
	if m.Nick != "" {
		return m.Nick
	}
	return m.User.EffectiveName()
}

// MemberFlags represents the flags of a member (discord.Member)
// https://discord.com/developers/docs/resources/guild#guild-member-object-guild-member-flags
type MemberFlags uint64

const MemberFlagsNone MemberFlags = 0
const (
	MemberFlagDidRejoin MemberFlags = 1 << iota
	MemberFlagCompletedOnboarding
	MemberFlagBypassesVerification
	MemberFlagStartedOnboarding
)

// ComputeMemberFlags creates a new member flags structure (discord.MemberFlags) from the given member flags
func ComputeMemberFlags(memberFlags ...MemberFlags) MemberFlags {
	return flags.Compute(memberFlags...)
}

// Add adds the given member flags (discord.MemberFlags)
func (f MemberFlags) Add(memberFlags ...MemberFlags) MemberFlags {
	return flags.Add(f, memberFlags...)
}

// Remove removes the given member flags (discord.MemberFlags)
func (f MemberFlags) Remove(memberFlags ...MemberFlags) MemberFlags {
	return flags.Remove(f, memberFlags...)
}

// Toggle toggles the given member flags (discord.MemberFlags)
func (f MemberFlags) Toggle(memberFlags ...MemberFlags) MemberFlags {
	return flags.Toggle(f, memberFlags...)
}

// Has checks if all the given member flags (discord.MemberFlags) are set
func (f MemberFlags) Has(memberFlags ...MemberFlags) bool {
	return flags.Has(f, memberFlags...)
}

// HasAny checks if any of the given member flags (discord.MemberFlags) is set
func (f MemberFlags) HasAny(memberFlags ...MemberFlags) bool {
	return flags.HasAny(f, memberFlags...)
}

// HasNot checks if all the given member flags (discord.MemberFlags) are not set
func (f MemberFlags) HasNot(memberFlags ...MemberFlags) bool {
	return flags.HasNot(f, memberFlags...)
}

// HasNotAny checks if any of the given member flags (discord.MemberFlags) is not set
func (f MemberFlags) HasNotAny(memberFlags ...MemberFlags) bool {
	return flags.HasNotAny(f, memberFlags...)
}

// ModifyGuildMember represents the payload to send to Discord to modify a guild member (discord.Member)
// https://discord.com/developers/docs/resources/guild#modify-guild-member-json-params
type ModifyGuildMember struct {
	Nick                       *string         `json:"nick,omitempty"`
	Roles                      ArraySnowflakes `json:"roles,omitempty"`
	Deaf                       *bool           `json:"deaf,omitempty"`
	Mute                       *bool           `json:"mute,omitempty"`
	ChannelID                  *Snowflake      `json:"channel_id,omitempty"`
	CommunicationDisabledUntil *time.Time      `json:"communication_disabled_until,omitempty"`
	Flags                      *MemberFlags    `json:"flags,omitempty"`

	AuditLogReason string `json:"-"`
}
