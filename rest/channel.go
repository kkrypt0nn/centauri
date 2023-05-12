package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"strconv"
	"time"
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

// GetChannelInvites returns a list of discord.InviteWithMetadata with discord.InviteMetadata that are available in the specific channel
func (c *Client) GetChannelInvites(channelID string) ([]discord.InviteWithMetadata, error) {
	return DoRequestAsList[discord.InviteWithMetadata](c, "GET", ChannelsEndpoint+"/"+channelID+"/invites", nil, 1)
}

// GetPinnedMessages returns a list of discord.Message that are pinned in the given channel
func (c *Client) GetPinnedMessages(channelID string) ([]discord.Message, error) {
	return DoRequestAsList[discord.Message](c, "GET", ChannelsEndpoint+"/"+channelID+"/pins", nil, 1)
}

// GetThreadMember returns the discord.ThreadMember if they are a member of the thread
func (c *Client) GetThreadMember(threadID, userID string) (*discord.ThreadMember, error) {
	return DoRequestAs[discord.ThreadMember](c, "GET", ChannelsEndpoint+"/"+threadID+"/thread-members/"+userID, nil, 1)
}

// ListThreadMembers returns a list of discord.ThreadMember that are members of the thread
func (c *Client) ListThreadMembers(threadID, after string, withMember bool, limit int) ([]discord.ThreadMember, error) {
	queryParams := make(QueryParameters)
	if after != "" {
		queryParams["after"] = after
	}
	if withMember {
		queryParams["with_member"] = "true"
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsList[discord.ThreadMember](c, "GET", ChannelsEndpoint+"/"+threadID+"/thread-members", queryParams, 1)
}

// ListPublicArchivedThreads returns discord.ArchivedThreads which represent public archived threads in the given channel
func (c *Client) ListPublicArchivedThreads(channelID string, before *time.Time, limit int) (*discord.ArchivedThreads, error) {
	queryParams := make(QueryParameters)
	if before != nil {
		queryParams["before"] = before.Format(time.RFC3339)
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAs[discord.ArchivedThreads](c, "GET", ChannelsEndpoint+"/"+channelID+"/threads/archived/public", queryParams, 1)
}

// ListPrivateArchivedThreads returns discord.ArchivedThreads which represent private archived threads in the given channel
func (c *Client) ListPrivateArchivedThreads(channelID string, before *time.Time, limit int) (*discord.ArchivedThreads, error) {
	queryParams := make(QueryParameters)
	if before != nil {
		queryParams["before"] = before.Format(time.RFC3339)
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAs[discord.ArchivedThreads](c, "GET", ChannelsEndpoint+"/"+channelID+"/threads/archived/private", queryParams, 1)
}

// ListJoinedPrivateArchivedThreads returns discord.ArchivedThreads which represent private archived threads in the given channel that the user has joined
func (c *Client) ListJoinedPrivateArchivedThreads(channelID, before string, limit int) (*discord.ArchivedThreads, error) {
	queryParams := make(QueryParameters)
	if before != "" {
		queryParams["before"] = before
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAs[discord.ArchivedThreads](c, "GET", ChannelsEndpoint+"/"+channelID+"/users/@me/threads/archived/private", queryParams, 1)
}
