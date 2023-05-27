package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
)

// GetAutoModerationRules returns a list of auto moderation rule structures (discord.AutoModerationRule) for the given guild ID
func (c *Client) GetAutoModerationRules(guildID string) ([]discord.AutoModerationRule, error) {
	return DoRequestAsList[discord.AutoModerationRule](c, "GET", GuildsEndpoint+"/"+guildID+"/auto-moderation/rules", nil, nil, 1)
}

// GetAutoModerationRule returns an auto moderation rule structure (discord.AutoModerationRule) for the given guild and rule ID
func (c *Client) GetAutoModerationRule(guildID, ruleID string) (*discord.AutoModerationRule, error) {
	return DoRequestAsStructure[discord.AutoModerationRule](c, "GET", GuildsEndpoint+"/"+guildID+"/auto-moderation/rules/"+ruleID, nil, nil, 1)
}

// CreateAutoModerationRule creates a new auto moderation rule (discord.AutoModerationRule) and returns its structure
func (c *Client) CreateAutoModerationRule(guildID string, autoModerationRule discord.CreateAutoModerationRule) (*discord.AutoModerationRule, error) {
	return DoRequestAsStructure[discord.AutoModerationRule](c, "POST", GuildsEndpoint+"/"+guildID+"/auto-moderation/rules", autoModerationRule, nil, 1, WithReason(autoModerationRule.AuditLogReason))
}

// ModifyAutoModerationRule modifies an existing auto moderation rule (discord.AutoModerationRule) and returns its structure
func (c *Client) ModifyAutoModerationRule(guildID, ruleID string, autoModerationRule discord.ModifyAutoModerationRule) (*discord.AutoModerationRule, error) {
	return DoRequestAsStructure[discord.AutoModerationRule](c, "PATCH", GuildsEndpoint+"/"+guildID+"/auto-moderation/rules/"+ruleID, autoModerationRule, nil, 1, WithReason(autoModerationRule.AuditLogReason))
}

// DeleteAutoModerationRule deletes an existing auto moderation rule (discord.AutoModerationRule)
func (c *Client) DeleteAutoModerationRule(guildID, ruleID string) error {
	_, _, err := c.DoRequest("DELETE", GuildsEndpoint+"/"+guildID+"/auto-moderation/rules/"+ruleID, nil, nil, 1)
	return err
}
