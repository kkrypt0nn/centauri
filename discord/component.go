package discord

import (
	"encoding/json"
	"fmt"
)

type Component interface {
	Type() ComponentType
}

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

type ButtonStyle int

const (
	ButtonStylePrimary ButtonStyle = 1 + iota
	ButtonStyleSecondary
	ButtonStyleSuccess
	ButtonStyleDanger
	ButtonStyleLink
)

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

type SelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Emoji       *Emoji `json:"emoji,omitempty"`
	Default     bool   `json:"default"`
}

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

type TextInputStyle int

const (
	TextInputStyleShort = 1 + iota
	TextInputStyleParagraph
)
