package discord

import "time"

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
