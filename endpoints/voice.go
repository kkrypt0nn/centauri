package endpoints

import "fmt"

const (
	VoiceEndpoint = Endpoint + "voice"
)

func VoiceRegions() string {
	return fmt.Sprintf("%s/regions", VoiceEndpoint)
}
