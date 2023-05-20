package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	StageInstancesEndpoint = Endpoint + "stage-instances"
)

// GetStageInstance returns a stage instance structure (discord.StageInstance) associated with the stage channel
func (c *Client) GetStageInstance(channelID string) (*discord.StageInstance, error) {
	return DoRequestAsStructure[discord.StageInstance](c, "GET", StageInstancesEndpoint+"/"+channelID, nil, nil, 1)
}
