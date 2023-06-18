package gateway

import (
	"github.com/kkrypt0nn/centauri/discord"
)

// EventIdentify represents the event to send to the gateway right after opening a connection
// https://discord.com/developers/docs/topics/gateway-events#identify
type EventIdentify struct {
	OpCode OpCode       `json:"op"`
	Data   IdentifyData `json:"d"`
}

// NewIdentifyEvent creates a new identifying event
func NewIdentifyEvent(token string, properties ConnectionProperties, intents discord.Intents) EventIdentify {
	return EventIdentify{
		OpCode: OpCodeIdentify,
		Data: IdentifyData{
			Token:      token,
			Properties: properties,
			Intents:    intents,
		},
	}
}

// IdentifyData represents the data to send when sending an identify event (discord.EventIdentify)
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

// Heartbeat represents a heartbeat event that should be sent do Discord to keep the connection alive
type Heartbeat struct {
	OpCode   OpCode `json:"op"`
	Sequence int64  `json:"d"`
}

// NewHeartbeat creates a new heartbeat event
func NewHeartbeat(sequence int64) Heartbeat {
	return Heartbeat{
		OpCode:   OpCodeHeartbeat,
		Sequence: sequence,
	}
}
