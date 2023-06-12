package discord

import "time"

// VoiceRegion represents an available voice region
// https://discord.com/developers/docs/resources/voice#voice-region-object-voice-region-structure
type VoiceRegion struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Optimal    bool   `json:"optimal"`
	Deprecated bool   `json:"deprecated"`
	Custom     bool   `json:"custom"`
}

// ModifyCurrentUserVoiceState represents the payload to send to Discord to modify the voice state of the current user (discord.User) in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#modify-current-user-voice-state-json-params
type ModifyCurrentUserVoiceState struct {
	ChannelID               *string    `json:"channel_id,omitempty"`
	Suppress                *bool      `json:"suppress,omitempty"`
	RequestToSpeakTimestamp *time.Time `json:"request_to_speak_timestamp,omitempty"`
}

// ModifyUserVoiceState represents the payload to send to Discord to modify the voice state of a user (discord.User) in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#modify-user-voice-state-json-params
type ModifyUserVoiceState struct {
	ChannelID *string `json:"channel_id,omitempty"`
	Suppress  *bool   `json:"suppress,omitempty"`
}
