package rest

import (
	"strconv"
	"strings"

	"github.com/kkrypt0nn/centauri/discord"
)

const (
	GuildsEndpoint = Endpoint + "guilds"
)

// CreateGuild creates a guild (discord.Guild) and returns its structure
func (c *Client) CreateGuild(guild discord.CreateGuild) (*discord.Guild, error) {
	return DoRequestAsStructure[discord.Guild](c, "POST", GuildsEndpoint, guild, nil, 1)
}

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

// ModifyGuild modifies an existing guild (discord.Guild) from the given guild ID and returns its new structure
func (c *Client) ModifyGuild(guildID string, guild discord.ModifyGuild) (*discord.Guild, error) {
	return DoRequestAsStructure[discord.Guild](c, "PATCH", GuildsEndpoint+"/"+guildID, guild, nil, 1, WithReason(guild.AuditLogReason))
}

// DeleteGuild deletes an existing guild (discord.Guild) from the given guild ID
func (c *Client) DeleteGuild(guildID string) error {
	_, _, err := c.DoRequest("DELETE", GuildsEndpoint+"/"+guildID, nil, nil, 1)
	return err
}

// GetGuildChannels returns a list of channel structures (discord.Channel) for the given guild ID
func (c *Client) GetGuildChannels(guildID string) ([]discord.Channel, error) {
	return DoRequestAsList[discord.Channel](c, "GET", GuildsEndpoint+"/"+guildID+"/channels", nil, nil, 1)
}

// CreateGuildChannel creates a channel (discord.Channel) for the given guild ID and returns its structure
func (c *Client) CreateGuildChannel(guildID string, channel discord.CreateGuildChannel) (*discord.Channel, error) {
	return DoRequestAsStructure[discord.Channel](c, "POST", GuildsEndpoint+"/"+guildID+"/channels", channel, nil, 1, WithReason(channel.AuditLogReason))
}

