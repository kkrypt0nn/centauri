package discord

import "github.com/kkrypt0nn/centauri/utils/flags"

// Intents represents the gateway intents
// https://discord.com/developers/docs/topics/gateway#gateway-intents
type Intents uint64

const (
	IntentsGuilds Intents = 1 << iota
	IntentsGuildMembers
	IntentsGuildModeration
	IntentsGuildEmojisAndStickers
	IntentsGuildIntegrations
	IntentsGuildWebhooks
	IntentsGuildInvites
	IntentsGuildVoiceStates
	IntentsGuildPresences
	IntentsGuildMessages
	IntentsGuildMessageReactions
	IntentsGuildMessageTyping
	IntentsDirectMessages
	IntentsDirectMessageReactions
	IntentsDirectMessageTyping
	IntentsMessageContent
	IntentsGuildScheduledEvents
	_
	_
	_
	IntentsAutoModerationConfiguration
	IntentsAutoModerationExecution
)

// ComputeIntents creates a new intents structure (discord.Intents) from the given intents
func ComputeIntents(intents ...Intents) Intents {
	return flags.Compute(intents...)
}

// Add adds the given intents (discord.Intents)
func (f Intents) Add(intents ...Intents) Intents {
	return flags.Add(f, intents...)
}

// Remove removes the given intents (discord.Intents)
func (f Intents) Remove(intents ...Intents) Intents {
	return flags.Remove(f, intents...)
}

// Toggle toggles the given intents (discord.Intents)
func (f Intents) Toggle(intents ...Intents) Intents {
	return flags.Toggle(f, intents...)
}

// Has checks if all the given intents (discord.Intents) are set
func (f Intents) Has(intents ...Intents) bool {
	return flags.Has(f, intents...)
}

// HasAny checks if any of the given intents (discord.Intents) is set
func (f Intents) HasAny(intents ...Intents) bool {
	return flags.HasAny(f, intents...)
}

// HasNot checks if all the given intents (discord.Intents) are not set
func (f Intents) HasNot(intents ...Intents) bool {
	return flags.HasNot(f, intents...)
}

// HasNotAny checks if any of the given intents (discord.Intents) is not set
func (f Intents) HasNotAny(intents ...Intents) bool {
	return flags.HasNotAny(f, intents...)
}
