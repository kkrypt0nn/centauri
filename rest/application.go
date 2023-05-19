package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	ApplicationsEndpoint = Endpoint + "applications"
)

// GetApplicationRoleConnectionMetadata returns a list of application role connection metadata structures (discord.ApplicationRoleConnectionMetadata) for the given application ID
func (c *Client) GetApplicationRoleConnectionMetadata(applicationID string) (*discord.ApplicationRoleConnectionMetadata, error) {
	return DoRequestAs[discord.ApplicationRoleConnectionMetadata](c, "GET", ApplicationsEndpoint+"/"+applicationID+"/role-connections/metadata", nil, 1)
}
