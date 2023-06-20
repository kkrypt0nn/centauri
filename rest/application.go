package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
)

// GetApplicationRoleConnectionMetadataRecords returns a list of application role connection metadata structures (discord.ApplicationRoleConnectionMetadata) for the given application ID
func (c *Client) GetApplicationRoleConnectionMetadataRecords() ([]discord.ApplicationRoleConnectionMetadata, error) {
	return DoRequestAsList[discord.ApplicationRoleConnectionMetadata](c, "GET", endpoints.ApplicationRoleConnectionMetadataRecords(c.selfUser.ID), nil, nil, 1)
}

// UpdateApplicationRoleConnectionMetadataRecords returns a list of application role connection metadata structures (discord.ApplicationRoleConnectionMetadata) for the given application ID
func (c *Client) UpdateApplicationRoleConnectionMetadataRecords(records []discord.ApplicationRoleConnectionMetadata) ([]discord.ApplicationRoleConnectionMetadata, error) {
	return DoRequestAsList[discord.ApplicationRoleConnectionMetadata](c, "PUT", endpoints.ApplicationRoleConnectionMetadataRecords(c.selfUser.ID), records, nil, 1)
}
