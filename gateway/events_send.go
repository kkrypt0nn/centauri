package gateway

import (
	"github.com/kkrypt0nn/centauri/discord"
)

// Identify represents the event to send to the gateway right after opening a connection
// https://discord.com/developers/docs/topics/gateway-events#identify
type Identify struct {
	OpCode OpCode       `json:"op"`
	Data   IdentifyData `json:"d"`
}

// NewIdentifyEvent creates a new Identify event
func NewIdentifyEvent(token string, properties ConnectionProperties, intents discord.Intents) Identify {
	return Identify{
		OpCode: OpCodeIdentify,
		Data: IdentifyData{
			Token:      token,
			Properties: properties,
			Intents:    intents,
		},
	}
}

// IdentifyData represents the data to send when sending an Identify event
type IdentifyData struct {
	Token      string               `json:"token"`
	Properties ConnectionProperties `json:"properties"`
	Intents    discord.Intents      `json:"intents"`
}

// ConnectionProperties represents properties about the client connecting
type ConnectionProperties struct {
	OS      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

// Resume represents a resume event that should be sent to the gateway when wanting to resume a session
// https://discord.com/developers/docs/topics/gateway-events#resume
type Resume struct {
	OpCode OpCode     `json:"op"`
	Data   ResumeData `json:"d"`
}

// NewResumeEvent creates a new Resume event
func NewResumeEvent(token, sessionID string, sequence int64) Resume {
	return Resume{
		OpCode: OpCodeResume,
		Data: ResumeData{
			Token:     token,
			SessionID: sessionID,
			Sequence:  sequence,
		},
	}
}

// ResumeData represents the data to send when sending a Resume event
type ResumeData struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Sequence  int64  `json:"seq"`
}

// Heartbeat represents a heartbeat event that should be sent to the gateway to keep the connection alive
// https://discord.com/developers/docs/topics/gateway-events#heartbeat
type Heartbeat struct {
	OpCode   OpCode `json:"op"`
	Sequence int64  `json:"d"`
}

// NewHeartbeatEvent creates a new Heartbeat event
func NewHeartbeatEvent(sequence int64) Heartbeat {
	return Heartbeat{
		OpCode:   OpCodeHeartbeat,
		Sequence: sequence,
	}
}

// EventSend is an interface that is implemented by all events that can be sent to the gateway
type EventSend interface {
	GetOpCode() OpCode
}

// RequestGuildMembers represents the event to send to the gateway to receive all members for a guild or list of guilds
// https://discord.com/developers/docs/topics/gateway-events#request-guild-members
type RequestGuildMembers struct {
	OpCode OpCode                  `json:"op"`
	Data   RequestGuildMembersData `json:"d"`
}

// GetOpCode returns the OpCode of the event
func (e *RequestGuildMembers) GetOpCode() OpCode {
	return OpCodeRequestGuildMembers
}

// NewRequestGuildMembersEvent creates a new Request Guild Members event
func NewRequestGuildMembersEvent(guildID discord.Snowflake, query string, limit int, presences bool, userIDs discord.ArraySnowflakes, nonce string) *RequestGuildMembers {
	return &RequestGuildMembers{
		OpCode: OpCodeRequestGuildMembers,
		Data: RequestGuildMembersData{
			GuildID:   guildID,
			Query:     query,
			Limit:     limit,
			Presences: presences,
			UserIDs:   userIDs,
			Nonce:     nonce,
		},
	}
}

// RequestGuildMembersData represents the data to send when sending a Request Guild Members event
type RequestGuildMembersData struct {
	GuildID   discord.Snowflake       `json:"guild_id"`
	Query     string                  `json:"query"`
	Limit     int                     `json:"limit"`
	Presences bool                    `json:"presences"`
	UserIDs   discord.ArraySnowflakes `json:"user_ids"`
	Nonce     string                  `json:"nonce"`
}

// UpdateVoiceState represents the event to send to the gateway when wanting to update the voice state
// https://discord.com/developers/docs/topics/gateway-events#update-voice-state
type UpdateVoiceState struct {
	OpCode OpCode               `json:"op"`
	Data   UpdateVoiceStateData `json:"d"`
}

// GetOpCode returns the OpCode of the event
func (e *UpdateVoiceState) GetOpCode() OpCode {
	return OpCodeVoiceStateUpdate
}

// NewUpdateVoiceStateEvent creates a new Update Voice State event
func NewUpdateVoiceStateEvent(guildID, channelID discord.Snowflake, selfMute, selfDeaf bool) *UpdateVoiceState {
	return &UpdateVoiceState{
		OpCode: OpCodeVoiceStateUpdate,
		Data: UpdateVoiceStateData{
			GuildID:   guildID,
			ChannelID: channelID,
			SelfMute:  selfMute,
			SelfDeaf:  selfDeaf,
		},
	}
}

// UpdateVoiceStateData represents the data to send when sending an Update Voice State event
type UpdateVoiceStateData struct {
	GuildID   discord.Snowflake `json:"guild_id"`
	ChannelID discord.Snowflake `json:"channel_id,omitempty"`
	SelfMute  bool              `json:"self_mute"`
	SelfDeaf  bool              `json:"self_deaf"`
}

// UpdatePresence represents the event to send to the gateway when wanting to update the presence
// https://discord.com/developers/docs/topics/gateway-events#update-presence
type UpdatePresence struct {
	OpCode OpCode             `json:"op"`
	Data   UpdatePresenceData `json:"d"`
}

// GetOpCode returns the OpCode of the event
func (e *UpdatePresence) GetOpCode() OpCode {
	return OpCodePresenceUpdate
}

// NewUpdatePresenceEvent creates a new Update Presence event
func NewUpdatePresenceEvent(since int, activities []discord.Activity, status discord.StatusType, afk bool) *UpdatePresence {
	return &UpdatePresence{
		OpCode: OpCodePresenceUpdate,
		Data: UpdatePresenceData{
			Since:      since,
			Activities: activities,
			Status:     status,
			AFK:        afk,
		},
	}
}

// UpdatePresenceData represents the data to send when sending an Update Presence event
type UpdatePresenceData struct {
	Since      int                `json:"since"`
	Activities []discord.Activity `json:"activities"`
	Status     discord.StatusType `json:"status"`
	AFK        bool               `json:"afk"`
}
