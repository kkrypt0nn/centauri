// Package endpoints is a package that will facilitate getting Discord endpoints
package endpoints

const (
	BaseURL  = "https://discord.com/api/v"
	Version  = "10"
	Endpoint = BaseURL + Version + "/"

	GatewayBaseURL = "wss://gateway.discord.gg/?v="
	GatewayVersion = "10"
	GatewayURL     = GatewayBaseURL + GatewayVersion + "&encoding=json"
)
