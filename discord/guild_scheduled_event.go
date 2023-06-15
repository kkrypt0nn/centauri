package discord

import (
	"fmt"
	"time"
)

// GuildScheduledEvent represents a scheduled event in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-structure
type GuildScheduledEvent struct {
	ID                 Snowflake                          `json:"id"`
	GuildID            Snowflake                          `json:"guild_id"`
	ChannelID          Snowflake                          `json:"channel_id"`
	CreatorID          Snowflake                          `json:"creator_id"`
	Name               string                             `json:"name"`
	Description        string                             `json:"description"`
	ScheduledStartTime time.Time                          `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time                         `json:"scheduled_end_time"`
	PrivacyLevel       GuildScheduledEventPrivacyLevel    `json:"privacy_level"`
	Status             GuildScheduledEventStatus          `json:"status"`
	EntityType         GuildScheduledEventEntityType      `json:"entity_type"`
	EntityID           Snowflake                          `json:"entity_id"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata"`
	Creator            *User                              `json:"creator"`
	UserCount          int                                `json:"user_count"`
	Image              string                             `json:"image"`
}

// CreatedAt returns the creation time of the guild scheduled event (discord.GuildScheduledEvent)
func (e *GuildScheduledEvent) CreatedAt() time.Time {
	return e.ID.CreatedAt()
}

// CoverURL returns the cover URL of the guild scheduled event (discord.GuildScheduledEvent)
func (e *GuildScheduledEvent) CoverURL(asFormat ImageFormat) string {
	if e.Image != "" {
		if asFormat == "" {
			asFormat = "png"
		}

		suffix := fmt.Sprintf("%s.%s", e.Image, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/guild-events/%d/%s", e.ID, suffix)
	}
	return ""
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
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	User                  *User     `json:"user"`
	Member                *Member   `json:"member,omitempty"`
}

// CreateGuildScheduledEvent represents the payload to send to Discord to create a new scheduled event (discord.GuildScheduledEvent)
// https://discord.com/developers/docs/resources/guild-scheduled-event#create-guild-scheduled-event-json-params
type CreateGuildScheduledEvent struct {
	Name               string                             `json:"name"`
	EntityType         GuildScheduledEventEntityType      `json:"entity_type"`
	ChannelID          *Snowflake                         `json:"channel_id,omitempty"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata,omitempty"`
	PrivacyLevel       GuildScheduledEventPrivacyLevel    `json:"privacy_level"`
	ScheduledStartTime time.Time                          `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time                         `json:"scheduled_end_time,omitempty"`
	Description        *string                            `json:"description,omitempty"`
	Image              *string                            `json:"image,omitempty"`

	AuditLogReason string `json:"-"`
}

// ModifyGuildScheduledEvent represents the payload to send to Discord to modify an existing scheduled event (discord.GuildScheduledEvent)
// https://discord.com/developers/docs/resources/guild-scheduled-event#modify-guild-scheduled-event-json-params
type ModifyGuildScheduledEvent struct {
	Name               *string                            `json:"name,omitempty"`
	EntityType         *GuildScheduledEventEntityType     `json:"entity_type,omitempty"`
	ChannelID          *Snowflake                         `json:"channel_id,omitempty"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata,omitempty"`
	PrivacyLevel       *GuildScheduledEventPrivacyLevel   `json:"privacy_level,omitempty"`
	ScheduledStartTime *time.Time                         `json:"scheduled_start_time,omitempty"`
	ScheduledEndTime   *time.Time                         `json:"scheduled_end_time,omitempty"`
	Description        *string                            `json:"description,omitempty"`
	Image              *string                            `json:"image,omitempty"`
	Status             *GuildScheduledEventStatus         `json:"status,omitempty"`

	AuditLogReason string `json:"-"`
}
