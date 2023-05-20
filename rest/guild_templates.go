package rest

import "github.com/kkrypt0nn/centauri/discord"

// GetGuildTemplate returns a guild template structure (discord.GuildTemplate) for the given template code
func (c *Client) GetGuildTemplate(templateCode string) (*discord.GuildTemplate, error) {
	return DoRequestAsStructure[discord.GuildTemplate](c, "GET", GuildsEndpoint+"/templates/"+templateCode, nil, nil, 1)
}

// GetGuildTemplates returns a list of guild template structures (discord.GuildTemplate) for the given guild ID
func (c *Client) GetGuildTemplates(guildID string) ([]discord.GuildTemplate, error) {
	return DoRequestAsList[discord.GuildTemplate](c, "GET", GuildsEndpoint+"/"+guildID+"/templates", nil, nil, 1)
}
