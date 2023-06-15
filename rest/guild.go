package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"strconv"
)

const (
	GuildsEndpoint = Endpoint + "guilds"
)

// CreateGuild creates a guild (discord.Guild) and returns its structure
func (c *Client) CreateGuild(guild discord.CreateGuild) (*discord.Guild, error) {
	return DoRequestAsStructure[discord.Guild](c, "POST", GuildsEndpoint, guild, nil, 1)
}

// GetGuild returns a guild structure (discord.Guild) for the given guild ID
func (c *Client) GetGuild(guildID discord.Snowflake, withCounts bool) (*discord.Guild, error) {
	queryParams := make(QueryParameters)
	if withCounts {
		queryParams["with_counts"] = "true"
	}
	return DoRequestAsStructure[discord.Guild](c, "GET", GuildsEndpoint+"/"+guildID.String(), nil, queryParams, 1)
}

// GetGuildPreview returns a guild preview structure (discord.GuildPreview) for the given guild ID
func (c *Client) GetGuildPreview(guildID discord.Snowflake) (*discord.GuildPreview, error) {
	return DoRequestAsStructure[discord.GuildPreview](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/preview", nil, nil, 1)
}

// ModifyGuild modifies an existing guild (discord.Guild) from the given guild ID and returns its new structure
func (c *Client) ModifyGuild(guildID discord.Snowflake, guild discord.ModifyGuild) (*discord.Guild, error) {
	return DoRequestAsStructure[discord.Guild](c, "PATCH", GuildsEndpoint+"/"+guildID.String(), guild, nil, 1, WithReason(guild.AuditLogReason))
}

// DeleteGuild deletes an existing guild (discord.Guild) from the given guild ID
func (c *Client) DeleteGuild(guildID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", GuildsEndpoint+"/"+guildID.String(), nil, nil, 1)
}

// GetGuildChannels returns a list of channel structures (discord.Channel) for the given guild ID
func (c *Client) GetGuildChannels(guildID discord.Snowflake) ([]discord.Channel, error) {
	return DoRequestAsList[discord.Channel](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/channels", nil, nil, 1)
}

// CreateGuildChannel creates a channel (discord.Channel) for the given guild ID and returns its structure
func (c *Client) CreateGuildChannel(guildID discord.Snowflake, channel discord.CreateGuildChannel) (*discord.Channel, error) {
	return DoRequestAsStructure[discord.Channel](c, "POST", GuildsEndpoint+"/"+guildID.String()+"/channels", channel, nil, 1, WithReason(channel.AuditLogReason))
}

// ModifyGuildChannelPositions modifies the position of existing channels (discord.Channel) for the given guild ID
func (c *Client) ModifyGuildChannelPositions(guildID discord.Snowflake, channelPositions []discord.ModifyGuildChannelPosition) error {
	return DoEmptyRequest(c, "PATCH", GuildsEndpoint+"/"+guildID.String()+"/channels", channelPositions, nil, 1)
}

// ListActiveThreads returns an active threads structure (discord.ActiveThreads) for the given guild ID
func (c *Client) ListActiveThreads(guildID discord.Snowflake) (*discord.ActiveThreads, error) {
	return DoRequestAsStructure[discord.ActiveThreads](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/threads/active", nil, nil, 1)
}

// GetGuildMember returns a member structure (discord.Member) for the given guild and user IDs
func (c *Client) GetGuildMember(guildID, userID discord.Snowflake) (*discord.Member, error) {
	member, err := DoRequestAsStructure[discord.Member](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/members/"+userID.String(), nil, nil, 1)
	if member != nil {
		member.GuildID = guildID
	}
	return member, err
}

// SearchGuildMember returns a list of member structures (discord.Member) matching the given query in the given guild ID
func (c *Client) SearchGuildMember(guildID discord.Snowflake, query string, limit int) ([]discord.Member, error) {
	queryParams := make(QueryParameters)
	queryParams["query"] = query
	if limit >= 1 && limit <= 1000 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	members, err := DoRequestAsList[discord.Member](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/members/search", nil, queryParams, 1)
	if members != nil {
		for i := 0; i < len(members); i++ {
			members[i].GuildID = guildID
		}
	}
	return members, err
}

// ModifyGuildMember modifies a guild member (discord.Member) for the given guild and user IDs and returns the new structure
func (c *Client) ModifyGuildMember(guildID, userID discord.Snowflake, member discord.ModifyGuildMember) (*discord.Member, error) {
	newMember, err := DoRequestAsStructure[discord.Member](c, "PATCH", GuildsEndpoint+"/"+guildID.String()+"/members/"+userID.String(), member, nil, 1, WithReason(member.AuditLogReason))
	if newMember != nil {
		newMember.GuildID = guildID
	}
	return newMember, err
}

// AddGuildMemberRole adds a role (discord.Role) for the given guild, user and role IDs
func (c *Client) AddGuildMemberRole(guildID, userID, roleID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "PUT", GuildsEndpoint+"/"+guildID.String()+"/members/"+userID.String()+"/roles/"+roleID.String(), nil, nil, 1, WithReason(reason))
}

// RemoveGuildMemberRole removes a role (discord.Role) for the given guild, user and role IDs
func (c *Client) RemoveGuildMemberRole(guildID, userID, roleID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", GuildsEndpoint+"/"+guildID.String()+"/members/"+userID.String()+"/roles/"+roleID.String(), nil, nil, 1, WithReason(reason))
}

// RemoveGuildMember removes a member (discord.Member) from a guild for the given guild and user IDs
func (c *Client) RemoveGuildMember(guildID, userID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", GuildsEndpoint+"/"+guildID.String()+"/members/"+userID.String(), nil, nil, 1, WithReason(reason))
}

// GetGuildBans returns a list of guild ban structures (discord.GuildBan) for the given guild ID
func (c *Client) GetGuildBans(guildID, before, after discord.Snowflake, limit int) ([]discord.GuildBan, error) {
	queryParams := make(QueryParameters)
	if before != 0 {
		queryParams["before"] = before.String()
	}
	if after != 0 {
		queryParams["after"] = after.String()
	}
	if limit >= 1 && limit <= 100 {
		queryParams["limit"] = strconv.Itoa(limit)
	}
	return DoRequestAsList[discord.GuildBan](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/bans", nil, queryParams, 1)
}

// GetGuildBan returns a guild ban structure (discord.GuildBan) for the given guild and ban ID
func (c *Client) GetGuildBan(guildID, banID discord.Snowflake) (*discord.GuildBan, error) {
	return DoRequestAsStructure[discord.GuildBan](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/bans/"+banID.String(), nil, nil, 1)
}

// CreateGuildBan creates a guild ban (discord.GuildBan) for the given guild ID
func (c *Client) CreateGuildBan(guildID discord.Snowflake, ban discord.CreateGuildBan) error {
	return DoEmptyRequest(c, "PUT", GuildsEndpoint+"/"+guildID.String()+"/bans/"+ban.UserID.String(), ban, nil, 1, WithReason(ban.AuditLogReason))
}

// RemoveGuildBan removes an existing guild ban (discord.GuildBan) for the given guild and user IDs
func (c *Client) RemoveGuildBan(guildID, userID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", GuildsEndpoint+"/"+guildID.String()+"/bans/"+userID.String(), nil, nil, 1, WithReason(reason))
}

// GetGuildRoles returns a list role structures (discord.Role)
func (c *Client) GetGuildRoles(guildID discord.Snowflake) ([]discord.Role, error) {
	return DoRequestAsList[discord.Role](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/roles", nil, nil, 1)
}

// CreateGuildRole creates a role (discord.Role) for the given guild ID and returns its structure
func (c *Client) CreateGuildRole(guildID discord.Snowflake, role discord.CreateGuildRole) (*discord.Role, error) {
	return DoRequestAsStructure[discord.Role](c, "POST", GuildsEndpoint+"/"+guildID.String()+"/roles", role, nil, 1, WithReason(role.AuditLogReason))
}

// ModifyGuildRolePositions modifies the position of existing roles (discord.Role) for the given guild ID and returns the list of role structures in the guild
func (c *Client) ModifyGuildRolePositions(guildID discord.Snowflake, rolePositions []discord.GuildRolePosition) ([]discord.Role, error) {
	return DoRequestAsList[discord.Role](c, "PATCH", GuildsEndpoint+"/"+guildID.String()+"/roles", rolePositions, nil, 1)
}

// ModifyGuildRole modifies an existing role (discord.Role) for the given guild and role IDs and returns its new structure
func (c *Client) ModifyGuildRole(guildID, roleID discord.Snowflake, role discord.ModifyGuildRole) (*discord.Role, error) {
	return DoRequestAsStructure[discord.Role](c, "PATCH", GuildsEndpoint+"/"+guildID.String()+"/roles/"+roleID.String(), role, nil, 1, WithReason(role.AuditLogReason))
}

// ModifyGuildMFALevel modifies the MFA level (discord.MFALevel) for the given guild ID and returns its new structure
func (c *Client) ModifyGuildMFALevel(guildID discord.Snowflake, mfaLevel discord.ModifyGuildMFALevel) (*discord.MFALevel, error) {
	return DoRequestAsStructure[discord.MFALevel](c, "POST", GuildsEndpoint+"/"+guildID.String()+"/mfa", mfaLevel, nil, 1, WithReason(mfaLevel.AuditLogReason))
}

// DeleteGuildRole deletes an existing role (discord.Role) for the given guild and role IDs
func (c *Client) DeleteGuildRole(guildID, roleID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", GuildsEndpoint+"/"+guildID.String()+"/roles/"+roleID.String(), nil, nil, 1, WithReason(reason))
}

// GetGuildPruneCount returns a guild prune structure (discord.GuildPrune) for the given guild IDs
func (c *Client) GetGuildPruneCount(guildID discord.Snowflake, days int, includeRoles discord.ArraySnowflakes) (*discord.GuildPrune, error) {
	queryParams := make(QueryParameters)
	if days >= 1 && days <= 30 {
		queryParams["days"] = strconv.Itoa(days)
	}
	if len(includeRoles) >= 1 {
		queryParams["include_roles"] = includeRoles.String()
	}
	return DoRequestAsStructure[discord.GuildPrune](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/prune", nil, queryParams, 1)
}

// BeginGuildPrune begins a guild prune operation (discord.GuildPrune) for the given guild ID and returns the number of members that were removed
func (c *Client) BeginGuildPrune(guildID discord.Snowflake, prune discord.BeginGuildPrune) (*discord.GuildPrune, error) {
	return DoRequestAsStructure[discord.GuildPrune](c, "POST", GuildsEndpoint+"/"+guildID.String()+"/prune", prune, nil, 1, WithReason(prune.AuditLogReason))
}

// GetGuildVoiceRegions returns a list of voice region structures (discord.VoiceRegion) for the given guild ID
func (c *Client) GetGuildVoiceRegions(guildID discord.Snowflake) ([]discord.VoiceRegion, error) {
	return DoRequestAsList[discord.VoiceRegion](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/regions", nil, nil, 1)
}

// GetGuildInvites returns a list of invite with metadata structures (discord.InviteWithMetadata) for the given guild ID
func (c *Client) GetGuildInvites(guildID discord.Snowflake) ([]discord.InviteWithMetadata, error) {
	return DoRequestAsList[discord.InviteWithMetadata](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/invites", nil, nil, 1)
}

// GetGuildIntegrations returns a list of integrations structures (discord.Integration) for the given guild ID
func (c *Client) GetGuildIntegrations(guildID discord.Snowflake) ([]discord.Integration, error) {
	return DoRequestAsList[discord.Integration](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/integrations", nil, nil, 1)
}

// DeleteGuildIntegration deletes an existing integration (discord.Integration) for the given guild and integration IDs
func (c *Client) DeleteGuildIntegration(guildID, integrationID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", GuildsEndpoint+"/"+guildID.String()+"/integrations/"+integrationID.String(), nil, nil, 1, WithReason(reason))
}

// GetGuildWidgetSetting returns a widget setting structure (discord.WidgetSetting) for the given guild ID
func (c *Client) GetGuildWidgetSetting(guildID discord.Snowflake) (*discord.WidgetSetting, error) {
	return DoRequestAsStructure[discord.WidgetSetting](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/widget", nil, nil, 1)
}

// ModifyGuildWidget modifies an existing widget (discord.WidgetSetting) for the given guild ID and returns its new structure
func (c *Client) ModifyGuildWidget(guildID discord.Snowflake, widget discord.ModifyGuildWidget) (*discord.WidgetSetting, error) {
	return DoRequestAsStructure[discord.WidgetSetting](c, "PATCH", GuildsEndpoint+"/"+guildID.String()+"/widget", widget, nil, 1, WithReason(widget.AuditLogReason))
}

// GetGuildWidget returns a widget structure (discord.Widget) for the given guild ID
func (c *Client) GetGuildWidget(guildID discord.Snowflake) (*discord.Widget, error) {
	return DoRequestAsStructure[discord.Widget](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/widget.json", nil, nil, 1)
}

// GetGuildVanityURL returns the vanity invite, which is a partial invite structure (discord.PartialInvite), for the given guild ID
func (c *Client) GetGuildVanityURL(guildID discord.Snowflake) (*discord.PartialInvite, error) {
	return DoRequestAsStructure[discord.PartialInvite](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/vanity-url", nil, nil, 1)
}

// GetGuildWelcomeScreen returns a welcome screen structure (discord.WelcomeScreen) for the given guild ID
func (c *Client) GetGuildWelcomeScreen(guildID discord.Snowflake) (*discord.WelcomeScreen, error) {
	return DoRequestAsStructure[discord.WelcomeScreen](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/welcome-screen", nil, nil, 1)
}

// ModifyGuildWelcomeScreen modifies an existing welcome screen (discord.WelcomeScreen) for the given guild ID and returns its new structure
func (c *Client) ModifyGuildWelcomeScreen(guildID discord.Snowflake, welcomeScreen discord.ModifyGuildWelcomeScreen) (*discord.WelcomeScreen, error) {
	return DoRequestAsStructure[discord.WelcomeScreen](c, "PATCH", GuildsEndpoint+"/"+guildID.String()+"/welcome-screen", welcomeScreen, nil, 1, WithReason(welcomeScreen.AuditLogReason))
}

// GetGuildOnboarding returns an onboarding structure (discord.Onboarding) for the given guild ID
func (c *Client) GetGuildOnboarding(guildID discord.Snowflake) (*discord.Onboarding, error) {
	return DoRequestAsStructure[discord.Onboarding](c, "GET", GuildsEndpoint+"/"+guildID.String()+"/onboarding", nil, nil, 1)
}

// ModifyCurrentUserVoiceState modifies the voice state of the current user for the given guild ID
func (c *Client) ModifyCurrentUserVoiceState(guildID discord.Snowflake, voiceState discord.ModifyCurrentUserVoiceState) error {
	return DoEmptyRequest(c, "PATCH", GuildsEndpoint+"/"+guildID.String()+"/voice-states/@me", voiceState, nil, 1)
}

// ModifyUserVoiceState modifies the voice state of the current user for the given guild and user IDs
func (c *Client) ModifyUserVoiceState(guildID, userID discord.Snowflake, voiceState discord.ModifyUserVoiceState) error {
	return DoEmptyRequest(c, "PATCH", GuildsEndpoint+"/"+guildID.String()+"/voice-states/"+userID.String(), voiceState, nil, 1)
}
