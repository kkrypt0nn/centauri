package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	ApplicationsEndpoint = Endpoint + "applications"
)

// GetApplicationRoleConnectionMetadataRecords returns a list of application role connection metadata structures (discord.ApplicationRoleConnectionMetadata) for the given application ID
func (c *Client) GetApplicationRoleConnectionMetadataRecords(applicationID string) (*discord.ApplicationRoleConnectionMetadata, error) {
	return DoRequestAsStructure[discord.ApplicationRoleConnectionMetadata](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/role-connections/metadata", nil, nil, 1)
}

// UpdateApplicationRoleConnectionMetadataRecords returns a list of application role connection metadata structures (discord.ApplicationRoleConnectionMetadata) for the given application ID
func (c *Client) UpdateApplicationRoleConnectionMetadataRecords(applicationID string, records []discord.ApplicationRoleConnectionMetadata) ([]discord.ApplicationRoleConnectionMetadata, error) {
	return DoRequestAsList[discord.ApplicationRoleConnectionMetadata](c, "PUT", ApplicationsEndpoint+"/"+applicationID+"/role-connections/metadata", records, nil, 1)
}
