package rest

import (
	"github.com/kkrypt0nn/centauri/endpoints"
	"strconv"
	"time"

	"github.com/kkrypt0nn/centauri/discord"
)

// GetChannel returns a channel structure (discord.Channel) for the given channel ID
func (c *Client) GetChannel(channelID discord.Snowflake) (*discord.Channel, error) {
	return DoRequestAsStructure[discord.Channel](c, "GET", endpoints.Channel(channelID), nil, nil, 1)
}

// ModifyChannel modifies an existing channel (discord.Channel) and returns its structure
func (c *Client) ModifyChannel(channelID discord.Snowflake, newChannel discord.ModifyChannel) (*discord.Channel, error) {
	return DoRequestAsStructure[discord.Channel](c, "PATCH", endpoints.Channel(channelID), newChannel, nil, 1, WithReason(newChannel.AuditLogReason))
}

// DeleteChannel deletes an existing channel (discord.Channel)
func (c *Client) DeleteChannel(channelID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.Channel(channelID), nil, nil, 1, WithReason(reason))
}

// GetChannelMessages returns a list of message structures (discord.Message) for the given channel ID
func (c *Client) GetChannelMessages(channelID discord.Snowflake) ([]discord.Message, error) {
	return DoRequestAsList[discord.Message](c, "GET", endpoints.ChannelMessages(channelID), nil, nil, 1)
}

// GetChannelMessage returns a message (discord.Message) for the given channel and message IDs
func (c *Client) GetChannelMessage(channelID, messageID discord.Snowflake) (*discord.Message, error) {
	return DoRequestAsStructure[discord.Message](c, "GET", endpoints.ChannelMessage(channelID, messageID), nil, nil, 1)
}

// CreateMessage posts a message (discord.Message) in a given channel ID and returns its structure
func (c *Client) CreateMessage(channelID discord.Snowflake, message discord.CreateMessage) (*discord.Message, error) {
	if len(message.Files) >= 1 {
		for id, file := range message.Files {
			message.Attachments = append(message.Attachments, discord.AttachmentSend{
				ID:          id,
				Description: file.Description,
			})
		}
		contentType, body, err := CreateMultipartBodyWithJSON(message, message.Files)
		if err != nil {
			return nil, err
		}
		return DoRequestWithFiles[discord.Message](c, "POST", endpoints.ChannelMessages(channelID), body, nil, 1, WithHeader("Content-Type", contentType))
	} else {
		return DoRequestAsStructure[discord.Message](c, "POST", endpoints.ChannelMessages(channelID), message, nil, 1)
	}
}

// CrosspostMessage cross-posts a given message ID from a given announcement channel (discord.ChannelTypeGuildAnnouncement) ID to following channels (discord.Channel) and returns a message structure (discord.Message)
func (c *Client) CrosspostMessage(channelID, messageID discord.Snowflake) (*discord.Message, error) {
	return DoRequestAsStructure[discord.Message](c, "POST", endpoints.CrosspostMessage(channelID, messageID), nil, nil, 1)
}

// CreateReaction creates a reaction (discord.Reaction) for the given channel and message IDs
func (c *Client) CreateReaction(channelID, messageID discord.Snowflake, emoji string) error {
	return DoEmptyRequest(c, "PUT", endpoints.Reaction(channelID, messageID, emoji), nil, nil, 1)
}

// DeleteOwnReaction deletes a reaction (discord.Reaction) the current user has made for the given channel and message IDs
func (c *Client) DeleteOwnReaction(channelID, messageID discord.Snowflake, emoji string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.Reaction(channelID, messageID, emoji), nil, nil, 1)
}

// DeleteUserReaction deletes a reaction (discord.Reaction) a given user ID has made for the given channel and message IDs
func (c *Client) DeleteUserReaction(channelID, messageID, userID discord.Snowflake, emoji string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.UserReaction(channelID, messageID, userID, emoji), nil, nil, 1)
}

// GetReactions returns a list of user structures (discord.User) that reacted to the given message ID in the given channel ID with the given emoji
func (c *Client) GetReactions(channelID, messageID discord.Snowflake, emoji string, after discord.Snowflake, limit int) ([]discord.User, error) {
	queryParams := make(QueryParameters)
	if after != 0 {
		queryParams["after"] = after.String()
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsList[discord.User](c, "GET", endpoints.ReactionsEmoji(channelID, messageID, emoji), nil, queryParams, 1)
}

// DeleteAllReactions deletes all reactions (discord.Reaction) for the given channel and message IDs
func (c *Client) DeleteAllReactions(channelID, messageID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", endpoints.Reactions(channelID, messageID), nil, nil, 1)
}

// DeleteAllReactionsForEmoji deletes all reactions (discord.Reaction) for the given channel and message IDs and the emoji
func (c *Client) DeleteAllReactionsForEmoji(channelID, messageID discord.Snowflake, emoji string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.ReactionsEmoji(channelID, messageID, emoji), nil, nil, 1)
}

