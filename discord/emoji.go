package discord

import (
	"fmt"
)

// Emoji represents a custom, or not, emoji on Discord
// https://discord.com/developers/docs/resources/emoji#emoji-object-emoji-structure
type Emoji struct {
	ID            string   `json:"id"`
	Name          string   `json:"name,omitempty"`
	Roles         []string `json:"roles,omitempty"`
	User          *User    `json:"user,omitempty"`
	RequireColons bool     `json:"require_colons,omitempty"`
	Managed       bool     `json:"managed,omitempty"`
	Animated      bool     `json:"animated,omitempty"`
	Available     bool     `json:"available,omitempty"`
}

// URL returns the URL for the emoji (discord.Emoji)
func (e *Emoji) URL(asFormat ImageFormat) string {
	if e.ID != "" {
		if asFormat == "" {
			asFormat = "png"
		}

		suffix := fmt.Sprintf("%s.%s", e.ID, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/emojis/%s", suffix)
	}
	return ""
}

// CreateGuildEmoji represents the payload to send to Discord to create a new emoji (discord.Emoji) in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/emoji#create-guild-emoji-json-params
type CreateGuildEmoji struct {
	Name  string   `json:"name"`
	Image string   `json:"image"`
	Roles []string `json:"roles,omitempty"`

	AuditLogReason string `json:"-"`
}

// ModifyGuildEmoji represents the payload to send to Discord to modify an existing emoji (discord.Emoji) in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/emoji#modify-guild-emoji-json-params
type ModifyGuildEmoji struct {
	Name  *string  `json:"name,omitempty"`
	Roles []string `json:"roles,omitempty"`

	AuditLogReason string `json:"-"`
}
