package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	StickersEndpoint     = Endpoint + "stickers"
	StickerPacksEndpoint = Endpoint + "sticker-packs"
)

// GetSticker returns a sticker structure (discord.Sticker) for the given sticker ID
func (c *Client) GetSticker(stickerID string) (*discord.Sticker, error) {
	return DoRequestAs[discord.Sticker](c, "GET", StickersEndpoint+"/"+stickerID, nil, 1)
}

// ListNitroStickerPacks returns the list of sticker pack structures (discord.NitroStickerPacks) available to Nitro subscribers
func (c *Client) ListNitroStickerPacks() (*discord.NitroStickerPacks, error) {
	return DoRequestAs[discord.NitroStickerPacks](c, "GET", StickerPacksEndpoint, nil, 1)
}

// ListGuildStickers returns a list of sticker structures (discord.Sticker) for the given guild ID
func (c *Client) ListGuildStickers(guildID string) ([]discord.Sticker, error) {
	return DoRequestAsList[discord.Sticker](c, "GET", GuildsEndpoint+"/"+guildID+"/stickers", nil, 1)
}

// GetGuildSticker returns a sticker structure (discord.Sticker) for the given guild and sticker IDs
func (c *Client) GetGuildSticker(guildID, stickerID string) (*discord.Sticker, error) {
	return DoRequestAs[discord.Sticker](c, "GET", GuildsEndpoint+"/"+guildID+"/stickers/"+stickerID, nil, 1)
}
