package discord

import (
	"encoding/json"
	"fmt"
	"time"
)

// ApplicationCommand represents an application command that is natively implemented within Discord
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-structure
type ApplicationCommand struct {
	ID                       Snowflake                 `json:"id"`
	Type                     ApplicationCommandType    `json:"type,omitempty"`
	ApplicationID            Snowflake                 `json:"application_id"`
	GuildID                  Snowflake                 `json:"guild_id,omitempty"`
	Name                     string                    `json:"name"`
	NameLocalizations        map[Locale]string         `json:"name_localizations,omitempty"`
	Description              string                    `json:"description"`
	DescriptionLocalizations map[Locale]string         `json:"description_localizations,omitempty"`
	Options                  ApplicationCommandOptions `json:"options,omitempty"`
	DefaultMemberPermissions *Permissions              `json:"default_member_permissions,omitempty"`
	DMPermission             bool                      `json:"dm_permission,omitempty"`
	NSFW                     bool                      `json:"nsfw,omitempty"`
	Version                  Snowflake                 `json:"version"`
}

// CreatedAt returns the creation time of the application command (discord.ApplicationCommand)
func (c *ApplicationCommand) CreatedAt() time.Time {
	return c.ID.CreatedAt()
}

// ApplicationCommandType represents the type of command it is
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types
type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = 1 + iota
	ApplicationCommandTypeUser
	ApplicationCommandTypeMessage
)

// ApplicationCommandOption is an interface that will be implemented by all application command options
type ApplicationCommandOption interface {
	Type() ApplicationCommandOptionType
}

// UnidentifiedApplicationCommandOption represents an application command option (discord.ApplicationCommandOption) that hasn't had its type identified yet
type UnidentifiedApplicationCommandOption struct {
	Type ApplicationCommandOptionType `json:"type"`

	data ApplicationCommandOption
}

func (o *UnidentifiedApplicationCommandOption) UnmarshalJSON(b []byte) error {
	type unidentifiedOption UnidentifiedApplicationCommandOption

	if err := json.Unmarshal(b, (*unidentifiedOption)(o)); err != nil {
		return fmt.Errorf("failed to unmarshal: %s", err)
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
		return fmt.Errorf("unknown type: %d", o.Type)
	}

	if err := json.Unmarshal(b, o.data); err != nil {
		return fmt.Errorf("failed to unmarshal type %d", o.Type)
	}

	return nil
}

// ApplicationCommandOptions is a list of application command options (discord.ApplicationCommandOption)
type ApplicationCommandOptions []ApplicationCommandOption

func (o *ApplicationCommandOptions) UnmarshalJSON(b []byte) error {
	var unidentifiedOptions []UnidentifiedApplicationCommandOption
	if err := json.Unmarshal(b, &unidentifiedOptions); err != nil {
		return err
	}

	if len(unidentifiedOptions) == 0 {
		*o = nil
		return nil
	}

	*o = make([]ApplicationCommandOption, len(unidentifiedOptions))
	for i, option := range unidentifiedOptions {
		(*o)[i] = option.data
	}

	return nil
}

// SubCommandOption represents a sub-command option to an application command (discord.ApplicationCommand)
// https://discord.com/developers/docs/interactions/application-commands#subcommands-and-subcommand-groups
type SubCommandOption struct {
	Name                     string                    `json:"name"`
	NameLocalizations        map[Locale]string         `json:"name_localizations,omitempty"`
	Description              string                    `json:"description"`
	DescriptionLocalizations map[Locale]string         `json:"description_localizations,omitempty"`
	Required                 bool                      `json:"required"`
	Options                  ApplicationCommandOptions `json:"options,omitempty"`
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

// SubCommandGroupOption represents a sub-command group option to an application command (discord.ApplicationCommand)
// https://discord.com/developers/docs/interactions/application-commands#subcommands-and-subcommand-groups
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

// StringCommandOption represents an application command option (discord.ApplicationCommandOption) to pass a string
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

// StringChoice is a choice which can be given in a string command option (discord.StringCommandOption)
type StringChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[Locale]string `json:"name_localizations,omitempty"`
	Value             string            `json:"value"`
}

// IntegerCommandOption represents an application command option (discord.ApplicationCommandOption) to pass an integer
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

// IntegerChoice is a choice which can be given in an integer command option (discord.IntegerCommandOption)
type IntegerChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[Locale]string `json:"name_localizations,omitempty"`
	Value             int               `json:"value"`
}

// BooleanCommandOption represents an application command option (discord.ApplicationCommandOption) to pass an integer
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

// UserCommandOption represents an application command option (discord.ApplicationCommandOption) to pass a user (discord.User)
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

// ChannelCommandOption represents an application command option (discord.ApplicationCommandOption) to pass a channel (discord.Channel)
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

// RoleCommandOption represents an application command option (discord.ApplicationCommandOption) to pass a role (discord.Role)
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

// MentionableCommandOption represents an application command option (discord.ApplicationCommandOption) to pass a mentionable (discord.User & discord.Role)
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

