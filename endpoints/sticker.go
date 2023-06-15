package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

const (
	StickersEndpoint     = Endpoint + "stickers"
	StickerPacksEndpoint = Endpoint + "sticker-packs"
)

func Sticker(stickerID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d", StickersEndpoint, stickerID)
}

func StickerPacks() string {
	return StickerPacksEndpoint
}

func GuildStickers(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/stickers", GuildsEndpoint, guildID)
}

func GuildSticker(guildID, stickerID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/stickers/%d", GuildsEndpoint, guildID, stickerID)
}
