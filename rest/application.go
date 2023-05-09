package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	ApplicationsEndpoint = Endpoint + "applications"
)

func (c *Client) GetApplicationRoleConnectionMetadata(appID string) (*discord.ApplicationRoleConnectionMetadata, error) {
	return DoRequestAs[discord.ApplicationRoleConnectionMetadata](c, "GET", ApplicationsEndpoint+"/"+appID+"/role-connections/metadata", nil, 1)
}
