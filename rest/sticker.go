package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	StickersEndpoint     = Endpoint + "stickers"
	StickerPacksEndpoint = Endpoint + "sticker-packs"
)

// GetSticker returns a sticker structure (discord.Sticker) for the given sticker ID
func (c *Client) GetSticker(stickerID string) (*discord.Sticker, error) {
	return DoRequestAsStructure[discord.Sticker](c, "GET", StickersEndpoint+"/"+stickerID, nil, nil, 1)
}

// ListNitroStickerPacks returns the list of sticker pack structures (discord.NitroStickerPacks) available to Nitro subscribers
func (c *Client) ListNitroStickerPacks() (*discord.NitroStickerPacks, error) {
	return DoRequestAsStructure[discord.NitroStickerPacks](c, "GET", StickerPacksEndpoint, nil, nil, 1)
}

// ListGuildStickers returns a list of sticker structures (discord.Sticker) for the given guild ID
func (c *Client) ListGuildStickers(guildID string) ([]discord.Sticker, error) {
	return DoRequestAsList[discord.Sticker](c, "GET", GuildsEndpoint+"/"+guildID+"/stickers", nil, nil, 1)
}

// GetGuildSticker returns a sticker structure (discord.Sticker) for the given guild and sticker IDs
func (c *Client) GetGuildSticker(guildID, stickerID string) (*discord.Sticker, error) {
	return DoRequestAsStructure[discord.Sticker](c, "GET", GuildsEndpoint+"/"+guildID+"/stickers/"+stickerID, nil, nil, 1)
}

// TODO: Support 'Create Guild Sticker' endpoint

// ModifyGuildSticker modifies an existing sticker (discord.Sticker) for the given guild and sticker IDs and returns its new structure
func (c *Client) ModifyGuildSticker(guildID, stickerID string, sticker discord.ModifyGuildSticker) (*discord.Sticker, error) {
	return DoRequestAsStructure[discord.Sticker](c, "PATCH", GuildsEndpoint+"/"+guildID+"/stickers/"+stickerID, sticker, nil, 1, WithReason(sticker.AuditLogReason))
}

// DeleteGuildSticker deletes an existing sticker (discord.Sticker) from the given guild and sticker ID
func (c *Client) DeleteGuildSticker(guildID, stickerID, reason string) error {
	_, _, err := c.DoRequest("DELETE", GuildsEndpoint+"/"+guildID+"/stickers/"+stickerID, nil, nil, 1, WithReason(reason))
	return err
}
