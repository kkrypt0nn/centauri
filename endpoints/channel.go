package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
	"net/url"
)

const (
	ChannelsEndpoint = Endpoint + "channels"
)

func Channel(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d", ChannelsEndpoint, channelID)
}

func ChannelMessages(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/messages", ChannelsEndpoint, channelID)
}

func ChannelMessage(channelID, messageID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/messages/%d", ChannelsEndpoint, channelID, messageID)
}

func CrosspostMessage(channelID, messageID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/messages/%d/crosspost", ChannelsEndpoint, channelID, messageID)
}

func Reactions(channelID, messageID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/messages/%d/reactions", ChannelsEndpoint, channelID, messageID)
}

func ReactionsEmoji(channelID, messageID discord.Snowflake, emoji string) string {
	return fmt.Sprintf("%s/%d/messages/%d/reactions/%s", ChannelsEndpoint, channelID, messageID, url.QueryEscape(emoji))
}

func Reaction(channelID, messageID discord.Snowflake, emoji string) string {
	return fmt.Sprintf("%s/%d/messages/%d/reactions/%s/@me", ChannelsEndpoint, channelID, messageID, url.QueryEscape(emoji))
}

func UserReaction(channelID, messageID, userID discord.Snowflake, emoji string) string {
	return fmt.Sprintf("%s/%d/messages/%d/reactions/%s/%d", ChannelsEndpoint, channelID, messageID, url.QueryEscape(emoji), userID)
}

func BulkDeleteMessages(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/messages/bulk-delete", ChannelsEndpoint, channelID)
}

func ChannelPermissions(channelID, overwriteID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/permissions/%d", ChannelsEndpoint, channelID, overwriteID)
}

func ChannelInvites(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/invites", ChannelsEndpoint, channelID)
}

func ChannelFollowers(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/followers", ChannelsEndpoint, channelID)
}

func ChannelTyping(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/typing", ChannelsEndpoint, channelID)
}

func ChannelPins(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/pins", ChannelsEndpoint, channelID)
}

func ChannelPin(channelID, messageID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/pins/%d", ChannelsEndpoint, channelID, messageID)
}

func ThreadFromMessage(channelID, messageID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/messages/%d/threads", ChannelsEndpoint, channelID, messageID)
}

func ChannelThreads(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/threads", ChannelsEndpoint, channelID)
}

func ThreadMembers(threadID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/thread-members", ChannelsEndpoint, threadID)
}

func ThreadMember(threadID, userID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/thread-members/%d", ChannelsEndpoint, threadID, userID)
}

func ThreadMemberSelf(threadID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/thread-members/@me", ChannelsEndpoint, threadID)
}

func PublicArchivedThreads(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/threads/archived/public", ChannelsEndpoint, channelID)
}

func PrivateArchivedThreads(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/threads/archived/private", ChannelsEndpoint, channelID)
}

func JoinedPrivateArchivedThreads(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/users/@me/threads/archived/private", ChannelsEndpoint, channelID)
}
