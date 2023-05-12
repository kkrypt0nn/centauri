package rest

import "github.com/kkrypt0nn/centauri/discord"

// ListGuildEmojis returns a list of discord.Emoji available in the given guild
func (c *Client) ListGuildEmojis(guildID string) ([]discord.Emoji, error) {
	return DoRequestAsList[discord.Emoji](c, "GET", GuildsEndpoint+"/"+guildID+"/emojis", nil, 1)
}

// GetGuildEmoji returns a discord.Emoji given its ID
func (c *Client) GetGuildEmoji(guildID, emojiID string) (*discord.Emoji, error) {
	return DoRequestAs[discord.Emoji](c, "GET", GuildsEndpoint+"/"+guildID+"/emojis/"+emojiID, nil, 1)
}
