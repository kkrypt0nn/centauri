package rest

import "github.com/kkrypt0nn/centauri/discord"

const (
	VoiceEndpoint = Endpoint + "voice"
)

// ListVoiceRegions returns a list of voice region structures (discord.VoiceRegion) that can be used when setting a voice or stage channel's `rtc_region`
func (c *Client) ListVoiceRegions() ([]discord.VoiceRegion, error) {
	return DoRequestAsList[discord.VoiceRegion](c, "GET", VoiceEndpoint+"/regions", nil, nil, 1)
}
