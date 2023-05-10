package discord

type ApplicationRoleConnectionMetadata struct {
	Type                     ApplicationRoleConnectionMetadataType `json:"type"`
	Key                      string                                `json:"key"`
	Name                     string                                `json:"name"`
	NameLocalizations        map[Locale]string                     `json:"name_localizations,omitempty"` // Don't know how that looks like
	Description              string                                `json:"description"`
	DescriptionLocalizations map[Locale]string                     `json:"description_localizations,omitempty"` // Don't know how that looks like
}

type ApplicationRoleConnectionMetadataType int

const (
	ApplicationRoleConnectionMetadataTypeIntegerLessThanOrEqual ApplicationRoleConnectionMetadataType = iota + 1
	ApplicationRoleConnectionMetadataTypeIntegerGreaterThanOrEqual
	ApplicationRoleConnectionMetadataTypeIntegerEqual
	ApplicationRoleConnectionMetadataTypeIntegerNotEqual
	ApplicationRoleConnectionMetadataTypeDatetimeLessThanOrEqual
	ApplicationRoleConnectionMetadataTypeDatetimeGreaterThanOrEqual
	ApplicationRoleConnectionMetadataTypeBooleanEqual
	ApplicationRoleConnectionMetadataTypeBooleanNotEqual
)
