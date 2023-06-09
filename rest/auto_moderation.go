package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
)

// GetAutoModerationRules returns a list of auto moderation rule structures (discord.AutoModerationRule) for the given guild ID
func (c *Client) GetAutoModerationRules(guildID discord.Snowflake) ([]discord.AutoModerationRule, error) {
	return DoRequestAsList[discord.AutoModerationRule](c, "GET", endpoints.AutoModerationRules(guildID), nil, nil, 1)
}

// GetAutoModerationRule returns an auto moderation rule structure (discord.AutoModerationRule) for the given guild and rule ID
func (c *Client) GetAutoModerationRule(guildID, ruleID discord.Snowflake) (*discord.AutoModerationRule, error) {
	return DoRequestAsStructure[discord.AutoModerationRule](c, "GET", endpoints.AutoModerationRule(guildID, ruleID), nil, nil, 1)
}

// CreateAutoModerationRule creates a new auto moderation rule (discord.AutoModerationRule) and returns its structure
func (c *Client) CreateAutoModerationRule(guildID discord.Snowflake, autoModerationRule discord.CreateAutoModerationRule) (*discord.AutoModerationRule, error) {
	return DoRequestAsStructure[discord.AutoModerationRule](c, "POST", endpoints.AutoModerationRules(guildID), autoModerationRule, nil, 1, WithReason(autoModerationRule.AuditLogReason))
}

// ModifyAutoModerationRule modifies an existing auto moderation rule (discord.AutoModerationRule) and returns its structure
func (c *Client) ModifyAutoModerationRule(guildID, ruleID discord.Snowflake, autoModerationRule discord.ModifyAutoModerationRule) (*discord.AutoModerationRule, error) {
	return DoRequestAsStructure[discord.AutoModerationRule](c, "PATCH", endpoints.AutoModerationRule(guildID, ruleID), autoModerationRule, nil, 1, WithReason(autoModerationRule.AuditLogReason))
}

// DeleteAutoModerationRule deletes an existing auto moderation rule (discord.AutoModerationRule)
func (c *Client) DeleteAutoModerationRule(guildID, ruleID discord.Snowflake, reason string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.AutoModerationRule(guildID, ruleID), nil, nil, 1, WithReason(reason))
}
