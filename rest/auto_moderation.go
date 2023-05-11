package rest

import "github.com/kkrypt0nn/centauri/discord"

// GetAutoModerationRules returns a list of discord.AutoModerationRule that are available in a guild
func (c *Client) GetAutoModerationRules(guildID string) ([]discord.AutoModerationRule, error) {
	return DoRequestAsList[discord.AutoModerationRule](c, "GET", GuildsEndpoint+"/"+guildID+"/auto-moderation/rules", nil, 1)
}

// GetAutoModerationRule returns a discord.AutoModerationRule given its ID
func (c *Client) GetAutoModerationRule(guildID, ruleID string) (*discord.AutoModerationRule, error) {
	return DoRequestAs[discord.AutoModerationRule](c, "GET", GuildsEndpoint+"/"+guildID+"/auto-moderation/rules/"+ruleID, nil, 1)
}
