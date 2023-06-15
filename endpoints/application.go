package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

const (
	ApplicationsEndpoint = Endpoint + "applications"
)

func ApplicationRoleConnectionMetadataRecords(applicationID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/role-connections/metadata", ApplicationsEndpoint, applicationID)
}
