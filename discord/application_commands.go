package discord

import (
	"encoding/json"
	"fmt"
)

// ApplicationCommand represents an application command that is natively implemented within Discord
// https://discord.com/developers/docs/interactions/application-commands#application-command-object
type ApplicationCommand struct {
	ID                       string                    `json:"id"`
	Type                     ApplicationCommandType    `json:"type,omitempty"`
	ApplicationID            string                    `json:"application_id"`
	GuildID                  string                    `json:"guild_id,omitempty"`
	Name                     string                    `json:"name"`
	NameLocalizations        map[Locale]string         `json:"name_localizations,omitempty"`
	Description              string                    `json:"description"`
	DescriptionLocalizations map[Locale]string         `json:"description_localizations,omitempty"`
	Options                  ApplicationCommandOptions `json:"options,omitempty"`
	DefaultMemberPermissions uint64                    `json:"default_member_permissions,string,omitempty"`
	DMPermission             bool                      `json:"dm_permission,omitempty"`
	DefaultPermission        bool                      `json:"default_permission,omitempty"`
	NSFW                     bool                      `json:"nsfw,omitempty"`
	Version                  string                    `json:"version"`
}

// ApplicationCommandType represents the type of command it is
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types
type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = 1 + iota
	ApplicationCommandTypeUser
	ApplicationCommandTypeMessage
)

type ApplicationCommandOption interface {
	Type() ApplicationCommandOptionType
}

type EmptyApplicationCommandOption struct {
	Name string                       `json:"name"`
	Type ApplicationCommandOptionType `json:"type"`

	data ApplicationCommandOption
}

func (o *EmptyApplicationCommandOption) UnmarshalJSON(b []byte) error {
	type emptyOption EmptyApplicationCommandOption

	if err := json.Unmarshal(b, (*emptyOption)(o)); err != nil {
		return fmt.Errorf("failed to unmarshal")
	}

	switch o.Type {
	case ApplicationCommandOptionTypeSubCommand:
		o.data = &SubCommandOption{}
	case ApplicationCommandOptionTypeSubCommandGroup:
		o.data = &SubCommandGroupOption{}
	case ApplicationCommandOptionTypeString:
		o.data = &StringCommandOption{}
	case ApplicationCommandOptionTypeInteger:
		o.data = &IntegerCommandOption{}
	case ApplicationCommandOptionTypeBoolean:
		o.data = &BooleanCommandOption{}
	case ApplicationCommandOptionTypeUser:
		o.data = &UserCommandOption{}
	case ApplicationCommandOptionTypeChannel:
		o.data = &ChannelCommandOption{}
	case ApplicationCommandOptionTypeRole:
		o.data = &RoleCommandOption{}
	case ApplicationCommandOptionTypeMentionable:
		o.data = &MentionableCommandOption{}
	case ApplicationCommandOptionTypeNumber:
		o.data = &NumberCommandOption{}
	case ApplicationCommandOptionTypeAttachment:
		o.data = &AttachmentCommandOption{}
	default:
		return fmt.Errorf("unknown type: %#v", b)
	}

	if err := json.Unmarshal(b, o.data); err != nil {
		return fmt.Errorf("failed to unmarshal type %d", o.Type)
	}

	return nil
}

type ApplicationCommandOptions []ApplicationCommandOption

func (o *ApplicationCommandOptions) UnmarshalJSON(b []byte) error {
	var emptyOptions []EmptyApplicationCommandOption
	if err := json.Unmarshal(b, &emptyOptions); err != nil {
		return err
	}

	if len(emptyOptions) == 0 {
		*o = nil
		return nil
	}

	*o = make([]ApplicationCommandOption, len(emptyOptions))
	for i, option := range emptyOptions {
		(*o)[i] = option.data
	}

	return nil
}

type SubCommandOption struct {
	Name                     string                     `json:"name"`
	NameLocalizations        map[Locale]string          `json:"name_localizations,omitempty"`
	Description              string                     `json:"description"`
	DescriptionLocalizations map[Locale]string          `json:"description_localizations,omitempty"`
	Required                 bool                       `json:"required"`
	Options                  []ApplicationCommandOption `json:"options,omitempty"`
}

