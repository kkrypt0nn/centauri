package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
)

// ListVoiceRegions returns a list of voice region structures (discord.VoiceRegion) that can be used when setting a voice or stage channel's `rtc_region`
func (c *Client) ListVoiceRegions() ([]discord.VoiceRegion, error) {
	return DoRequestAsList[discord.VoiceRegion](c, "GET", endpoints.VoiceRegions(), nil, nil, 1)
}
