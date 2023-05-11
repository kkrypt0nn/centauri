package discord

type Sticker struct {
	ID          string            `json:"id"`
	PackID      string            `json:"pack_id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Tags        []string          `json:"tags"`
	Type        StickerType       `json:"type"`
	FormatType  StickerFormatType `json:"format_type"`
	Available   bool              `json:"available"`
	GuildID     string            `json:"guild_id,omitempty"`
	User        *User             `json:"user,omitempty"`
	SortValue   int               `json:"sort_value"`
}

type StickerType int

const (
	StickerTypeStandard = 1 + iota
	StickerTypeGuild
)

type StickerFormatType int

const (
	StickerFormatTypePNG StickerFormatType = 1 + iota
	StickerFormatTypeAPNG
	StickerFormatTypeLottie
	StickerFormatTypeGIF
)

type StickerItem struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	FormatType StickerFormatType `json:"format_type"`
}
