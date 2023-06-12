package discord

// Sticker represents a sticker that can be sent in messages
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-structure
type Sticker struct {
	ID          string            `json:"id"`
	PackID      string            `json:"pack_id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Tags        string            `json:"tags"`
	Type        StickerType       `json:"type"`
	FormatType  StickerFormatType `json:"format_type"`
	Available   bool              `json:"available"`
	GuildID     string            `json:"guild_id,omitempty"`
	User        *User             `json:"user,omitempty"`
	SortValue   int               `json:"sort_value"`
}

// StickerType represents the type of sticker (discord.Sticker)
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
type StickerType int

const (
	StickerTypeStandard = 1 + iota
	StickerTypeGuild
)

// StickerFormatType represents the type of format of the sticker (discord.Sticker)
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-format-types
type StickerFormatType int

const (
	StickerFormatTypePNG StickerFormatType = 1 + iota
	StickerFormatTypeAPNG
	StickerFormatTypeLottie
	StickerFormatTypeGIF
)

// StickerItem represents a partial sticker (discord.Sticker)
// https://discord.com/developers/docs/resources/sticker#sticker-item-object-sticker-item-structure
type StickerItem struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	FormatType StickerFormatType `json:"format_type"`
}

// StickerPack represents a pack containing multiple standard stickers (discord.StickerTypeStandard)
// https://discord.com/developers/docs/resources/sticker#sticker-pack-object-sticker-pack-structure
type StickerPack struct {
	ID             string    `json:"id"`
	Stickers       []Sticker `json:"stickers"`
	Name           string    `json:"name"`
	SkuID          string    `json:"sku_id"`
	CoverStickerID string    `json:"cover_sticker_id"`
	Description    string    `json:"description"`
	BannerAssetID  string    `json:"banner_asset_id"`
}

// NitroStickerPacks represents a list of default stickers (discord.Sticker) available for Nitro subscribers
// https://discord.com/developers/docs/resources/sticker#list-nitro-sticker-packs-response-structure
type NitroStickerPacks struct {
	StickerPacks []StickerPack `json:"sticker_packs"`
}

// ModifyGuildSticker represents the payload to send to Discord to modify an existing sticker (discord.Sticker)
// https://discord.com/developers/docs/resources/sticker#modify-guild-sticker-json-params
type ModifyGuildSticker struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags        *string `json:"tags,omitempty"`

	AuditLogReason string `json:"-"`
}
