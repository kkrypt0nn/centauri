package discord

import "time"

// GuildScheduledEvent represents a scheduled event in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-structure
type GuildScheduledEvent struct {
	ID                 string                             `json:"id"`
	GuildID            string                             `json:"guild_id"`
	ChannelID          string                             `json:"channel_id"`
	CreatorID          string                             `json:"creator_id"`
	Name               string                             `json:"name"`
	Description        string                             `json:"description"`
	ScheduledStartTime time.Time                          `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time                         `json:"scheduled_end_time"`
	PrivacyLevel       GuildScheduledEventPrivacyLevel    `json:"privacy_level"`
	Status             GuildScheduledEventStatus          `json:"status"`
	EntityType         GuildScheduledEventEntityType      `json:"entity_type"`
	EntityID           string                             `json:"entity_id"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata"`
	Creator            *User                              `json:"creator"`
	UserCount          int                                `json:"user_count"`
	Image              string                             `json:"image"`
}

// GuildScheduledEventPrivacyLevel represents the privacy level for the scheduled event (discord.GuildScheduledEvent)
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-privacy-level
type GuildScheduledEventPrivacyLevel int

const (
	_ GuildScheduledEventPrivacyLevel = 1 + iota
	GuildScheduledEventPrivacyLevelGuildOnly
)

// GuildScheduledEventStatus represents the current status of the scheduled event (discord.GuildScheduledEvent)
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-status
type GuildScheduledEventStatus int

const (
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = 1 + iota
	GuildScheduledEventStatusActive
	GuildScheduledEventStatusCompleted
	GuildScheduledEventStatusCancelled
)

// GuildScheduledEventEntityType represents the type of the scheduled event (discord.GuildScheduledEvent)
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-entity-types
type GuildScheduledEventEntityType int

const (
	GuildScheduledEventEntityTypeStageInstance = 1 + iota
	GuildScheduledEventEntityTypeVoice
	GuildScheduledEventEntityTypeExternal
)

// GuildScheduledEventEntityMetadata represents additional metadata for the scheduled event (discord.GuildScheduledEvent) of type external (discord.GuildScheduledEventEntityTypeExternal)
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-entity-metadata
type GuildScheduledEventEntityMetadata struct {
	Location string `json:"location,omitempty"`
}

// GuildScheduledEventUser represents a user (discord.User) and member (discord.Member), if any, that is subscribed to a scheduled event (discord.GuildScheduledEvent)
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-user-object-guild-scheduled-event-user-structure
type GuildScheduledEventUser struct {
	GuildScheduledEventID string  `json:"guild_scheduled_event_id"`
	User                  *User   `json:"user"`
	Member                *Member `json:"member,omitempty"`
}
