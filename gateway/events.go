package gateway

import (
	"encoding/json"
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/gateway/packets"
	"io"
)

// EventType represents the event type
type EventType string

const (
	EventTypeReady         EventType = "READY"
	EventTypeResumed       EventType = "RESUMED"
	EventTypeMessageCreate EventType = "MESSAGE_CREATE"
)

// Event is the structure of a gateway event
// https://discord.com/developers/docs/topics/gateway#gateway-events-example-gateway-event
type Event struct {
	OpCode   packets.OpCode  `json:"op"`
	Sequence int64           `json:"s"`
	Type     string          `json:"t"`
	Data     json.RawMessage `json:"d"`

	Struct interface{} `json:"-"`
}

// NewEvent parses a websocket message and returns a new event
func NewEvent(message io.Reader) (*Event, error) {
	var event *Event
	err := json.NewDecoder(message).Decode(&event)
	if err != nil {
		return nil, err
	}
	return event, err
}

// EventHello represents an event sent by the gateway to inform we should now identify
// https://discord.com/developers/docs/topics/gateway-events#hello
type EventHello struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// EventIdentify represents the event to send to the gateway right after opening a connection
// https://discord.com/developers/docs/topics/gateway-events#identify
type EventIdentify struct {
	OpCode packets.OpCode `json:"op"`
	Data   IdentifyData   `json:"d"`
}

// NewIdentifyEvent creates a new identifying event
func NewIdentifyEvent(token string, properties ConnectionProperties, intents discord.Intents) EventIdentify {
	return EventIdentify{
		OpCode: packets.OpCodeIdentify,
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
	OpCode   packets.OpCode `json:"op"`
	Sequence int64          `json:"d"`
}

// NewHeartbeat creates a new heartbeat event
func NewHeartbeat(sequence int64) Heartbeat {
	return Heartbeat{
		OpCode:   packets.OpCodeHeartbeat,
		Sequence: sequence,
	}
}
