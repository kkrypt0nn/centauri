package gateway

import (
	"encoding/json"
	"github.com/kkrypt0nn/centauri/gateway/packets"
)

// EventName represents the name of the events
type EventName string

const (
	EventNameReady EventName = "READY"
)

// Event is the structure of a Gateway event
// https://discord.com/developers/docs/topics/gateway#gateway-events-example-gateway-event
type Event struct {
	OpCode   packets.OpCode  `json:"op"`
	Sequence int64           `json:"s"`
	Type     string          `json:"t"`
	Data     json.RawMessage `json:"d"`
}

type EventHandler interface {
	Handle(c *Client, i interface{})
}