func (o *SubCommandOption) Type() ApplicationCommandOptionType {
	return ApplicationCommandOptionTypeSubCommand
}

func (o *SubCommandOption) MarshalJSON() ([]byte, error) {
	type subCommandOption SubCommandOption
	return json.Marshal(struct {
		Type ApplicationCommandOptionType `json:"type"`
		*subCommandOption
	}{
		Type:             ApplicationCommandOptionTypeSubCommand,
		subCommandOption: (*subCommandOption)(o),
	})
}

type SubCommandGroupOption struct {
	Name                     string             `json:"name"`
	NameLocalizations        map[Locale]string  `json:"name_localizations,omitempty"`
	Description              string             `json:"description"`
	DescriptionLocalizations map[Locale]string  `json:"description_localizations,omitempty"`
	Required                 bool               `json:"required,omitempty"`
	SubCommands              []SubCommandOption `json:"options,omitempty"`
}

func (o *SubCommandGroupOption) Type() ApplicationCommandOptionType {
	return ApplicationCommandOptionTypeSubCommandGroup
}

func (o *SubCommandGroupOption) MarshalJSON() ([]byte, error) {
	type subCommandGroupOption SubCommandGroupOption
	return json.Marshal(struct {
		Type ApplicationCommandOptionType `json:"type"`
		*subCommandGroupOption
	}{
		Type:                  ApplicationCommandOptionTypeSubCommandGroup,
		subCommandGroupOption: (*subCommandGroupOption)(o),
	})
}

type StringCommandOption struct {
	Name                     string            `json:"name"`
	NameLocalizations        map[Locale]string `json:"name_localizations,omitempty"`
	Description              string            `json:"description"`
	DescriptionLocalizations map[Locale]string `json:"description_localizations,omitempty"`
	Required                 bool              `json:"required,omitempty"`
	Choices                  []StringChoice    `json:"choices,omitempty"`
	MinLength                *int              `json:"min_length,omitempty"`
	MaxLength                *int              `json:"max_length,omitempty"`
	AutoComplete             bool              `json:"auto_complete,omitempty"`
}

func (o *StringCommandOption) Type() ApplicationCommandOptionType {
	return ApplicationCommandOptionTypeString
}

func (o *StringCommandOption) MarshalJSON() ([]byte, error) {
	type stringCommandOption StringCommandOption
	return json.Marshal(struct {
		Type ApplicationCommandOptionType `json:"type"`
		*stringCommandOption
	}{
		Type:                ApplicationCommandOptionTypeString,
		stringCommandOption: (*stringCommandOption)(o),
	})
}

type StringChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[Locale]string `json:"name_localizations,omitempty"`
	Value             string            `json:"value"`
}

type IntegerCommandOption struct {
	Name                     string            `json:"name"`
	NameLocalizations        map[Locale]string `json:"name_localizations,omitempty"`
	Description              string            `json:"description"`
	DescriptionLocalizations map[Locale]string `json:"description_localizations,omitempty"`
	Required                 bool              `json:"required,omitempty"`
	Choices                  []IntegerChoice   `json:"choices,omitempty"`
	MinValue                 *int              `json:"min_value,omitempty"`
	MaxValue                 *int              `json:"max_value,omitempty"`
	AutoComplete             bool              `json:"auto_complete,omitempty"`
}

func (o *IntegerCommandOption) Type() ApplicationCommandOptionType {
	return ApplicationCommandOptionTypeInteger
}

func (o *IntegerCommandOption) MarshalJSON() ([]byte, error) {
	type integerCommandOption IntegerCommandOption
	return json.Marshal(struct {
		Type ApplicationCommandOptionType `json:"type"`
		*integerCommandOption
	}{
		Type:                 ApplicationCommandOptionTypeInteger,
		integerCommandOption: (*integerCommandOption)(o),
	})
}

type IntegerChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[Locale]string `json:"name_localizations,omitempty"`
	Value             int               `json:"value"`
}

type BooleanCommandOption struct {
	Name                     string            `json:"name"`
	NameLocalizations        map[Locale]string `json:"name_localizations,omitempty"`
	Description              string            `json:"description"`
	DescriptionLocalizations map[Locale]string `json:"description_localizations,omitempty"`
	Required                 bool              `json:"required,omitempty"`
}

func (o *BooleanCommandOption) Type() ApplicationCommandOptionType {
	return ApplicationCommandOptionTypeBoolean
}

func (o *BooleanCommandOption) MarshalJSON() ([]byte, error) {
	type booleanCommandOption BooleanCommandOption
	return json.Marshal(struct {
		Type ApplicationCommandOptionType `json:"type"`
		*booleanCommandOption
	}{
		Type:                 ApplicationCommandOptionTypeBoolean,
		booleanCommandOption: (*booleanCommandOption)(o),
	})
}

type UserCommandOption struct {
	Name                     string            `json:"name"`
	NameLocalizations        map[Locale]string `json:"name_localizations,omitempty"`
	Description              string            `json:"description"`
	DescriptionLocalizations map[Locale]string `json:"description_localizations,omitempty"`
	Required                 bool              `json:"required,omitempty"`
}

func (o *UserCommandOption) Type() ApplicationCommandOptionType {
	return ApplicationCommandOptionTypeUser
}

func (o *UserCommandOption) MarshalJSON() ([]byte, error) {
	type userCommandOption UserCommandOption
	return json.Marshal(struct {
		Type ApplicationCommandOptionType `json:"type"`
		*userCommandOption
	}{
		Type:              ApplicationCommandOptionTypeUser,
		userCommandOption: (*userCommandOption)(o),
	})
}

type ChannelCommandOption struct {
	Name                     string            `json:"name"`
	NameLocalizations        map[Locale]string `json:"name_localizations,omitempty"`
	Description              string            `json:"description"`
	DescriptionLocalizations map[Locale]string `json:"description_localizations,omitempty"`
	Required                 bool              `json:"required,omitempty"`
	ChannelTypes             []ChannelType     `json:"channel_types,omitempty"`
}

func (o *ChannelCommandOption) Type() ApplicationCommandOptionType {
	return ApplicationCommandOptionTypeChannel
}

func (o *ChannelCommandOption) MarshalJSON() ([]byte, error) {
	type channelCommandOption ChannelCommandOption
	return json.Marshal(struct {
		Type ApplicationCommandOptionType `json:"type"`
		*channelCommandOption
	}{
		Type:                 ApplicationCommandOptionTypeChannel,
		channelCommandOption: (*channelCommandOption)(o),
	})
}

type RoleCommandOption struct {
	Name                     string            `json:"name"`
	NameLocalizations        map[Locale]string `json:"name_localizations,omitempty"`
	Description              string            `json:"description"`
	DescriptionLocalizations map[Locale]string `json:"description_localizations,omitempty"`
	Required                 bool              `json:"required,omitempty"`
}

func (o *RoleCommandOption) Type() ApplicationCommandOptionType {
	return ApplicationCommandOptionTypeRole
}

func (o *RoleCommandOption) MarshalJSON() ([]byte, error) {
	type roleCommandOption RoleCommandOption
	return json.Marshal(struct {
		Type ApplicationCommandOptionType `json:"type"`
		*roleCommandOption
	}{
		Type:              ApplicationCommandOptionTypeRole,
		roleCommandOption: (*roleCommandOption)(o),
	})
}

type MentionableCommandOption struct {
	Name                     string            `json:"name"`
	NameLocalizations        map[Locale]string `json:"name_localizations,omitempty"`
	Description              string            `json:"description"`
	DescriptionLocalizations map[Locale]string `json:"description_localizations,omitempty"`
	Required                 bool              `json:"required,omitempty"`
}

func (o *MentionableCommandOption) Type() ApplicationCommandOptionType {
	return ApplicationCommandOptionTypeMentionable
}