// EditMessage edits a message (discord.Message) from the given channel and message IDs and returns its structure
func (c *Client) EditMessage(channelID, messageID discord.Snowflake, message discord.EditMessage) (*discord.Message, error) {
	if len(message.Files) >= 1 {
		for id, file := range message.Files {
			message.Attachments = append(message.Attachments, discord.AttachmentSend{
				ID:          id,
				Description: file.Description,
			})
		}
		contentType, body, err := CreateMultipartBodyWithJSON(message, message.Files)
		if err != nil {
			return nil, err
		}
		return DoRequestWithFiles[discord.Message](c, "PATCH", endpoints.ChannelMessage(channelID, messageID), body, nil, 1, WithHeader("Content-Type", contentType))
	} else {
		return DoRequestAsStructure[discord.Message](c, "PATCH", endpoints.ChannelMessage(channelID, messageID), message, nil, 1)
	}
}

// DeleteMessage deletes a message from the given channel and message IDs
func (c *Client) DeleteMessage(channelID, messageID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.ChannelMessage(channelID, messageID), nil, nil, 1, WithReason(reason))
}

// BulkDeleteMessages deletes multiple messages from the given channel ID in a single request
func (c *Client) BulkDeleteMessages(channelID discord.Snowflake, bulkDeleteMessages discord.BulkDeleteMessages) error {
	return DoEmptyRequest(c, "POST", endpoints.BulkDeleteMessages(channelID), bulkDeleteMessages, nil, 1, WithReason(bulkDeleteMessages.AuditLogReason))
}

// EditChannelPermissions edits the permissions for the given channel and overwrite IDs
func (c *Client) EditChannelPermissions(channelID, overwriteID discord.Snowflake, channelPermissions discord.EditChannelPermissions) error {
	return DoEmptyRequest(c, "PUT", endpoints.ChannelPermissions(channelID, overwriteID), channelPermissions, nil, 1, WithReason(channelPermissions.AuditLogReason))
}

// GetChannelInvites returns a list of invite with metadata structures (discord.InviteWithMetadata) for a given channel ID
func (c *Client) GetChannelInvites(channelID discord.Snowflake) ([]discord.InviteWithMetadata, error) {
	return DoRequestAsList[discord.InviteWithMetadata](c, "GET", endpoints.ChannelInvites(channelID), nil, nil, 1)
}

// CreateChannelInvite creates a new invite (discord.Invite) for the given channel ID and returns its new structure
func (c *Client) CreateChannelInvite(channelID discord.Snowflake, invite discord.CreateChannelInvite) (*discord.Invite, error) {
	return DoRequestAsStructure[discord.Invite](c, "POST", endpoints.ChannelInvites(channelID), invite, nil, 1, WithReason(invite.AuditLogReason))
}

// DeleteChannelPermissions delete the permissions for the given channel and overwrite IDs
func (c *Client) DeleteChannelPermissions(channelID, overwriteID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.ChannelPermissions(channelID, overwriteID), nil, nil, 1, WithReason(reason))
}

// FollowAnnouncementChannel follows the given channel ID to the given target channel ID and returns a followed channel structure (discord.FollowedChannel)
func (c *Client) FollowAnnouncementChannel(channelID, targetID discord.Snowflake) (*discord.FollowedChannel, error) {
	payload := struct {
		WebhookChannelID discord.Snowflake `json:"webhook_channel_id"`
	}{
		WebhookChannelID: targetID,
	}
	return DoRequestAsStructure[discord.FollowedChannel](c, "POST", endpoints.ChannelFollowers(channelID), payload, nil, 1)
}

// TriggerTyping triggers the "<User> is typing..." indicator for the given channel ID
func (c *Client) TriggerTyping(channelID discord.Snowflake) error {
	return DoEmptyRequest(c, "POST", endpoints.ChannelTyping(channelID), nil, nil, 1)
}

// GetPinnedMessages returns a list of message structures (discord.Message) that are pinned in the given channel ID
func (c *Client) GetPinnedMessages(channelID discord.Snowflake) ([]discord.Message, error) {
	return DoRequestAsList[discord.Message](c, "GET", endpoints.ChannelPins(channelID), nil, nil, 1)
}

// PinMessage pins the given message ID in the given channel ID
func (c *Client) PinMessage(channelID, messageID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "PUT", endpoints.ChannelPin(channelID, messageID), nil, nil, 1, WithReason(reason))
}

// UnpinMessage unpins the given message ID in the given channel ID
func (c *Client) UnpinMessage(channelID, messageID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.ChannelPin(channelID, messageID), nil, nil, 1, WithReason(reason))
}

// StartThreadFromMessage creates a new thread from the given message ID and returns a channel structure (discord.Channel)
func (c *Client) StartThreadFromMessage(channelID, messageID discord.Snowflake, thread discord.StartThreadFromMessage) (*discord.Channel, error) {
	return DoRequestAsStructure[discord.Channel](c, "POST", endpoints.ThreadFromMessage(channelID, messageID), thread, nil, 1, WithReason(thread.AuditLogReason))
}

