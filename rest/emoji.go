package rest

import "github.com/kkrypt0nn/centauri/discord"

// ListGuildEmojis returns a list of emoji structures (discord.Emoji) for the given guild ID
func (c *Client) ListGuildEmojis(guildID string) ([]discord.Emoji, error) {
	return DoRequestAsList[discord.Emoji](c, "GET", GuildsEndpoint+"/"+guildID+"/emojis", nil, nil, 1)
}

// GetGuildEmoji returns an emoji structure (discord.Emoji) for the given guild and emoji IDs
func (c *Client) GetGuildEmoji(guildID, emojiID string) (*discord.Emoji, error) {
	return DoRequestAsStructure[discord.Emoji](c, "GET", GuildsEndpoint+"/"+guildID+"/emojis/"+emojiID, nil, nil, 1)
}

// CreateGuildEmoji creates a new emoji (discord.Emoji) for the given guild ID and returns its structure
func (c *Client) CreateGuildEmoji(guildID string, emoji discord.CreateGuildEmoji) (*discord.Emoji, error) {
	return DoRequestAsStructure[discord.Emoji](c, "POST", GuildsEndpoint+"/"+guildID+"/emojis", emoji, nil, 1, WithReason(emoji.AuditLogReason))
}

// ModifyGuildEmoji modifies an existing emoji (discord.Emoji) for the given guild and emoji IDs and returns the edited emoji structure
func (c *Client) ModifyGuildEmoji(guildID, emojiID string, emoji discord.ModifyGuildEmoji) (*discord.Emoji, error) {
	return DoRequestAsStructure[discord.Emoji](c, "PATCH", GuildsEndpoint+"/"+guildID+"/emojis/"+emojiID, emoji, nil, 1, WithReason(emoji.AuditLogReason))
}

// DeleteGuildEmoji deletes an existing emoji (discord.Emoji) for the given guild and emoji IDs
func (c *Client) DeleteGuildEmoji(guildID, emojiID, reason string) error {
	return DoEmptyRequest(c, "DELETE", GuildsEndpoint+"/"+guildID+"/emojis/"+emojiID, nil, nil, 1, WithReason(reason))
}
