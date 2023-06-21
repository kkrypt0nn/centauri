package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

const (
	InteractionsEndpoint = Endpoint + "interactions"
)

func InteractionResponse(interactionID discord.Snowflake, interactionToken string) string {
	return fmt.Sprintf("%s/%d/%s/callback", InteractionsEndpoint, interactionID, interactionToken)
}

func OriginalInteractionResponse(applicationID discord.Snowflake, interactionToken string) string {
	return fmt.Sprintf("%s/%d/%s/messages/@original", WebhooksEndpoint, applicationID, interactionToken)
}

func CreateFollowupMessage(applicationID discord.Snowflake, interactionToken string) string {
	return fmt.Sprintf("%s/%d/%s", WebhooksEndpoint, applicationID, interactionToken)
}

func FollowupMessage(applicationID discord.Snowflake, interactionToken string, messageID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d/%s/messages/%d", WebhooksEndpoint, applicationID, interactionToken, messageID)
}