// StartThreadWithoutMessage creates a new thread and returns a channel structure (discord.Channel)
func (c *Client) StartThreadWithoutMessage(channelID discord.Snowflake, thread discord.StartThreadWithoutMessage) (*discord.Channel, error) {
	return DoRequestAsStructure[discord.Channel](c, "POST", endpoints.ChannelThreads(channelID), thread, nil, 1, WithReason(thread.AuditLogReason))
}

// StartThreadInForumChannel creates a new thread in the given channel ID and returns a channel structure (discord.Channel) with a nested message structure (discord.Message)
func (c *Client) StartThreadInForumChannel(channelID discord.Snowflake, thread discord.StartThreadInForumChannel) (*discord.Channel, error) {
	if len(thread.Message.Files) >= 1 {
		for id, file := range thread.Message.Files {
			thread.Message.Attachments = append(thread.Message.Attachments, discord.AttachmentSend{
				ID:          id,
				Description: file.Description,
			})
		}
		contentType, body, err := CreateMultipartBodyWithJSON(thread, thread.Message.Files)
		if err != nil {
			return nil, err
		}
		return DoRequestWithFiles[discord.Channel](c, "POST", endpoints.ChannelThreads(channelID), body, nil, 1, WithHeader("Content-Type", contentType))
	} else {
		return DoRequestAsStructure[discord.Channel](c, "POST", endpoints.ChannelThreads(channelID), thread, nil, 1)
	}
}

// AddThreadMember adds a member to a thread given the thread and user IDs
func (c *Client) AddThreadMember(threadID, userID discord.Snowflake) error {
	return DoEmptyRequest(c, "PUT", endpoints.ThreadMember(threadID, userID), nil, nil, 1)
}

// LeaveThread removes the current user from a given thread given its ID
func (c *Client) LeaveThread(threadID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", endpoints.ThreadMemberSelf(threadID), nil, nil, 1)
}

// RemoveThreadMember removes another member from a thread given the thread and user IDs
func (c *Client) RemoveThreadMember(threadID, userID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", endpoints.ThreadMember(threadID, userID), nil, nil, 1)
}

// GetThreadMember returns a thread member structure (discord.ThreadMember) for a given thread and user ID if they are a member of said thread
func (c *Client) GetThreadMember(threadID, userID discord.Snowflake) (*discord.ThreadMember, error) {
	return DoRequestAsStructure[discord.ThreadMember](c, "GET", endpoints.ThreadMember(threadID, userID), nil, nil, 1)
}

// ListThreadMembers returns a list of thread member structures (discord.ThreadMember) that are members of the given thread ID
func (c *Client) ListThreadMembers(threadID, after discord.Snowflake, withMember bool, limit int) ([]discord.ThreadMember, error) {
	queryParams := make(QueryParameters)
	if after != 0 {
		queryParams["after"] = after.String()
	}
	if withMember {
		queryParams["with_member"] = "true"
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsList[discord.ThreadMember](c, "GET", endpoints.ThreadMembers(threadID), nil, queryParams, 1)
}

// ListPublicArchivedThreads returns an archived threads structure (discord.ArchivedThreads) which contains the public archived threads in the given channel ID
func (c *Client) ListPublicArchivedThreads(channelID discord.Snowflake, before *time.Time, limit int) (*discord.ArchivedThreads, error) {
	queryParams := make(QueryParameters)
	if before != nil {
		queryParams["before"] = before.Format(time.RFC3339)
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsStructure[discord.ArchivedThreads](c, "GET", endpoints.PublicArchivedThreads(channelID), nil, queryParams, 1)
}

// ListPrivateArchivedThreads returns an archived threads structure (discord.ArchivedThreads) which contains the private archived threads in the given channel ID
func (c *Client) ListPrivateArchivedThreads(channelID discord.Snowflake, before *time.Time, limit int) (*discord.ArchivedThreads, error) {
	queryParams := make(QueryParameters)
	if before != nil {
		queryParams["before"] = before.Format(time.RFC3339)
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsStructure[discord.ArchivedThreads](c, "GET", endpoints.PrivateArchivedThreads(channelID), nil, queryParams, 1)
}

// ListJoinedPrivateArchivedThreads returns an archived threads structure (discord.ArchivedThreads) which contains the private archived threads in the given channel ID that the user has joined
func (c *Client) ListJoinedPrivateArchivedThreads(channelID, before discord.Snowflake, limit int) (*discord.ArchivedThreads, error) {
	queryParams := make(QueryParameters)
	if before != 0 {
		queryParams["before"] = before.String()
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsStructure[discord.ArchivedThreads](c, "GET", endpoints.JoinedPrivateArchivedThreads(channelID), nil, queryParams, 1)
}
