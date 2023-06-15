package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	StageInstancesEndpoint = Endpoint + "stage-instances"
)

// CreateStageInstance creates a new stage instance (discord.StageInstance) and returns its structure
func (c *Client) CreateStageInstance(stageInstance discord.CreateStageInstance) (*discord.StageInstance, error) {
	return DoRequestAsStructure[discord.StageInstance](c, "POST", StageInstancesEndpoint, stageInstance, nil, 1, WithReason(stageInstance.AuditLogReason))
}

// GetStageInstance returns a stage instance structure (discord.StageInstance) associated with the stage channel
func (c *Client) GetStageInstance(channelID discord.Snowflake) (*discord.StageInstance, error) {
	return DoRequestAsStructure[discord.StageInstance](c, "GET", StageInstancesEndpoint+"/"+channelID.String(), nil, nil, 1)
}

// ModifyStageInstance modifies an existing stage instance (discord.StageInstance) for the given channel ID and returns its new structure
func (c *Client) ModifyStageInstance(channelID discord.Snowflake, stageInstance discord.ModifyStageInstance) (*discord.StageInstance, error) {
	return DoRequestAsStructure[discord.StageInstance](c, "PATCH", StageInstancesEndpoint+"/"+channelID.String(), stageInstance, nil, 1, WithReason(stageInstance.AuditLogReason))
}

// DeleteStageInstance deletes an existing stage instance (discord.StageInstance) for the given channel ID
func (c *Client) DeleteStageInstance(channelID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", StageInstancesEndpoint+"/"+channelID.String(), nil, nil, 1)
}
