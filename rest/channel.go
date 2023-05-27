package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"net/url"
	"strconv"
	"time"
)

const (
	ChannelsEndpoint = Endpoint + "channels"
)

// GetChannel returns a channel structure (discord.Channel) for the given channel ID
func (c *Client) GetChannel(channelID string) (*discord.Channel, error) {
	return DoRequestAsStructure[discord.Channel](c, "GET", ChannelsEndpoint+"/"+channelID, nil, nil, 1)
}

// ModifyChannel modifies an existing channel (discord.Channel) and returns its structure
func (c *Client) ModifyChannel(channelID string, newChannel discord.ModifyChannel) (*discord.Channel, error) {
	return DoRequestAsStructure[discord.Channel](c, "PATCH", ChannelsEndpoint+"/"+channelID, newChannel, nil, 1, WithReason(newChannel.AuditLogReason))
}

// DeleteChannel deletes an existing channel (discord.Channel)
func (c *Client) DeleteChannel(channelID, reason string) error {
	_, _, err := c.DoRequest("DELETE", ChannelsEndpoint+"/"+channelID, nil, nil, 1, WithReason(reason))
	return err
}

// GetChannelMessages returns a list of message structures (discord.Message) for the given channel ID
func (c *Client) GetChannelMessages(channelID string) ([]discord.Message, error) {
	return DoRequestAsList[discord.Message](c, "GET", ChannelsEndpoint+"/"+channelID+"/messages", nil, nil, 1)
}

// GetChannelMessage returns a message (discord.Message) for the given channel and message IDs
func (c *Client) GetChannelMessage(channelID, messageID string) (*discord.Message, error) {
	return DoRequestAsStructure[discord.Message](c, "GET", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID, nil, nil, 1)
}

// CreateMessage posts a message (discord.Message) in a given channel ID and returns its structure
func (c *Client) CreateMessage(channelID string, message discord.CreateMessage) (*discord.Message, error) {
	return DoRequestAsStructure[discord.Message](c, "POST", ChannelsEndpoint+"/"+channelID+"/messages", message, nil, 1)
}

// CrosspostMessage cross-posts a given message ID from a given announcement channel (discord.ChannelTypeGuildAnnouncement) ID to following channels (discord.Channel) and returns a message structure (discord.Message)
func (c *Client) CrosspostMessage(channelID, messageID string) (*discord.Message, error) {
	return DoRequestAsStructure[discord.Message](c, "POST", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID+"/crosspost", nil, nil, 1)
}

// CreateReaction creates a reaction (discord.Reaction) for the given channel and message IDs
func (c *Client) CreateReaction(channelID, messageID, emoji string) error {
	_, _, err := c.DoRequest("PUT", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID+"/reactions/"+url.QueryEscape(emoji)+"/@me", nil, nil, 1)
	return err
}

// DeleteOwnReaction deletes a reaction (discord.Reaction) the current user has made for the given channel and message IDs
func (c *Client) DeleteOwnReaction(channelID, messageID, emoji string) error {
	_, _, err := c.DoRequest("DELETE", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID+"/reactions/"+url.QueryEscape(emoji)+"/@me", nil, nil, 1)
	return err
}

// DeleteUserReaction deletes a reaction (discord.Reaction) a given user ID has made for the given channel and message IDs
func (c *Client) DeleteUserReaction(channelID, messageID, userID, emoji string) error {
	_, _, err := c.DoRequest("DELETE", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID+"/reactions/"+url.QueryEscape(emoji)+"/"+userID, nil, nil, 1)
	return err
}

