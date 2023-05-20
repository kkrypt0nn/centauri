package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"strconv"
	"strings"
)

const (
	GuildsEndpoint = Endpoint + "guilds"
)

// GetGuild returns a guild structure (discord.Guild) for the given guild ID
func (c *Client) GetGuild(guildID string, withCounts bool) (*discord.Guild, error) {
	queryParams := make(QueryParameters)
	if withCounts {
		queryParams["with_counts"] = "true"
	}
	return DoRequestAsStructure[discord.Guild](c, "GET", GuildsEndpoint+"/"+guildID, nil, queryParams, 1)
}

// GetGuildPreview returns a guild preview structure (discord.GuildPreview) for the given guild ID
func (c *Client) GetGuildPreview(guildID string) (*discord.GuildPreview, error) {
	return DoRequestAsStructure[discord.GuildPreview](c, "GET", GuildsEndpoint+"/"+guildID+"/preview", nil, nil, 1)
}

// GetGuildChannels returns a list of channel structures (discord.Channel) for the given guild ID
func (c *Client) GetGuildChannels(guildID string) ([]discord.Channel, error) {
	return DoRequestAsList[discord.Channel](c, "GET", GuildsEndpoint+"/"+guildID+"/channels", nil, nil, 1)
}

// ListActiveThreads returns an active threads structure (discord.ActiveThreads) for the given guild ID
func (c *Client) ListActiveThreads(guildID string) (*discord.ActiveThreads, error) {
	return DoRequestAsStructure[discord.ActiveThreads](c, "GET", GuildsEndpoint+"/"+guildID+"/threads/active", nil, nil, 1)
}

// GetGuildMember returns a member structure (discord.Member) for the given guild and user IDs
func (c *Client) GetGuildMember(guildID, userID string) (*discord.Member, error) {
	return DoRequestAsStructure[discord.Member](c, "GET", GuildsEndpoint+"/"+guildID+"/members/"+userID, nil, nil, 1)
}

// SearchGuildMember returns a list of member structures (discord.Member) matching the given query in the given guild ID
func (c *Client) SearchGuildMember(guildID, query string, limit int) ([]discord.Member, error) {
	queryParams := make(QueryParameters)
	queryParams["query"] = query
	if limit >= 1 && limit <= 1000 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsList[discord.Member](c, "GET", GuildsEndpoint+"/"+guildID+"/members/search", nil, queryParams, 1)
}

// GetGuildBans returns a list of guild ban structures (discord.GuildBan) for the given guild ID
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
	return DoRequestAsList[discord.GuildBan](c, "GET", GuildsEndpoint+"/"+guildID+"/bans", nil, queryParams, 1)
}

// GetGuildBan returns a guild ban structure (discord.GuildBan) for the given guild and ban ID
func (c *Client) GetGuildBan(guildID, banID string) (*discord.GuildBan, error) {
	return DoRequestAsStructure[discord.GuildBan](c, "GET", GuildsEndpoint+"/"+guildID+"/bans/"+banID, nil, nil, 1)
}

// GetGuildRoles returns a list role structures (discord.Role)
func (c *Client) GetGuildRoles(guildID string) ([]discord.Role, error) {
	return DoRequestAsList[discord.Role](c, "GET", GuildsEndpoint+"/"+guildID+"/roles", nil, nil, 1)
}

// GetGuildPruneCount returns a guild prune structure (discord.GuildPrune) for the given guild IDs
func (c *Client) GetGuildPruneCount(guildID string, days int, includeRoles []string) (*discord.GuildPrune, error) {
	queryParams := make(QueryParameters)
	if days >= 1 && days <= 30 {
		queryParams["days"] = strconv.Itoa(days)
	}
	if len(includeRoles) >= 1 {
		queryParams["include_roles"] = strings.Join(includeRoles, ",")
	}
	return DoRequestAsStructure[discord.GuildPrune](c, "GET", GuildsEndpoint+"/"+guildID+"/prune", nil, queryParams, 1)
}

// GetGuildVoiceRegions returns a list of voice region structures (discord.VoiceRegion) for the given guild ID
func (c *Client) GetGuildVoiceRegions(guildID string) ([]discord.VoiceRegion, error) {
	return DoRequestAsList[discord.VoiceRegion](c, "GET", GuildsEndpoint+"/"+guildID+"/regions", nil, nil, 1)
}

// GetGuildInvites returns a list of invite with metadata structures (discord.InviteWithMetadata) for the given guild ID
func (c *Client) GetGuildInvites(guildID string) ([]discord.InviteWithMetadata, error) {
	return DoRequestAsList[discord.InviteWithMetadata](c, "GET", GuildsEndpoint+"/"+guildID+"/invites", nil, nil, 1)
}

// GetGuildIntegrations returns a list of integrations structures (discord.Integration) for the given guild ID
func (c *Client) GetGuildIntegrations(guildID string) ([]discord.Integration, error) {
	return DoRequestAsList[discord.Integration](c, "GET", GuildsEndpoint+"/"+guildID+"/integrations", nil, nil, 1)
}

// GetGuildWidgetSetting returns a widget setting structure (discord.WidgetSetting) for the given guild ID
func (c *Client) GetGuildWidgetSetting(guildID string) (*discord.WidgetSetting, error) {
	return DoRequestAsStructure[discord.WidgetSetting](c, "GET", GuildsEndpoint+"/"+guildID+"/widget", nil, nil, 1)
}

// GetGuildWidget returns a widget structure (discord.Widget) for the given guild ID
func (c *Client) GetGuildWidget(guildID string) (*discord.Widget, error) {
	return DoRequestAsStructure[discord.Widget](c, "GET", GuildsEndpoint+"/"+guildID+"/widget.json", nil, nil, 1)
}

// GetGuildVanityURL returns the vanity invite, which is a partial invite structure (discord.PartialInvite), for the given guild ID
func (c *Client) GetGuildVanityURL(guildID string) (*discord.PartialInvite, error) {
	return DoRequestAsStructure[discord.PartialInvite](c, "GET", GuildsEndpoint+"/"+guildID+"/vanity-url", nil, nil, 1)
}

// GetGuildWelcomeScreen returns a welcome screen structure (discord.WelcomeScreen) for the given guild ID
func (c *Client) GetGuildWelcomeScreen(guildID string) (*discord.WelcomeScreen, error) {
	return DoRequestAsStructure[discord.WelcomeScreen](c, "GET", GuildsEndpoint+"/"+guildID+"/welcome-screen", nil, nil, 1)
}

// GetGuildOnboarding returns a onboarding structure (discord.Onboarding) for the given guild ID
func (c *Client) GetGuildOnboarding(guildID string) (*discord.Onboarding, error) {
	return DoRequestAsStructure[discord.Onboarding](c, "GET", GuildsEndpoint+"/"+guildID+"/onboarding", nil, nil, 1)
}
