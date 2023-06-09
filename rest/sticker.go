package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
)

// GetSticker returns a sticker structure (discord.Sticker) for the given sticker ID
func (c *Client) GetSticker(stickerID discord.Snowflake) (*discord.Sticker, error) {
	return DoRequestAsStructure[discord.Sticker](c, "GET", endpoints.Sticker(stickerID), nil, nil, 1)
}

// ListNitroStickerPacks returns the list of sticker pack structures (discord.NitroStickerPacks) available to Nitro subscribers
func (c *Client) ListNitroStickerPacks() (*discord.NitroStickerPacks, error) {
	return DoRequestAsStructure[discord.NitroStickerPacks](c, "GET", endpoints.StickerPacks(), nil, nil, 1)
}

// ListGuildStickers returns a list of sticker structures (discord.Sticker) for the given guild ID
func (c *Client) ListGuildStickers(guildID discord.Snowflake) ([]discord.Sticker, error) {
	return DoRequestAsList[discord.Sticker](c, "GET", endpoints.GuildStickers(guildID), nil, nil, 1)
}

// GetGuildSticker returns a sticker structure (discord.Sticker) for the given guild and sticker IDs
func (c *Client) GetGuildSticker(guildID, stickerID discord.Snowflake) (*discord.Sticker, error) {
	return DoRequestAsStructure[discord.Sticker](c, "GET", endpoints.GuildSticker(guildID, stickerID), nil, nil, 1)
}

// CreateGuildSticker creates a new sticker (discord.Sticker) for the given guild ID and returns its structure
// **Warning**: Seems to not be working at the moment, always giving "Invalid Asset"
func (c *Client) CreateGuildSticker(guildID discord.Snowflake, sticker discord.CreateGuildSticker) (*discord.Sticker, error) {
	contentType, body, err := CreateMultipartBodyWithJSON(sticker, []discord.File{
		sticker.File,
	})
	if err != nil {
		return nil, err
	}
	return DoRequestWithFiles[discord.Sticker](c, "POST", endpoints.GuildStickers(guildID), body, nil, 1, WithReason(sticker.AuditLogReason), WithHeader("Content-Type", contentType))
}

// ModifyGuildSticker modifies an existing sticker (discord.Sticker) for the given guild and sticker IDs and returns its new structure
func (c *Client) ModifyGuildSticker(guildID, stickerID discord.Snowflake, sticker discord.ModifyGuildSticker) (*discord.Sticker, error) {
	return DoRequestAsStructure[discord.Sticker](c, "PATCH", endpoints.GuildSticker(guildID, stickerID), sticker, nil, 1, WithReason(sticker.AuditLogReason))
}

// DeleteGuildSticker deletes an existing sticker (discord.Sticker) from the given guild and sticker ID
func (c *Client) DeleteGuildSticker(guildID, stickerID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.GuildSticker(guildID, stickerID), nil, nil, 1, WithReason(reason))
}