// NumberCommandOption represents an application command option (discord.ApplicationCommandOption) to pass a number
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

// NumberChoice is a choice which can be given in a number command option (discord.NumberCommandOption)
type NumberChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[Locale]string `json:"name_localizations,omitempty"`
	Value             float64           `json:"value"`
}

// AttachmentCommandOption represents an application command option (discord.ApplicationCommandOption) to pass an attachment (discord.Attachment)
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

// GuildApplicationCommandPermissions represents the permissions for an app's command in a guild
// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-guild-application-command-permissions-structure
type GuildApplicationCommandPermissions struct {
	ID            Snowflake                      `json:"id"`
	ApplicationID Snowflake                      `json:"application_id"`
	GuildID       Snowflake                      `json:"guild_id"`
	Permissions   []ApplicationCommandPermission `json:"permissions"`
}

// ApplicationCommandPermission represents the permissions for an application command (discord.ApplicationCommand)
// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permissions-structure
type ApplicationCommandPermission struct {
	ID         Snowflake                        `json:"id"`
	Type       ApplicationCommandPermissionType `json:"type"`
	Permission bool                             `json:"permission"`
}

// CreatedAt returns the creation time of the application command permission (discord.ApplicationCommandPermission)
func (p *ApplicationCommandPermission) CreatedAt() time.Time {
	return p.ID.CreatedAt()
}

// ApplicationCommandPermissionType represents the type of application command permission (discord.ApplicationCommandPermission) it is
// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permission-type
type ApplicationCommandPermissionType int

const (
	ApplicationCommandPermissionTypeRole ApplicationCommandPermissionType = 1 + iota
	ApplicationCommandPermissionTypeUser
	ApplicationCommandPermissionTypeChannel
)

// CreateGlobalApplicationCommand represents the payload to send to Discord to create a new global application command (discord.ApplicationCommand)
// https://discord.com/developers/docs/interactions/application-commands#create-global-application-command-json-params
type CreateGlobalApplicationCommand struct {
	Name                     string                    `json:"name"`
	NameLocalizations        map[Locale]string         `json:"name_localizations,omitempty"`
	Description              *string                   `json:"description,omitempty"`
	DescriptionLocalizations map[Locale]string         `json:"description_localizations,omitempty"`
	Options                  ApplicationCommandOptions `json:"options,omitempty"`
	DefaultMemberPermissions *Permissions              `json:"default_member_permissions,omitempty"`
	DMPermission             *bool                     `json:"dm_permission,omitempty"`
	Type                     *ApplicationCommandType   `json:"type,omitempty"`
	NSFW                     *bool                     `json:"nsfw,omitempty"`
}

// CreateGuildApplicationCommand represents the payload to send to Discord to create a new application command (discord.ApplicationCommand)
// https://discord.com/developers/docs/interactions/application-commands#create-guild-application-command-json-params
type CreateGuildApplicationCommand struct {
	Name                     string                    `json:"name"`
	NameLocalizations        map[Locale]string         `json:"name_localizations,omitempty"`
	Description              *string                   `json:"description,omitempty"`
	DescriptionLocalizations map[Locale]string         `json:"description_localizations,omitempty"`
	Options                  ApplicationCommandOptions `json:"options,omitempty"`
	DefaultMemberPermissions *Permissions              `json:"default_member_permissions,omitempty"`
	Type                     *ApplicationCommandType   `json:"type,omitempty"`
	NSFW                     *bool                     `json:"nsfw,omitempty"`
}

// EditGlobalApplicationCommand represents the payload to send to Discord to edit and existing application command (discord.ApplicationCommand)
// https://discord.com/developers/docs/interactions/application-commands#edit-global-application-command-json-params
type EditGlobalApplicationCommand struct {
	Name                     *string                   `json:"name,omitempty"`
	NameLocalizations        map[Locale]string         `json:"name_localizations,omitempty"`
	Description              *string                   `json:"description,omitempty"`
	DescriptionLocalizations map[Locale]string         `json:"description_localizations,omitempty"`
	Options                  ApplicationCommandOptions `json:"options,omitempty"`
	DefaultMemberPermissions *Permissions              `json:"default_member_permissions,omitempty"`
	DMPermission             *bool                     `json:"dm_permission,omitempty"`
	NSFW                     *bool                     `json:"nsfw,omitempty"`
}

// EditGuildApplicationCommand represents the payload to send to Discord to edit and existing application command (discord.ApplicationCommand)
// https://discord.com/developers/docs/interactions/application-commands#edit-guild-application-command-json-params
type EditGuildApplicationCommand struct {
	Name                     *string                   `json:"name,omitempty"`
	NameLocalizations        map[Locale]string         `json:"name_localizations,omitempty"`
	Description              *string                   `json:"description,omitempty"`
	DescriptionLocalizations map[Locale]string         `json:"description_localizations,omitempty"`
	Options                  ApplicationCommandOptions `json:"options,omitempty"`
	DefaultMemberPermissions *Permissions              `json:"default_member_permissions,omitempty"`
	NSFW                     *bool                     `json:"nsfw,omitempty"`
}
