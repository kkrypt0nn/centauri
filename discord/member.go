package discord

import (
	"fmt"
	"strings"
	"time"
)

// Member represents a user (discord.User) that is a member/that has joined a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#guild-member-object-guild-member-structure
type Member struct {
	User                       *User       `json:"user,omitempty"`
	Nick                       string      `json:"nick,omitempty"`
	Avatar                     string      `json:"avatar,omitempty"`
	Roles                      []string    `json:"roles,omitempty"`
	JoinedAt                   time.Time   `json:"joined_at"`
	PremiumSince               *time.Time  `json:"premium_since,omitempty"`
	Deaf                       bool        `json:"deaf"`
	Mute                       bool        `json:"mute"`
	Flags                      MemberFlags `json:"flags"`
	Pending                    bool        `json:"pending,omitempty"`
	Permissions                Permissions `json:"permissions,string,omitempty"`
	CommunicationDisabledUntil *time.Time  `json:"communication_disabled_until,omitempty"`

	// Usually that's set by Centauri
	GuildID string `json:"guild_id,omitempty"`
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
		return fmt.Sprintf("https://cdn.discordapp.com/guilds/%s/users/%s/avatars/%s", m.GuildID, m.User.ID, suffix)
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

// ModifyGuildMember represents the payload to send to Discord to modify a guild member (discord.Member)
// https://discord.com/developers/docs/resources/guild#modify-guild-member-json-params
type ModifyGuildMember struct {
	Nick                       *string      `json:"nick,omitempty"`
	Roles                      []string     `json:"roles,omitempty"`
	Deaf                       *bool        `json:"deaf,omitempty"`
	Mute                       *bool        `json:"mute,omitempty"`
	ChannelID                  *string      `json:"channel_id,omitempty"`
	CommunicationDisabledUntil *time.Time   `json:"communication_disabled_until,omitempty"`
	Flags                      *MemberFlags `json:"flags,omitempty"`

	AuditLogReason string `json:"-"`
}
