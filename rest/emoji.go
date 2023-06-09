package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
)

// ListGuildEmojis returns a list of emoji structures (discord.Emoji) for the given guild ID
func (c *Client) ListGuildEmojis(guildID discord.Snowflake) ([]discord.Emoji, error) {
	return DoRequestAsList[discord.Emoji](c, "GET", endpoints.GuildEmojis(guildID), nil, nil, 1)
}

// GetGuildEmoji returns an emoji structure (discord.Emoji) for the given guild and emoji IDs
func (c *Client) GetGuildEmoji(guildID, emojiID discord.Snowflake) (*discord.Emoji, error) {
	return DoRequestAsStructure[discord.Emoji](c, "GET", endpoints.GuildEmoji(guildID, emojiID), nil, nil, 1)
}

// CreateGuildEmoji creates a new emoji (discord.Emoji) for the given guild ID and returns its structure
func (c *Client) CreateGuildEmoji(guildID discord.Snowflake, emoji discord.CreateGuildEmoji) (*discord.Emoji, error) {
	return DoRequestAsStructure[discord.Emoji](c, "POST", endpoints.GuildEmojis(guildID), emoji, nil, 1, WithReason(emoji.AuditLogReason))
}

// ModifyGuildEmoji modifies an existing emoji (discord.Emoji) for the given guild and emoji IDs and returns the edited emoji structure
func (c *Client) ModifyGuildEmoji(guildID, emojiID discord.Snowflake, emoji discord.ModifyGuildEmoji) (*discord.Emoji, error) {
	return DoRequestAsStructure[discord.Emoji](c, "PATCH", endpoints.GuildEmoji(guildID, emojiID), emoji, nil, 1, WithReason(emoji.AuditLogReason))
}

// DeleteGuildEmoji deletes an existing emoji (discord.Emoji) for the given guild and emoji IDs
func (c *Client) DeleteGuildEmoji(guildID, emojiID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.GuildEmoji(guildID, emojiID), nil, nil, 1, WithReason(reason))
}
