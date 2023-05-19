package discord

// InteractionType represents the type of interaction that has been executed
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-type
type InteractionType int

const (
	InteractionTypePing InteractionType = 1 + iota
	InteractionTypeApplicationCommand
	InteractionTypeMessageComponent
	InteractionTypeApplicationCommandAutocomplete
	InteractionTypeModalSubmit
)
