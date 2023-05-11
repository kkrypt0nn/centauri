package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	ChannelsEndpoint = Endpoint + "channels"
)

// GetChannel returns a discord.Channel given its ID
func (c *Client) GetChannel(channelID string) (*discord.Channel, error) {
	return DoRequestAs[discord.Channel](c, "GET", ChannelsEndpoint+"/"+channelID, nil, 1)
}

// GetChannelMessages returns a list of discord.Message from a channel
func (c *Client) GetChannelMessages(channelID string) ([]discord.Message, error) {
	return DoRequestAsList[discord.Message](c, "GET", ChannelsEndpoint+"/"+channelID+"/messages", nil, 1)
}

// GetChannelMessage returns a discord.Message given its ID
func (c *Client) GetChannelMessage(channelID, messageID string) (*discord.Message, error) {
	return DoRequestAs[discord.Message](c, "GET", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID, nil, 1)
}
