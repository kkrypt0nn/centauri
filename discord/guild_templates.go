package discord

import "time"

// GuildTemplate represents a template that can be used to creates a guild (discord.Guild) based on a snapshot of an existing guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild-template#guild-template-object-guild-template-structure
type GuildTemplate struct {
	Code                  string        `json:"code"`
	Name                  string        `json:"name"`
	Description           string        `json:"description,omitempty"`
	UsageCount            int           `json:"usage_count"`
	CreatorID             string        `json:"creator_id"`
	Creator               *User         `json:"creator"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
	SourceGuildID         string        `json:"source_guild_id"`
	SerializedSourceGuild *PartialGuild `json:"serialized_source_guild"`
	IsDirty               bool          `json:"is_dirty"`
}

// CreateGuildFromTemplate represents the payload to send to Discord to create a guild (discord.Guild) based on a template code
// https://discord.com/developers/docs/resources/guild-template#create-guild-from-guild-template-json-params
type CreateGuildFromTemplate struct {
	Name  string  `json:"name"`
	Image *string `json:"image,omitempty"`
}

// CreateGuildTemplate represents the payload to send to Discord to create a guild template (discord.GuildTemplate)
// https://discord.com/developers/docs/resources/guild-template#create-guild-template-json-params
type CreateGuildTemplate struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// ModifyGuildTemplate represents the payload to send to Discord to modify an existing a guild template (discord.GuildTemplate)
// https://discord.com/developers/docs/resources/guild-template#modify-guild-template-json-params
type ModifyGuildTemplate struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
