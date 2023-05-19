package discord

import (
	"encoding/json"
	"fmt"
)

// Component represents an interface that is implemented by all message components below
// https://discord.com/developers/docs/interactions/message-components#message-components
type Component interface {
	Type() ComponentType
}

// ComponentType represents the type of component (discord.Component)
// https://discord.com/developers/docs/interactions/message-components#component-object-component-types
type ComponentType int

const (
	ComponentTypeActionRow ComponentType = 1 + iota
	ComponentTypeButton
	ComponentTypeSelectMenu
	ComponentTypeTextInput
	ComponentTypeUserSelect
	ComponentTypeRoleSelect
	ComponentTypeMentionableSelect
	ComponentTypeChannelSelect
)

// UnidentifiedComponent represents a component (discord.Component) that hasn't had its type identified yet
type UnidentifiedComponent struct {
	Type ComponentType `json:"type"`
	data Component
}

func (c *UnidentifiedComponent) UnmarshalJSON(b []byte) error {
	type emptyComponent UnidentifiedComponent

	if err := json.Unmarshal(b, (*emptyComponent)(c)); err != nil {
		return fmt.Errorf("failed to unmarshal: %s", err)
	}

	switch c.Type {
	case ComponentTypeActionRow:
		c.data = &ActionRow{}
	case ComponentTypeButton:
		c.data = &Button{}
	case ComponentTypeSelectMenu, ComponentTypeUserSelect, ComponentTypeRoleSelect, ComponentTypeMentionableSelect, ComponentTypeChannelSelect:
		c.data = &SelectMenu{}
	case ComponentTypeTextInput:
		c.data = &TextInput{}
	default:
		return fmt.Errorf("unknown type: %d", c.Type)
	}

	if err := json.Unmarshal(b, c.data); err != nil {
		return fmt.Errorf("failed to unmarshal type %d (%s)", c.Type, err)
	}

	return nil
}

// ActionRow represents a non-interactive container component for other types of components
// https://discord.com/developers/docs/interactions/message-components#action-rows
type ActionRow struct {
	Components []Component `json:"components"`
}

func (c *ActionRow) MarshalJSON() ([]byte, error) {
	type actionRow ActionRow
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		*actionRow
	}{
		Type:      ComponentTypeActionRow,
		actionRow: (*actionRow)(c),
	})
}

func (c *ActionRow) UnmarshalJSON(b []byte) error {
	var src struct {
		UnidentifiedComponents []UnidentifiedComponent `json:"components"`
	}
	if err := json.Unmarshal(b, &src); err != nil {
		return fmt.Errorf("failed unmarshalling: %s", err)
	}

	c.Components = make([]Component, len(src.UnidentifiedComponents))
	for i, unidentifiedComponent := range src.UnidentifiedComponents {
		c.Components[i] = unidentifiedComponent.data
	}
	return nil
}

func (c *ActionRow) Type() ComponentType {
	return ComponentTypeActionRow
}

// Button represents an interactive components that render in messages which can be clicked by users
// https://discord.com/developers/docs/interactions/message-components#button-object-button-structure
type Button struct {
	Style    ButtonStyle `json:"style"`
	Label    string      `json:"label"`
	Emoji    *Emoji      `json:"emoji,omitempty"`
	CustomID string      `json:"custom_id,omitempty"`
	URL      string      `json:"url,omitempty"`
	Disabled bool        `json:"disabled"`
}

func (c *Button) MarshalJSON() ([]byte, error) {
	type button Button
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		*button
	}{
		Type:   ComponentTypeButton,
		button: (*button)(c),
	})
}

func (c *Button) Type() ComponentType {
	return ComponentTypeButton
}

// ButtonStyle represents the style of the button (discord.Button)
// https://discord.com/developers/docs/interactions/message-components#button-object-button-styles
type ButtonStyle int

const (
	ButtonStylePrimary ButtonStyle = 1 + iota
	ButtonStyleSecondary
	ButtonStyleSuccess
	ButtonStyleDanger
	ButtonStyleLink
)

// SelectMenu represents an interactive components that allow users to select one or more options from a dropdown list in messages
// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-menu-structure
type SelectMenu struct {
	MenuType     SelectMenuType `json:"type"`
	CustomID     string         `json:"custom_id,omitempty"`
	Options      []SelectOption `json:"options,omitempty"`
	ChannelTypes []ChannelType  `json:"channel_types,omitempty"`
	Placeholder  string         `json:"placeholder,omitempty"`
	MinValues    int            `json:"min_values,omitempty"`
	MaxValues    int            `json:"max_values,omitempty"`
	Disabled     bool           `json:"disabled"`
}

func (c *SelectMenu) Type() ComponentType {
	if c.MenuType != 0 {
		return ComponentType(c.MenuType)
	}
	return ComponentTypeSelectMenu
}

// SelectMenuType represents the type of select menu (discord.SelectMenu)
// https://discord.com/developers/docs/interactions/message-components#component-object-component-types
type SelectMenuType int

const (
	_ SelectMenuType = 1 + iota
	_
	SelectMenuTypeText
	_
	SelectMenuTypeUser
	SelectMenuTypeRole
	SelectMenuTypeMentionable
	SelectMenuTypeChannels
)

// SelectOption represents the choices a user can choose from
// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-option-structure
type SelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Emoji       *Emoji `json:"emoji,omitempty"`
	Default     bool   `json:"default"`
}

// TextInput represents an interactive component that render on modals
// https://discord.com/developers/docs/interactions/message-components#text-inputs-text-input-structure
type TextInput struct {
	CustomID    string         `json:"custom_id,omitempty"`
	Style       TextInputStyle `json:"style"`
	Label       string         `json:"label"`
	MinLength   int            `json:"min_length,omitempty"`
	MaxLength   int            `json:"max_length,omitempty"`
	Required    bool           `json:"required"`
	Value       string         `json:"value,omitempty"`
	Placeholder string         `json:"placeholder"`
}

func (c *TextInput) MarshalJSON() ([]byte, error) {
	type textInput TextInput
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		*textInput
	}{
		Type:      ComponentTypeTextInput,
		textInput: (*textInput)(c),
	})
}

func (c *TextInput) Type() ComponentType {
	return ComponentTypeTextInput
}

// TextInputStyle represents the style to use for the text input (discord.TextInput)
// https://discord.com/developers/docs/interactions/message-components#text-inputs-text-input-styles
type TextInputStyle int

const (
	TextInputStyleShort = 1 + iota
	TextInputStyleParagraph
)
