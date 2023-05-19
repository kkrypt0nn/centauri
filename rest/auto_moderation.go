package rest

import "github.com/kkrypt0nn/centauri/discord"

// GetAutoModerationRules returns a list of auto moderation rule structures (discord.AutoModerationRule) for the given guild ID
func (c *Client) GetAutoModerationRules(guildID string) ([]discord.AutoModerationRule, error) {
	return DoRequestAsList[discord.AutoModerationRule](c, "GET", GuildsEndpoint+"/"+guildID+"/auto-moderation/rules", nil, 1)
}

// GetAutoModerationRule returns a auto moderation rule structure (discord.AutoModerationRule) for the given guild and rule ID
func (c *Client) GetAutoModerationRule(guildID, ruleID string) (*discord.AutoModerationRule, error) {
	return DoRequestAs[discord.AutoModerationRule](c, "GET", GuildsEndpoint+"/"+guildID+"/auto-moderation/rules/"+ruleID, nil, 1)
}
