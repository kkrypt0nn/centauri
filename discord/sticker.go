package discord

import (
	"fmt"
	"time"
)

// Sticker represents a sticker that can be sent in messages
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-structure
type Sticker struct {
	ID          Snowflake         `json:"id"`
	PackID      Snowflake         `json:"pack_id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Tags        string            `json:"tags"`
	Type        StickerType       `json:"type"`
	FormatType  StickerFormatType `json:"format_type"`
	Available   bool              `json:"available"`
	GuildID     Snowflake         `json:"guild_id,omitempty"`
	User        *User             `json:"user,omitempty"`
	SortValue   int               `json:"sort_value"`
}

// CreatedAt returns the creation date of the sticker (discord.Sticker)
func (s *Sticker) CreatedAt() time.Time {
	return s.ID.CreatedAt()
}

// URL returns the URL for the sticker (discord.Sticker)
func (s *Sticker) URL(asFormat ImageFormat) string {
	if s.ID != 0 {
		if asFormat == "" {
			asFormat = "png"
		}

		suffix := fmt.Sprintf("%d.%s", s.ID, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/stickers/%s", suffix)
	}
	return ""
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
	ID         Snowflake         `json:"id"`
	Name       string            `json:"name"`
	FormatType StickerFormatType `json:"format_type"`
}

// CreatedAt returns the creation date of the sticker item (discord.StickerItem)
func (i *StickerItem) CreatedAt() time.Time {
	return i.ID.CreatedAt()
}

// StickerPack represents a pack containing multiple standard stickers (discord.StickerTypeStandard)
// https://discord.com/developers/docs/resources/sticker#sticker-pack-object-sticker-pack-structure
type StickerPack struct {
	ID             Snowflake `json:"id"`
	Stickers       []Sticker `json:"stickers"`
	Name           string    `json:"name"`
	SkuID          Snowflake `json:"sku_id"`
	CoverStickerID Snowflake `json:"cover_sticker_id"`
	Description    string    `json:"description"`
	BannerAssetID  Snowflake `json:"banner_asset_id"`
}

// CreatedAt returns the creation date of the sticker pack (discord.StickerPack)
func (p *StickerPack) CreatedAt() time.Time {
	return p.ID.CreatedAt()
}

// BannerURL returns the banner URL of the sticker pack (discord.StickerPack)
func (p *StickerPack) BannerURL() string {
	if p.BannerAssetID != 0 {
		return fmt.Sprintf("https://cdn.discordapp.com/app-icons/710982414301790216/store/%d.png", p.BannerAssetID)
	}
	return ""
}

// NitroStickerPacks represents a list of default stickers (discord.Sticker) available for Nitro subscribers
// https://discord.com/developers/docs/resources/sticker#list-nitro-sticker-packs-response-structure
type NitroStickerPacks struct {
	StickerPacks []StickerPack `json:"sticker_packs"`
}

// CreateGuildSticker represents the payload to send to Discord to create a new guild sticker (discord.Sticker)
// https://discord.com/developers/docs/resources/sticker#create-guild-sticker-form-params
type CreateGuildSticker struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tags        string `json:"tags"`

	File           File   `json:"-"`
	AuditLogReason string `json:"-"`
}

// ModifyGuildSticker represents the payload to send to Discord to modify an existing sticker (discord.Sticker)
// https://discord.com/developers/docs/resources/sticker#modify-guild-sticker-json-params
type ModifyGuildSticker struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags        *string `json:"tags,omitempty"`

	AuditLogReason string `json:"-"`
}
