package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"strconv"
	"strings"
)

const (
	GuildsEndpoint = Endpoint + "guilds"
)

// GetGuild returns a discord.Guild based on its ID
func (c *Client) GetGuild(guildID string, withCounts bool) (*discord.Guild, error) {
	queryParams := make(QueryParameters)
	if withCounts {
		queryParams["with_counts"] = "true"
	}
	return DoRequestAs[discord.Guild](c, "GET", GuildsEndpoint+"/"+guildID, queryParams, 1)
}

// GetGuildPreview returns a discord.GuildPreview based on its ID
func (c *Client) GetGuildPreview(guildID string) (*discord.GuildPreview, error) {
	return DoRequestAs[discord.GuildPreview](c, "GET", GuildsEndpoint+"/"+guildID+"/preview", nil, 1)
}

// GetGuildChannels returns a list of discord.Channel based on the guild ID
func (c *Client) GetGuildChannels(guildID string) ([]discord.Channel, error) {
	return DoRequestAsList[discord.Channel](c, "GET", GuildsEndpoint+"/"+guildID+"/channels", nil, 1)
}

// ListActiveThreads returns a discord.ActiveThreads structure
func (c *Client) ListActiveThreads(guildID string) (*discord.ActiveThreads, error) {
	return DoRequestAs[discord.ActiveThreads](c, "GET", GuildsEndpoint+"/"+guildID+"/threads/active", nil, 1)
}

// GetGuildMember returns a discord.Member structure
func (c *Client) GetGuildMember(guildID, userID string) (*discord.Member, error) {
	return DoRequestAs[discord.Member](c, "GET", GuildsEndpoint+"/"+guildID+"/members/"+userID, nil, 1)
}

// SearchGuildMember returns a list of discord.Member structure matching the given query
func (c *Client) SearchGuildMember(guildID, query string, limit int) ([]discord.Member, error) {
	queryParams := make(QueryParameters)
	queryParams["query"] = query
	if limit >= 1 && limit <= 1000 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsList[discord.Member](c, "GET", GuildsEndpoint+"/"+guildID+"/members/search", queryParams, 1)
}

// GetGuildBans returns a list of discord.GuildBan structures
func (c *Client) GetGuildBans(guildID, before, after string, limit int) ([]discord.GuildBan, error) {
	queryParams := make(QueryParameters)
	if before != "" {
		queryParams["before"] = before
	}
	if after != "" {
		queryParams["after"] = after
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsList[discord.GuildBan](c, "GET", GuildsEndpoint+"/"+guildID+"/bans", queryParams, 1)
}

// GetGuildBan returns a discord.GuildBan structure matching the given ID
func (c *Client) GetGuildBan(guildID, banID string) (*discord.GuildBan, error) {
	return DoRequestAs[discord.GuildBan](c, "GET", GuildsEndpoint+"/"+guildID+"/bans/"+banID, nil, 1)
}

// GetGuildRoles returns a list of discord.Role structures
func (c *Client) GetGuildRoles(guildID string) ([]discord.Role, error) {
	return DoRequestAsList[discord.Role](c, "GET", GuildsEndpoint+"/"+guildID+"/roles", nil, 1)
}

// GetGuildPruneCount returns a discord.GuildPrune structure
func (c *Client) GetGuildPruneCount(guildID string, days int, includeRoles []string) (*discord.GuildPrune, error) {
	queryParams := make(QueryParameters)
	if days >= 1 && days <= 30 {
		queryParams["days"] = strconv.Itoa(days)
	}
	if len(includeRoles) >= 1 {
		queryParams["include_roles"] = strings.Join(includeRoles, ",")
	}
	return DoRequestAs[discord.GuildPrune](c, "GET", GuildsEndpoint+"/"+guildID+"/prune", queryParams, 1)
}

// GetGuildVoiceRegions returns a list of discord.VoiceRegion structures
func (c *Client) GetGuildVoiceRegions(guildID string) ([]discord.VoiceRegion, error) {
	return DoRequestAsList[discord.VoiceRegion](c, "GET", GuildsEndpoint+"/"+guildID+"/regions", nil, 1)
}

// GetGuildInvites returns a list of discord.InviteWithMetadata structures
func (c *Client) GetGuildInvites(guildID string) ([]discord.InviteWithMetadata, error) {
	return DoRequestAsList[discord.InviteWithMetadata](c, "GET", GuildsEndpoint+"/"+guildID+"/invites", nil, 1)
}

// GetGuildIntegrations returns a list of discord.Integration structures
func (c *Client) GetGuildIntegrations(guildID string) ([]discord.Integration, error) {
	return DoRequestAsList[discord.Integration](c, "GET", GuildsEndpoint+"/"+guildID+"/integrations", nil, 1)
}

// GetGuildWidgetSetting returns a discord.WidgetSetting structure
func (c *Client) GetGuildWidgetSetting(guildID string) (*discord.WidgetSetting, error) {
	return DoRequestAs[discord.WidgetSetting](c, "GET", GuildsEndpoint+"/"+guildID+"/widget", nil, 1)
}

// GetGuildWidget returns a discord.Widget structure
func (c *Client) GetGuildWidget(guildID string) (*discord.Widget, error) {
	return DoRequestAs[discord.Widget](c, "GET", GuildsEndpoint+"/"+guildID+"/widget.json", nil, 1)
}

// GetGuildVanityURL returns a discord.PartialInvite structure
func (c *Client) GetGuildVanityURL(guildID string) (*discord.PartialInvite, error) {
	return DoRequestAs[discord.PartialInvite](c, "GET", GuildsEndpoint+"/"+guildID+"/vanity-url", nil, 1)
}

// GetGuildWelcomeScreen returns a discord.WelcomeScreen structure
func (c *Client) GetGuildWelcomeScreen(guildID string) (*discord.WelcomeScreen, error) {
	return DoRequestAs[discord.WelcomeScreen](c, "GET", GuildsEndpoint+"/"+guildID+"/welcome-screen", nil, 1)
}

// GetGuildOnboarding returns a discord.Onboarding structure
func (c *Client) GetGuildOnboarding(guildID string) (*discord.Onboarding, error) {
	return DoRequestAs[discord.Onboarding](c, "GET", GuildsEndpoint+"/"+guildID+"/onboarding", nil, 1)
}
