package gateway

import "github.com/kkrypt0nn/centauri/discord"

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

// ReadyEventHandler is the event handler for the Ready event (events.Ready)
type ReadyEventHandler func(client *Client, ready *Ready)

// Handle handles the Ready event (events.Ready)
func (handler ReadyEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*Ready); ok {
		handler(c, e)
	}
}
