package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

const (
	StageInstancesEndpoint = Endpoint + "stage-instances"
)

func StageInstances() string {
	return StageInstancesEndpoint
}

func StageInstance(channelID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d", StageInstancesEndpoint, channelID)
}
