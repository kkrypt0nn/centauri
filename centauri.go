package centauri

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/gateway"
	"github.com/kkrypt0nn/centauri/rest"
)

// NewRestClient returns a new REST client (rest.Client) to make REST API calls only, may be a bot or a user token
func NewRestClient(token string) *rest.Client {
	return rest.New(token)
}

// NewGatewayClient returns a new Gateway API client (gateway.Client) to handle real-time events from Discord
func NewGatewayClient(token string, intents discord.Intents) *gateway.Client {
	return gateway.New(token, intents)
}