// ModifyGuildChannelPositions modifies the position of existing channels (discord.Channel) for the given guild ID
func (c *Client) ModifyGuildChannelPositions(guildID string, channelPositions []discord.ModifyGuildChannelPosition) error {
	_, _, err := c.DoRequest("PATCH", GuildsEndpoint+"/"+guildID+"/channels", channelPositions, nil, 1)
	return err
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

// ModifyGuildMember modifies a guild member (discord.Member) for the given guild and user IDs and returns the new structure
func (c *Client) ModifyGuildMember(guildID, userID string, member discord.ModifyGuildMember) (*discord.Member, error) {
	return DoRequestAsStructure[discord.Member](c, "PATCH", GuildsEndpoint+"/"+guildID+"/members/"+userID, member, nil, 1, WithReason(member.AuditLogReason))
}

// AddGuildMemberRole adds a role (discord.Role) for the given guild, user and role IDs
func (c *Client) AddGuildMemberRole(guildID, userID, roleID, reason string) error {
	_, _, err := c.DoRequest("PUT", GuildsEndpoint+"/"+guildID+"/members/"+userID+"/roles/"+roleID, nil, nil, 1, WithReason(reason))
	return err
}

// RemoveGuildMemberRole removes a role (discord.Role) for the given guild, user and role IDs
func (c *Client) RemoveGuildMemberRole(guildID, userID, roleID, reason string) error {
	_, _, err := c.DoRequest("DELETE", GuildsEndpoint+"/"+guildID+"/members/"+userID+"/roles/"+roleID, nil, nil, 1, WithReason(reason))
	return err
}

// RemoveGuildMember removes a member (discord.Member) from a guild for the given guild and user IDs
func (c *Client) RemoveGuildMember(guildID, userID, reason string) error {
	_, _, err := c.DoRequest("DELETE", GuildsEndpoint+"/"+guildID+"/members/"+userID, nil, nil, 1, WithReason(reason))
	return err
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

// CreateGuildBan creates a guild ban (discord.GuildBan) for the given guild ID
func (c *Client) CreateGuildBan(guildID string, ban discord.CreateGuildBan) error {
	_, _, err := c.DoRequest("PUT", GuildsEndpoint+"/"+guildID+"/bans/"+ban.UserID, ban, nil, 1, WithReason(ban.AuditLogReason))
	return err
}

// RemoveGuildBan removes an existing guild ban (discord.GuildBan) for the given guild and user IDs
func (c *Client) RemoveGuildBan(guildID, userID, reason string) error {
	_, _, err := c.DoRequest("DELETE", GuildsEndpoint+"/"+guildID+"/bans/"+userID, nil, nil, 1, WithReason(reason))
	return err
}

// GetGuildRoles returns a list role structures (discord.Role)
func (c *Client) GetGuildRoles(guildID string) ([]discord.Role, error) {
	return DoRequestAsList[discord.Role](c, "GET", GuildsEndpoint+"/"+guildID+"/roles", nil, nil, 1)
}

// CreateGuildRole creates a role (discord.Role) for the given guild ID and returns its structure
func (c *Client) CreateGuildRole(guildID string, role discord.CreateGuildRole) (*discord.Role, error) {
	return DoRequestAsStructure[discord.Role](c, "POST", GuildsEndpoint+"/"+guildID+"/roles", role, nil, 1, WithReason(role.AuditLogReason))
}

// ModifyGuildRolePositions modifies the position of existing roles (discord.Role) for the given guild ID and returns the list of role structures in the guild
func (c *Client) ModifyGuildRolePositions(guildID string, rolePositions []discord.GuildRolePosition) ([]discord.Role, error) {
	return DoRequestAsList[discord.Role](c, "PATCH", GuildsEndpoint+"/"+guildID+"/roles", rolePositions, nil, 1)
}

// ModifyGuildRole modifies an existing role (discord.Role) for the given guild and role IDs and returns its new structure
func (c *Client) ModifyGuildRole(guildID, roleID string, role discord.ModifyGuildRole) (*discord.Role, error) {
	return DoRequestAsStructure[discord.Role](c, "PATCH", GuildsEndpoint+"/"+guildID+"/roles/"+roleID, role, nil, 1, WithReason(role.AuditLogReason))
}

// ModifyGuildMFALevel modifies the MFA level (discord.MFALevel) for the given guild ID and returns its new structure
func (c *Client) ModifyGuildMFALevel(guildID string, mfaLevel discord.ModifyGuildMFALevel) (*discord.MFALevel, error) {
	return DoRequestAsStructure[discord.MFALevel](c, "POST", GuildsEndpoint+"/"+guildID+"/mfa", mfaLevel, nil, 1, WithReason(mfaLevel.AuditLogReason))
}

// DeleteGuildRole deletes an existing role (discord.Role) for the given guild and role IDs
func (c *Client) DeleteGuildRole(guildID, roleID, reason string) error {
	_, _, err := c.DoRequest("DELETE", GuildsEndpoint+"/"+guildID+"/roles/"+roleID, nil, nil, 1, WithReason(reason))
	return err
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

// BeginGuildPrune begins a guild prune operation (discord.GuildPrune) for the given guild ID and returns the number of members that were removed
func (c *Client) BeginGuildPrune(guildID string, prune discord.BeginGuildPrune) (*discord.GuildPrune, error) {
	return DoRequestAsStructure[discord.GuildPrune](c, "POST", GuildsEndpoint+"/"+guildID+"/prune", prune, nil, 1, WithReason(prune.AuditLogReason))
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

// DeleteGuildIntegration deletes an existing integration (discord.Integration) for the given guild and integration IDs
func (c *Client) DeleteGuildIntegration(guildID, integrationID, reason string) error {
	_, _, err := c.DoRequest("DELETE", GuildsEndpoint+"/"+guildID+"/integrations/"+integrationID, nil, nil, 1, WithReason(reason))
	return err
}

// GetGuildWidgetSetting returns a widget setting structure (discord.WidgetSetting) for the given guild ID
func (c *Client) GetGuildWidgetSetting(guildID string) (*discord.WidgetSetting, error) {
	return DoRequestAsStructure[discord.WidgetSetting](c, "GET", GuildsEndpoint+"/"+guildID+"/widget", nil, nil, 1)
}

// ModifyGuildWidget modifies an existing widget (discord.WidgetSetting) for the given guild ID and returns its new structure
func (c *Client) ModifyGuildWidget(guildID string, widget discord.ModifyGuildWidget) (*discord.WidgetSetting, error) {
	return DoRequestAsStructure[discord.WidgetSetting](c, "PATCH", GuildsEndpoint+"/"+guildID+"/widget", widget, nil, 1, WithReason(widget.AuditLogReason))
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

// ModifyGuildWelcomeScreen modifies an existing welcome screen (discord.WelcomeScreen) for the given guild ID and returns its new structure
func (c *Client) ModifyGuildWelcomeScreen(guildID string, welcomeScreen discord.ModifyGuildWelcomeScreen) (*discord.WelcomeScreen, error) {
	return DoRequestAsStructure[discord.WelcomeScreen](c, "PATCH", GuildsEndpoint+"/"+guildID+"/welcome-screen", welcomeScreen, nil, 1, WithReason(welcomeScreen.AuditLogReason))
}

// GetGuildOnboarding returns an onboarding structure (discord.Onboarding) for the given guild ID
func (c *Client) GetGuildOnboarding(guildID string) (*discord.Onboarding, error) {
	return DoRequestAsStructure[discord.Onboarding](c, "GET", GuildsEndpoint+"/"+guildID+"/onboarding", nil, nil, 1)
}

// ModifyCurrentUserVoiceState modifies the voice state of the current user for the given guild ID
func (c *Client) ModifyCurrentUserVoiceState(guildID string, voiceState discord.ModifyCurrentUserVoiceState) error {
	_, _, err := c.DoRequest("PATCH", GuildsEndpoint+"/"+guildID+"/voice-states/@me", voiceState, nil, 1)
	return err
}

// ModifyUserVoiceState modifies the voice state of the current user for the given guild and user IDs
func (c *Client) ModifyUserVoiceState(guildID, userID string, voiceState discord.ModifyUserVoiceState) error {
	_, _, err := c.DoRequest("PATCH", GuildsEndpoint+"/"+guildID+"/voice-states/"+userID, voiceState, nil, 1)
	return err
}
