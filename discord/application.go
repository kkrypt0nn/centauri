package discord

type ApplicationRoleConnectionMetadata struct {
	Type                     ApplicationRoleConnectionMetadataType `json:"type"`
	Key                      string                                `json:"key"`
	Name                     string                                `json:"name"`
	NameLocalizations        interface{}                           `json:"name_localizations,omitempty"` // Don't know how that looks like
	Description              string                                `json:"description"`
	DescriptionLocalizations interface{}                           `json:"description_localizations,omitempty"` // Don't know how that looks like
}

type ApplicationRoleConnectionMetadataType int

const (
	IntegerLessThanOrEqual ApplicationRoleConnectionMetadataType = iota + 1
	IntegerGreaterThanOrEqual
	IntegerEqual
	IntegerNotEqual
	DatetimeLessThanOrEqual
	DatetimeGreaterThanOrEqual
	BooleanEqual
	BooleanNotEqual
)
