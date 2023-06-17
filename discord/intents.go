package discord

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