func (o *MentionableCommandOption) MarshalJSON() ([]byte, error) {
	type mentionableCommandOption MentionableCommandOption
	return json.Marshal(struct {
		Type ApplicationCommandOptionType `json:"type"`
		*mentionableCommandOption
	}{
		Type:                     ApplicationCommandOptionTypeMentionable,
		mentionableCommandOption: (*mentionableCommandOption)(o),
	})
}

type NumberCommandOption struct {
	Name                     string            `json:"name"`
	NameLocalizations        map[Locale]string `json:"name_localizations,omitempty"`
	Description              string            `json:"description"`
	DescriptionLocalizations map[Locale]string `json:"description_localizations,omitempty"`
	Required                 bool              `json:"required,omitempty"`
	Choices                  []NumberChoice    `json:"choices,omitempty"`
	MinValue                 float64           `json:"min_value,omitempty"`
	MaxValue                 float64           `json:"max_value,omitempty"`
	Autocomplete             bool              `json:"autocomplete,omitempty"`
}

func (o *NumberCommandOption) Type() ApplicationCommandOptionType {
	return ApplicationCommandOptionTypeNumber
}

func (o *NumberCommandOption) MarshalJSON() ([]byte, error) {
	type numberCommandOption NumberCommandOption
	return json.Marshal(struct {
		Type ApplicationCommandOptionType `json:"type"`
		*numberCommandOption
	}{
		Type:                ApplicationCommandOptionTypeNumber,
		numberCommandOption: (*numberCommandOption)(o),
	})
}

type NumberChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[Locale]string `json:"name_localizations,omitempty"`
	Value             float64           `json:"value"`
}

type AttachmentCommandOption struct {
	Name                     string            `json:"name"`
	NameLocalizations        map[Locale]string `json:"name_localizations,omitempty"`
	Description              string            `json:"description"`
	DescriptionLocalizations map[Locale]string `json:"description_localizations,omitempty"`
	Required                 bool              `json:"required,omitempty"`
}

func (o *AttachmentCommandOption) Type() ApplicationCommandOptionType {
	return ApplicationCommandOptionTypeAttachment
}

func (o *AttachmentCommandOption) MarshalJSON() ([]byte, error) {
	type attachmentCommandOption AttachmentCommandOption
	return json.Marshal(struct {
		Type ApplicationCommandOptionType `json:"type"`
		*attachmentCommandOption
	}{
		Type:                    ApplicationCommandOptionTypeAttachment,
		attachmentCommandOption: (*attachmentCommandOption)(o),
	})
}

// ApplicationCommandOptionType represents the type of option it is
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-type
type ApplicationCommandOptionType int

const (
	ApplicationCommandOptionTypeSubCommand ApplicationCommandOptionType = 1 + iota
	ApplicationCommandOptionTypeSubCommandGroup
	ApplicationCommandOptionTypeString
	ApplicationCommandOptionTypeInteger
	ApplicationCommandOptionTypeBoolean
	ApplicationCommandOptionTypeUser
	ApplicationCommandOptionTypeChannel
	ApplicationCommandOptionTypeRole
	ApplicationCommandOptionTypeMentionable
	ApplicationCommandOptionTypeNumber
	ApplicationCommandOptionTypeAttachment
)

type GuildApplicationCommandPermissions struct {
	ID            string                          `json:"id"`
	ApplicationID string                          `json:"application_id"`
	GuildID       string                          `json:"guild_id"`
	Permissions   []ApplicationCommandPermissions `json:"permissions"`
}

type ApplicationCommandPermissions struct {
	ID         string                           `json:"id"`
	Type       ApplicationCommandPermissionType `json:"type"`
	Permission bool                             `json:"permission"`
}

type ApplicationCommandPermissionType int

const (
	ApplicationCommandPermissionTypeRole ApplicationCommandPermissionType = 1 + iota
	ApplicationCommandPermissionTypeUser
	ApplicationCommandPermissionTypeChannel
)
