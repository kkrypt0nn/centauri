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
