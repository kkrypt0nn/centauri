package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

const (
	GuildsEndpoint = Endpoint + "guilds"
)

func Guilds() string {
	return GuildsEndpoint
}

func Guild(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d", GuildsEndpoint, guildID)
}

func GuildPreview(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/preview", GuildsEndpoint, guildID)
}

func GuildChannels(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/channels", GuildsEndpoint, guildID)
}

func ActiveThreads(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/threads/active", GuildsEndpoint, guildID)
}

func GuildMember(guildID, userID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/members/%d", GuildsEndpoint, guildID, userID)
}

func GuildMembersSearch(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/members/search", GuildsEndpoint, guildID)
}

func MemberRole(guildID, userID, roleID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/members/%d/roles/%d", GuildsEndpoint, guildID, userID, roleID)
}

func GuildBans(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/bans", GuildsEndpoint, guildID)
}

func GuildBan(guildID, userID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/bans/%d", GuildsEndpoint, guildID, userID)
}

func GuildRoles(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/roles", GuildsEndpoint, guildID)
}

func GuildRole(guildID, roleID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/roles/%d", GuildsEndpoint, guildID, roleID)
}

func GuildMFA(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/mfa", GuildsEndpoint, guildID)
}

func GuildPrune(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/prune", GuildsEndpoint, guildID)
}

func GuildRegions(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/regions", GuildsEndpoint, guildID)
}

func GuildInvites(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/invites", GuildsEndpoint, guildID)
}

func GuildIntegrations(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/integrations", GuildsEndpoint, guildID)
}

func GuildIntegration(guildID, integrationID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/integrations/%d", GuildsEndpoint, guildID, integrationID)
}

func GuildWidget(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/widget", GuildsEndpoint, guildID)
}

func GuildVanityURL(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/vanity-url", GuildsEndpoint, guildID)
}

func GuildWelcomeScreen(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/welcome-screen", GuildsEndpoint, guildID)
}

func GuildOnboarding(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/onboarding", GuildsEndpoint, guildID)
}

func GuildSelfVoiceState(guildID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/voice-states/@me", GuildsEndpoint, guildID)
}

func GuildVoiceState(guildID, userID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/voice-states/%d", GuildsEndpoint, guildID, userID)
}
