package discord

import (
	"encoding/json"
	"fmt"
	"time"
)

// Interaction represents a generic Discord interaction
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-structure
type Interaction struct {
	ID             Snowflake       `json:"id"`
	ApplicationID  Snowflake       `json:"application_id"`
	Type           InteractionType `json:"type"`
	Data           InteractionData `json:"data"`
	GuildID        Snowflake       `json:"guild_id"`
	Channel        Channel         `json:"channel"`
	ChannelID      Snowflake       `json:"channel_id"`
	Member         *Member         `json:"member,omitempty"`
	User           *User           `json:"user,omitempty"`
	Token          string          `json:"token"`
	Version        int             `json:"version"`
	Message        *Message        `json:"message,omitempty"`
	AppPermissions Permissions     `json:"app_permissions"`
	Locale         Locale          `json:"locale"`
	GuildLocale    Locale          `json:"guild_locale"`
}

func (i *Interaction) UnmarshalJSON(b []byte) error {
	type interaction Interaction
	var source struct {
		interaction
		UnidentifiedData json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(b, &source); err != nil {
		return fmt.Errorf("failed unmarshalling: %s", err)
	}

	*i = Interaction(source.interaction)

	switch i.Type {
	case InteractionTypeApplicationCommand:
		data := &InteractionApplicationCommandData{}
		err := json.Unmarshal(source.UnidentifiedData, &data)
		if err != nil {
			return fmt.Errorf("failed unmarshalling: %s", err)
		}
		i.Data = data
	case InteractionTypeMessageComponent:
		data := &InteractionMessageComponentData{}
		err := json.Unmarshal(source.UnidentifiedData, &data)
		if err != nil {
			return fmt.Errorf("failed unmarshalling: %s", err)
		}
		i.Data = data
	case InteractionTypeModalSubmit:
		data := &InteractionModalSubmitData{}
		err := json.Unmarshal(source.UnidentifiedData, &data)
		if err != nil {
			return fmt.Errorf("failed unmarshalling: %s", err)
		}
		i.Data = data
	default:
		return fmt.Errorf("unknown type: %d", i.Type)
	}

	return nil
}

// CreatedAt returns the creation time of the interaction (discord.Interaction)
func (i *Interaction) CreatedAt() time.Time {
	return i.ID.CreatedAt()
}

// InteractionData represents an interface that is implemented by all interaction data structures below
type InteractionData interface {
	InteractionType() InteractionType
}

// InteractionApplicationCommandData represents the data of an application command interaction
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-application-command-data-structure
type InteractionApplicationCommandData struct {
	ID       Snowflake                 `json:"id"`
	Name     string                    `json:"name"`
	Type     ApplicationCommandType    `json:"type"`
	Resolved Resolved                  `json:"resolved,omitempty"`
	Options  ApplicationCommandOptions `json:"options,omitempty"`
	GuildID  Snowflake                 `json:"guild_id,omitempty"`
	TargetID Snowflake                 `json:"target_id,omitempty"`
}

// InteractionType returns the type of the interaction
func (d *InteractionApplicationCommandData) InteractionType() InteractionType {
	return InteractionTypeApplicationCommand
}

// InteractionMessageComponentData represents the data of a message component interaction
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-message-component-data-structure
type InteractionMessageComponentData struct {
	CustomID      string         `json:"custom_id"`
	ComponentType ComponentType  `json:"component_type"`
	Values        []SelectOption `json:"values"`
}

// InteractionType returns the type of the interaction
func (d *InteractionMessageComponentData) InteractionType() InteractionType {
	return InteractionTypeMessageComponent
}

// InteractionModalSubmitData represents the data of a modal submit interaction
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-modal-submit-data-structure
type InteractionModalSubmitData struct {
	CustomID   string      `json:"custom_id"`
	Components []Component `json:"components"`
}

func (d *InteractionModalSubmitData) UnmarshalJSON(b []byte) error {
	type data InteractionModalSubmitData
	var source struct {
		data
		UnidentifiedComponents []UnidentifiedComponent `json:"components"`
	}
	if err := json.Unmarshal(b, &source); err != nil {
		return fmt.Errorf("failed unmarshalling: %s", err)
	}

	*d = InteractionModalSubmitData(source.data)

	if len(source.UnidentifiedComponents) == 0 {
		d.Components = nil
		return nil
	}

	d.Components = make([]Component, len(source.UnidentifiedComponents))
	for i, unidentifiedComponents := range source.UnidentifiedComponents {
		d.Components[i] = unidentifiedComponents.data
	}

	return nil
}

// InteractionType returns the type of the interaction
func (d *InteractionModalSubmitData) InteractionType() InteractionType {
	return InteractionTypeModalSubmit
}

// InteractionType represents the type of interaction (discord.Interaction) that has been executed
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-type
type InteractionType int

const (
	InteractionTypePing InteractionType = 1 + iota
	InteractionTypeApplicationCommand
	InteractionTypeMessageComponent
	InteractionTypeApplicationCommandAutocomplete
	InteractionTypeModalSubmit
)

// Resolved represents the mapped data from Snowflakes to their structures
type Resolved struct {
	Users       map[Snowflake]User       `json:"users,omitempty"`
	Members     map[Snowflake]Member     `json:"members,omitempty"`
	Roles       map[Snowflake]Role       `json:"roles,omitempty"`
	Channels    map[Snowflake]Channel    `json:"channels,omitempty"`
	Messages    map[Snowflake]Message    `json:"messages,omitempty"`
	Attachments map[Snowflake]Attachment `json:"attachments,omitempty"`
}

// CreateInteractionResponse represents the payload to send to Discord to respond to an interaction (discord.Interaction)
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object
type CreateInteractionResponse struct {
	Type InteractionResponseType `json:"type"`
	Data InteractionResponse     `json:"data"`
}

// InteractionResponseType represents the interaction response types
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
type InteractionResponseType int

const (
	InteractionResponseTypePong InteractionResponseType = 1 + iota
	_
	_
	InteractionResponseTypeMessage
	InteractionResponseTypeDefer
	InteractionResponseTypeDeferUpdateMessage
	InteractionResponseTypeUpdateMessage
	InteractionResponseTypeAutoCompleteResult
	InteractionResponseTypeModal
)

// InteractionResponse represents an interface that is implemented by all interaction response types
type InteractionResponse interface {
	Type() InteractionResponseType
}

// MessageInteractionResponse represents a response to an interaction (discord.Interaction) with a message
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-messages
type MessageInteractionResponse struct {
	TTS             bool            `json:"tts"`
	Content         *string         `json:"content,omitempty"`
	Embeds          []Embed         `json:"embeds"`
	AllowedMentions AllowedMentions `json:"allowed_mentions"`
	Flags           MessageFlags    `json:"flags"`
	Components      []Component     `json:"components"`
	Attachments     []Attachment    `json:"attachments,omitempty"`

	Files []File `json:"-"`
}

// Type returns the type of interaction response
func (r MessageInteractionResponse) Type() InteractionResponseType {
	return InteractionResponseTypeMessage
}

// AutoCompleteInteractionResponse represents a response to an interaction (discord.Interaction) with auto complete choices
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-autocomplete
type AutoCompleteInteractionResponse[T IntegerChoice | StringChoice | NumberChoice] struct {
	Choices []T `json:"choices"`
}

// Type returns the type of interaction response
func (r AutoCompleteInteractionResponse[T]) Type() InteractionResponseType {
	return InteractionResponseTypeAutoCompleteResult
}

// ModalInteractionResponse represents a response to an interaction (discord.Interaction) with a modal
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-modal
type ModalInteractionResponse struct {
	CustomID   string      `json:"custom_id"`
	Title      string      `json:"title"`
	Components []TextInput `json:"components"`
}

// Type returns the type of interaction response
func (r ModalInteractionResponse) Type() InteractionResponseType {
	return InteractionResponseTypeModal
}

// EditOriginalInteractionResponse represents the payload to send to Discord to edit an original interaction response
// https://discord.com/developers/docs/resources/webhook#edit-webhook-message-jsonform-params
type EditOriginalInteractionResponse struct {
	EditWebhookMessage
}

// CreateFollowupMessage represents the payload to send to Discord to create a followup message
// https://discord.com/developers/docs/resources/webhook#execute-webhook-jsonform-params
type CreateFollowupMessage struct {
	Content         *string          `json:"content,omitempty"`
	TTS             *bool            `json:"tts,omitempty"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`
	Components      []Component      `json:"components,omitempty"`
	Flags           *MessageFlags    `json:"flags,omitempty"`
	ThreadName      *string          `json:"thread_name,omitempty"`
	Attachments     []Attachment     `json:"attachments,omitempty"`

	Files []File `json:"-"`
}

// EditFollowupMessage represents the payload to send to Discord to edit a followup message
// https://discord.com/developers/docs/resources/webhook#edit-webhook-message-jsonform-params
type EditFollowupMessage struct {
	EditWebhookMessage
}