// GetReactions returns a list of user structures (discord.User) that reacted to the given message ID in the given channel ID with the given emoji
func (c *Client) GetReactions(channelID, messageID, emoji, after string, limit int) ([]discord.User, error) {
	queryParams := make(QueryParameters)
	if after != "" {
		queryParams["after"] = after
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsList[discord.User](c, "GET", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID+"/reactions/"+emoji, nil, queryParams, 1)
}

// DeleteAllReactions deletes all reactions (discord.Reaction) for the given channel and message IDs
func (c *Client) DeleteAllReactions(channelID, messageID string) error {
	_, _, err := c.DoRequest("DELETE", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID+"/reactions", nil, nil, 1)
	return err
}

// DeleteAllReactionsForEmoji deletes all reactions (discord.Reaction) for the given channel and message IDs and the emoji
func (c *Client) DeleteAllReactionsForEmoji(channelID, messageID, emoji string) error {
	_, _, err := c.DoRequest("DELETE", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID+"/reactions/"+emoji, nil, nil, 1)
	return err
}

// EditMessage edits a message (discord.Message) from the given channel and message IDs and returns its structure
func (c *Client) EditMessage(channelID, messageID string, message discord.EditMessage) (*discord.Message, error) {
	return DoRequestAsStructure[discord.Message](c, "PATCH", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID, message, nil, 1)
}

// DeleteMessage deletes a message from the given channel and message IDs
func (c *Client) DeleteMessage(channelID, messageID, reason string) error {
	_, _, err := c.DoRequest("DELETE", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID, nil, nil, 1, WithReason(reason))
	return err
}

// BulkDeleteMessages deletes multiple messages from the given channel ID in a single request
func (c *Client) BulkDeleteMessages(channelID string, bulkDeleteMessages discord.BulkDeleteMessages) error {
	_, _, err := c.DoRequest("POST", ChannelsEndpoint+"/"+channelID+"/messages/bulk-delete", bulkDeleteMessages, nil, 1, WithReason(bulkDeleteMessages.AuditLogReason))
	return err
}

// EditChannelPermissions edits the permissions for the given channel and overwrite IDs
func (c *Client) EditChannelPermissions(channelID, overwriteID string, channelPermissions discord.EditChannelPermissions) error {
	_, _, err := c.DoRequest("PUT", ChannelsEndpoint+"/"+channelID+"/permissions/"+overwriteID, channelPermissions, nil, 1, WithReason(channelPermissions.AuditLogReason))
	return err
}

// GetChannelInvites returns a list of invite with metadata structures (discord.InviteWithMetadata) for a given channel ID
func (c *Client) GetChannelInvites(channelID string) ([]discord.InviteWithMetadata, error) {
	return DoRequestAsList[discord.InviteWithMetadata](c, "GET", ChannelsEndpoint+"/"+channelID+"/invites", nil, nil, 1)
}

// CreateChannelInvite creates a new invite (discord.Invite) for the given channel ID and returns its new structure
func (c *Client) CreateChannelInvite(channelID string, invite discord.CreateChannelInvite) (*discord.Invite, error) {
	return DoRequestAsStructure[discord.Invite](c, "POST", ChannelsEndpoint+"/"+channelID+"/invites", invite, nil, 1, WithReason(invite.AuditLogReason))
}

// DeleteChannelPermissions delete the permissions for the given channel and overwrite IDs
func (c *Client) DeleteChannelPermissions(channelID, overwriteID, reason string) error {
	_, _, err := c.DoRequest("DELETE", ChannelsEndpoint+"/"+channelID+"/permissions/"+overwriteID, nil, nil, 1, WithReason(reason))
	return err
}

// FollowAnnouncementChannel follows the given channel ID to the given target channel ID and returns a followed channel structure (discord.FollowedChannel)
func (c *Client) FollowAnnouncementChannel(channelID, targetID string) (*discord.FollowedChannel, error) {
	payload := struct {
		WebhookChannelID string `json:"webhook_channel_id"`
	}{
		WebhookChannelID: targetID,
	}
	return DoRequestAsStructure[discord.FollowedChannel](c, "POST", ChannelsEndpoint+"/"+channelID+"/followers", payload, nil, 1)
}

// TriggerTyping triggers the "<User> is typing..." indicator for the given channel ID
func (c *Client) TriggerTyping(channelID string) error {
	_, _, err := c.DoRequest("POST", ChannelsEndpoint+"/"+channelID+"/typing", nil, nil, 1)
	return err
}

// GetPinnedMessages returns a list of message structures (discord.Message) that are pinned in the given channel ID
func (c *Client) GetPinnedMessages(channelID string) ([]discord.Message, error) {
	return DoRequestAsList[discord.Message](c, "GET", ChannelsEndpoint+"/"+channelID+"/pins", nil, nil, 1)
}

// PinMessage pins the given message ID in the given channel ID
func (c *Client) PinMessage(channelID, messageID, reason string) error {
	_, _, err := c.DoRequest("PUT", ChannelsEndpoint+"/"+channelID+"/pins/"+messageID, nil, nil, 1, WithReason(reason))
	return err
}

// UnpinMessage unpins the given message ID in the given channel ID
func (c *Client) UnpinMessage(channelID, messageID, reason string) error {
	_, _, err := c.DoRequest("DELETE", ChannelsEndpoint+"/"+channelID+"/pins/"+messageID, nil, nil, 1, WithReason(reason))
	return err
}

// StartThreadFromMessage creates a new thread from the given message ID and returns a channel structure (discord.Channel)
func (c *Client) StartThreadFromMessage(channelID, messageID string, thread discord.StartThreadFromMessage) (*discord.Channel, error) {
	return DoRequestAsStructure[discord.Channel](c, "POST", ChannelsEndpoint+"/"+channelID+"/messages/"+messageID+"/threads", thread, nil, 1, WithReason(thread.AuditLogReason))
}

// StartThreadWithoutMessage creates a new thread and returns a channel structure (discord.Channel)
func (c *Client) StartThreadWithoutMessage(channelID string, thread discord.StartThreadWithoutMessage) (*discord.Channel, error) {
	return DoRequestAsStructure[discord.Channel](c, "POST", ChannelsEndpoint+"/"+channelID+"/threads", thread, nil, 1, WithReason(thread.AuditLogReason))
}

// StartThreadInForumChannel creates a new thread in the given channel ID and returns a channel structure (discord.Channel) with a nested message structure (discord.Message)
func (c *Client) StartThreadInForumChannel(channelID string, thread discord.StartThreadInForumChannel) (*discord.Channel, error) {
	return DoRequestAsStructure[discord.Channel](c, "POST", ChannelsEndpoint+"/"+channelID+"/threads", thread, nil, 1, WithReason(thread.AuditLogReason))
}

// AddThreadMember adds a member to a thread given the thread and user IDs
func (c *Client) AddThreadMember(threadID, userID string) error {
	_, _, err := c.DoRequest("PUT", ChannelsEndpoint+"/"+threadID+"/thread-members/"+userID, nil, nil, 1)
	return err
}

// LeaveThread removes the current user from a given thread given its ID
func (c *Client) LeaveThread(threadID string) error {
	_, _, err := c.DoRequest("DELETE", ChannelsEndpoint+"/"+threadID+"/thread-members/@me", nil, nil, 1)
	return err
}

// RemoveThreadMember removes another member from a thread given the thread and user IDs
func (c *Client) RemoveThreadMember(threadID, userID string) error {
	_, _, err := c.DoRequest("DELETE", ChannelsEndpoint+"/"+threadID+"/thread-members/"+userID, nil, nil, 1)
	return err
}

// GetThreadMember returns a thread member structure (discord.ThreadMember) for a given thread and user ID if they are a member of said thread
func (c *Client) GetThreadMember(threadID, userID string) (*discord.ThreadMember, error) {
	return DoRequestAsStructure[discord.ThreadMember](c, "GET", ChannelsEndpoint+"/"+threadID+"/thread-members/"+userID, nil, nil, 1)
}

// ListThreadMembers returns a list of thread member structures (discord.ThreadMember) that are members of the given thread ID
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
	return DoRequestAsList[discord.ThreadMember](c, "GET", ChannelsEndpoint+"/"+threadID+"/thread-members", nil, queryParams, 1)
}

