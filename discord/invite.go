package discord

import "time"

// Invite represents a code that when used, adds a user to a guild or group DM channel
// https://discord.com/developers/docs/resources/invite#invite-object-invite-structure
type Invite struct {
	Code                     string               `json:"code"`
	Guild                    *PartialGuild        `json:"guild,omitempty"`
	Channel                  *Channel             `json:"channel,omitempty"`
	Inviter                  *User                `json:"inviter,omitempty"`
	TargetType               InviteTargetType     `json:"target_type,omitempty"`
	TargetUser               *User                `json:"target_user,omitempty"`
	TargetApplication        *Application         `json:"target_application,omitempty"`
	ApproximatePresenceCount int                  `json:"approximate_presence_count"`
	ApproximateMemberCount   int                  `json:"approximate_member_count"`
	ExpiresAt                *time.Time           `json:"expires_at"`
	GuildScheduledEvent      *GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`
}

// CreateChannelInvite represents the payload to send to Discord to create an invite (discord.Invite) for a channel (discord.Channel)
type CreateChannelInvite struct {
	MaxAge              *int              `json:"max_age,omitempty"`
	MaxUses             *int              `json:"max_uses,omitempty"`
	Temporary           *bool             `json:"temporary,omitempty"`
	Unique              *bool             `json:"unique,omitempty"`
	TargetType          *InviteTargetType `json:"target_type,omitempty"`
	TargetUserID        string            `json:"target_user_id,omitempty"`
	TargetApplicationID string            `json:"target_application_id,omitempty"`

	AuditLogReason string `json:"-"`
}

// InviteWithMetadata represents an invite (discord.Invite) with additional metadata
// https://discord.com/developers/docs/resources/invite#invite-metadata-object-invite-metadata-structure
type InviteWithMetadata struct {
	Invite
	Uses      int       `json:"uses"`
	MaxUses   int       `json:"max_uses"`
	MaxAge    int       `json:"max_age"`
	Temporary bool      `json:"temporary"`
	CreatedAt time.Time `json:"created_at"`
}

// PartialInvite represents a partial invite (discord.Invite) - for so-called vanity invites
// https://discord.com/developers/docs/resources/guild#get-guild-vanity-url-example-partial-invite-object
type PartialInvite struct {
	Code string `json:"code"`
	Uses int    `json:"uses"`
}

// InviteTargetType represents the target type of the invite (discord.Invite)
// https://discord.com/developers/docs/resources/invite#invite-object-invite-target-types
type InviteTargetType int

const (
	InviteTargetTypeStream InviteTargetType = 1 + iota
	InviteTargetTypeEmbeddedApplication
)
