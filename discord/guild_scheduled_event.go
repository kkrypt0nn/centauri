package discord

import "time"

type GuildScheduledEvent struct {
	ID                 string                             `json:"id"`
	GuildID            string                             `json:"guild_id"`
	ChannelID          string                             `json:"channel_id"`
	CreatorID          string                             `json:"creator_id"`
	Name               string                             `json:"name"`
	Description        string                             `json:"description"`
	ScheduledStartTime time.Time                          `json:"scheduled_start_time"`
	ScheduledEndTime   time.Time                          `json:"scheduled_end_time"`
	PrivacyLevel       GuildScheduledEventPrivacyLevel    `json:"privacy_level"`
	Status             GuildScheduledEventStatus          `json:"status"`
	EntityType         GuildScheduledEventEntityType      `json:"entity_type"`
	EntityID           string                             `json:"entity_id"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata"`
	Creator            *User                              `json:"creator"`
	UserCount          int                                `json:"user_count"`
	Image              string                             `json:"image"`
}

type GuildScheduledEventPrivacyLevel int

const (
	_ GuildScheduledEventPrivacyLevel = 1 + iota
	GuildScheduledEventPrivacyLevelGuildOnly
)

type GuildScheduledEventStatus int

const (
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = 1 + iota
	GuildScheduledEventStatusActive
	GuildScheduledEventStatusCompleted
	GuildScheduledEventStatusCancelled
)

type GuildScheduledEventEntityType int

const (
	GuildScheduledEventEntityTypeStageInstance = 1 + iota
	GuildScheduledEventEntityTypeVoice
	GuildScheduledEventEntityTypeExternal
)

type GuildScheduledEventEntityMetadata struct {
	Location string `json:"location,omitempty"`
}
