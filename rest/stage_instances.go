package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
)

// CreateStageInstance creates a new stage instance (discord.StageInstance) and returns its structure
func (c *Client) CreateStageInstance(stageInstance discord.CreateStageInstance) (*discord.StageInstance, error) {
	return DoRequestAsStructure[discord.StageInstance](c, "POST", endpoints.StageInstances(), stageInstance, nil, 1, WithReason(stageInstance.AuditLogReason))
}

// GetStageInstance returns a stage instance structure (discord.StageInstance) associated with the stage channel
func (c *Client) GetStageInstance(channelID discord.Snowflake) (*discord.StageInstance, error) {
	return DoRequestAsStructure[discord.StageInstance](c, "GET", endpoints.StageInstance(channelID), nil, nil, 1)
}

// ModifyStageInstance modifies an existing stage instance (discord.StageInstance) for the given channel ID and returns its new structure
func (c *Client) ModifyStageInstance(channelID discord.Snowflake, stageInstance discord.ModifyStageInstance) (*discord.StageInstance, error) {
	return DoRequestAsStructure[discord.StageInstance](c, "PATCH", endpoints.StageInstance(channelID), stageInstance, nil, 1, WithReason(stageInstance.AuditLogReason))
}

// DeleteStageInstance deletes an existing stage instance (discord.StageInstance) for the given channel ID
func (c *Client) DeleteStageInstance(channelID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", endpoints.StageInstance(channelID), nil, nil, 1)
}
