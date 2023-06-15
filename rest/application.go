package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
)

// GetApplicationRoleConnectionMetadataRecords returns a list of application role connection metadata structures (discord.ApplicationRoleConnectionMetadata) for the given application ID
func (c *Client) GetApplicationRoleConnectionMetadataRecords(applicationID discord.Snowflake) (*discord.ApplicationRoleConnectionMetadata, error) {
	return DoRequestAsStructure[discord.ApplicationRoleConnectionMetadata](c, "GET", endpoints.ApplicationRoleConnectionMetadataRecords(applicationID), nil, nil, 1)
}

// UpdateApplicationRoleConnectionMetadataRecords returns a list of application role connection metadata structures (discord.ApplicationRoleConnectionMetadata) for the given application ID
func (c *Client) UpdateApplicationRoleConnectionMetadataRecords(applicationID discord.Snowflake, records []discord.ApplicationRoleConnectionMetadata) ([]discord.ApplicationRoleConnectionMetadata, error) {
	return DoRequestAsList[discord.ApplicationRoleConnectionMetadata](c, "PUT", endpoints.ApplicationRoleConnectionMetadataRecords(applicationID), records, nil, 1)
}
