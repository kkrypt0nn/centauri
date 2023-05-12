package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"strconv"
)

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

// GetReactions returns a list of discord.User that reacted to the message with the given emoji
func (c *Client) GetReactions(channelID, messageID, emoji, after string, limit int) ([]discord.User, error) {
	queryParams := make(QueryParameters)
	if after != "" {
		queryParams["after"] = after
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsList[discord.User](c, "GET", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID+"/reactions/"+emoji, queryParams, 1)
}

// GetChannelInvites returns a list of discord.Invite with discord.InviteMetadata that are available in the specific channel
func (c *Client) GetChannelInvites(channelID string) ([]discord.Invite, error) {
	return DoRequestAsList[discord.Invite](c, "GET", ChannelsEndpoint+"/"+channelID+"/invites", nil, 1)
}

// GetPinnedMessages returns a list of discord.Message that are pinned in the given channel
func (c *Client) GetPinnedMessages(channelID string) ([]discord.Message, error) {
	return DoRequestAsList[discord.Message](c, "GET", ChannelsEndpoint+"/"+channelID+"/pins", nil, 1)
}

// GetThreadMember returns the discord.ThreadMember if they are a member of the thread
func (c *Client) GetThreadMember(threadID, userID string) (*discord.ThreadMember, error) {
	return DoRequestAs[discord.ThreadMember](c, "GET", ChannelsEndpoint+"/"+threadID+"/thread-members/"+userID, nil, 1)
}
