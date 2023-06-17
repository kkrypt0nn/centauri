package gateway

import "github.com/kkrypt0nn/centauri/discord"

// EventInterface is a way to create a new interface of an event based on its name
type EventInterface interface {
	New() interface{}
}

// Ready is the event that is sent once the client has completed the handshake and is now ready
// https://discord.com/developers/docs/topics/gateway-events#ready-ready-event-fields
type Ready struct {
	Version          int                        `json:"v"`
	User             *discord.User              `json:"user"`
	Guilds           []discord.UnavailableGuild `json:"guilds"`
	SessionID        string                     `json:"session_id"`
	ResumeGatewayURL string                     `json:"resume_gateway_url"`
	Shard            [2]int                     `json:"shard"`
	Application      discord.Application        `json:"application"`
}

// MessageCreate is the data for the MESSAGE_CREATE event
type MessageCreate struct {
	discord.Message
}

// EventHandler are the various event handlers existing
type EventHandler interface {
	Handle(*Client, interface{})
}

// readyEventHandler is the event handler for the Ready event (gateway.Ready)
type readyEventHandler func(*Client, *Ready)

// Handle handles the Ready event (gateway.Ready)
func (h readyEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*Ready); ok {
		h(c, e)
	}
}

// New returns the Ready event (gateway.Ready) structure
func (h readyEventHandler) New() interface{} {
	return &Ready{}
}

// messageCreateEventHandler is the event handler for the Message Create event (gateway.MessageCreate)
type messageCreateEventHandler func(*Client, *MessageCreate)

// Handle handles the Message Create event (gateway.MessageCreate)
func (h messageCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*MessageCreate); ok {
		h(c, e)
	}
}

// New returns the Message Create event (gateway.MessageCreate) structure
func (h messageCreateEventHandler) New() interface{} {
	return &MessageCreate{}
}

func handlerForFunction(f interface{}) EventHandler {
	switch t := f.(type) {
	case func(*Client, *Ready):
		return readyEventHandler(t)
	case func(*Client, *MessageCreate):
		return messageCreateEventHandler(t)
	}
	return nil
}
