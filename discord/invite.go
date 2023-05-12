package discord

import "time"

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
	StageInstance            *InviteStageInstance `json:"stage_instance,omitempty"`
	GuildScheduledEvent      *GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`
}

type InviteWithMetadata struct {
	Invite
	Uses      int       `json:"uses"`
	MaxUses   int       `json:"max_uses"`
	MaxAge    int       `json:"max_age"`
	Temporary bool      `json:"temporary"`
	CreatedAt time.Time `json:"created_at"`
}

type PartialInvite struct {
	Code string `json:"code"`
	Uses int    `json:"uses"`
}

type InviteTargetType int

const (
	InviteTargetTypeStream InviteTargetType = 1 + iota
	InviteTargetTypeEmbeddedApplication
)

type InviteStageInstance struct {
	Members          []Member `json:"members"`
	ParticipantCount int      `json:"participant_count"`
	SpeakerCount     int      `json:"speaker_count"`
	Topic            string   `json:"topic"`
}