// ListPublicArchivedThreads returns an archived threads structure (discord.ArchivedThreads) which contains the public archived threads in the given channel ID
func (c *Client) ListPublicArchivedThreads(channelID string, before *time.Time, limit int) (*discord.ArchivedThreads, error) {
	queryParams := make(QueryParameters)
	if before != nil {
		queryParams["before"] = before.Format(time.RFC3339)
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsStructure[discord.ArchivedThreads](c, "GET", ChannelsEndpoint+"/"+channelID+"/threads/archived/public", nil, queryParams, 1)
}

// ListPrivateArchivedThreads returns an archived threads structure (discord.ArchivedThreads) which contains the private archived threads in the given channel ID
func (c *Client) ListPrivateArchivedThreads(channelID string, before *time.Time, limit int) (*discord.ArchivedThreads, error) {
	queryParams := make(QueryParameters)
	if before != nil {
		queryParams["before"] = before.Format(time.RFC3339)
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsStructure[discord.ArchivedThreads](c, "GET", ChannelsEndpoint+"/"+channelID+"/threads/archived/private", nil, queryParams, 1)
}

// ListJoinedPrivateArchivedThreads returns an archived threads structure (discord.ArchivedThreads) which contains the private archived threads in the given channel ID that the user has joined
func (c *Client) ListJoinedPrivateArchivedThreads(channelID, before string, limit int) (*discord.ArchivedThreads, error) {
	queryParams := make(QueryParameters)
	if before != "" {
		queryParams["before"] = before
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsStructure[discord.ArchivedThreads](c, "GET", ChannelsEndpoint+"/"+channelID+"/users/@me/threads/archived/private", nil, queryParams, 1)
}
