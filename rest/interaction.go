package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/endpoints"
)

// CreateInteractionResponse creates a response to an interaction (discord.Interaction) for the given interaction ID and token
func (c *Client) CreateInteractionResponse(interactionID discord.Snowflake, interactionToken string, interactionResponse discord.CreateInteractionResponse) error {
	if interactionResponse.Data != nil && interactionResponse.Data.Type() == discord.InteractionResponseTypeMessage {
		if len(interactionResponse.Data.(discord.MessageInteractionResponse).Files) >= 1 {
			contentType, body, err := CreateMultipartBodyWithJSON(interactionResponse, interactionResponse.Data.(discord.MessageInteractionResponse).Files)
			if err != nil {
				return err
			}
			return DoEmptyRequestWithFiles(c, "POST", endpoints.InteractionResponse(interactionID, interactionToken), body, nil, 1, WithHeader("Content-Type", contentType))
		} else {
			return DoEmptyRequest(c, "POST", endpoints.InteractionResponse(interactionID, interactionToken), interactionResponse, nil, 1)
		}
	}
	return DoEmptyRequest(c, "POST", endpoints.InteractionResponse(interactionID, interactionToken), interactionResponse, nil, 1)
}

// GetOriginalInteractionResponse returns the message structure (discord.Message) of the original interaction response for given the interaction token
func (c *Client) GetOriginalInteractionResponse(interactionToken string, threadID discord.Snowflake) (*discord.Message, error) {
	queryParams := make(QueryParameters)
	if threadID != 0 {
		queryParams["thread_id"] = threadID.String()
	}
	return DoRequestAsStructure[discord.Message](c, "GET", endpoints.OriginalInteractionResponse(c.selfUser.ID, interactionToken), nil, queryParams, 1)
}

// EditOriginalInteractionResponse edits the original interaction response message (discord.Message) for the given interaction token and returns its new structure
func (c *Client) EditOriginalInteractionResponse(interactionToken string, threadID discord.Snowflake, message discord.EditOriginalInteractionResponse) (*discord.Webhook, error) {
	queryParams := make(QueryParameters)
	if threadID != 0 {
		queryParams["thread_id"] = threadID.String()
	}
	if len(message.Files) >= 1 {
		contentType, body, err := CreateMultipartBodyWithJSON(message, message.Files)
		if err != nil {
			return nil, err
		}
		return DoRequestAsStructure[discord.Webhook](c, "PATCH", endpoints.OriginalInteractionResponse(c.selfUser.ID, interactionToken), body, queryParams, 1, WithHeader("Content-Type", contentType))
	} else {
		return DoRequestAsStructure[discord.Webhook](c, "PATCH", endpoints.OriginalInteractionResponse(c.selfUser.ID, interactionToken), message, queryParams, 1)
	}
}

// DeleteOriginalInteractionResponse deletes the original interaction response message (discord.Message) for the given interaction token
func (c *Client) DeleteOriginalInteractionResponse(interactionToken string) error {
	return DoEmptyRequest(c, "DELETE", endpoints.OriginalInteractionResponse(c.selfUser.ID, interactionToken), nil, nil, 1)
}

// CreateFollowupMessage creates a followup message (discord.Message) for the given interaction token
func (c *Client) CreateFollowupMessage(interactionToken string, message discord.CreateFollowupMessage) (*discord.Message, error) {
	if len(message.Files) >= 1 {
		contentType, body, err := CreateMultipartBodyWithJSON(message, message.Files)
		if err != nil {
			return nil, err
		}
		return DoRequestWithFiles[discord.Message](c, "POST", endpoints.CreateFollowupMessage(c.selfUser.ID, interactionToken), body, nil, 1, WithHeader("Content-Type", contentType))
	} else {
		return DoRequestAsStructure[discord.Message](c, "POST", endpoints.CreateFollowupMessage(c.selfUser.ID, interactionToken), message, nil, 1)
	}
}

// GetFollowupMessage returns the message structure (discord.Message) of the followup message for the given the interaction token and message ID
func (c *Client) GetFollowupMessage(interactionToken string, messageID discord.Snowflake) (*discord.Message, error) {
	return DoRequestAsStructure[discord.Message](c, "GET", endpoints.FollowupMessage(c.selfUser.ID, interactionToken, messageID), nil, nil, 1)
}

// EditFollowupMessage edits the followup message (discord.Message) for the given interaction token and message ID and returns its new structure
func (c *Client) EditFollowupMessage(interactionToken string, messageID, threadID discord.Snowflake, message discord.EditFollowupMessage) (*discord.Webhook, error) {
	queryParams := make(QueryParameters)
	if threadID != 0 {
		queryParams["thread_id"] = threadID.String()
	}
	if len(message.Files) >= 1 {
		contentType, body, err := CreateMultipartBodyWithJSON(message, message.Files)
		if err != nil {
			return nil, err
		}
		return DoRequestAsStructure[discord.Webhook](c, "PATCH", endpoints.FollowupMessage(c.selfUser.ID, interactionToken, messageID), body, queryParams, 1, WithHeader("Content-Type", contentType))
	} else {
		return DoRequestAsStructure[discord.Webhook](c, "PATCH", endpoints.FollowupMessage(c.selfUser.ID, interactionToken, messageID), message, queryParams, 1)
	}
}

// DeleteFollowupMessage deletes the followup response message (discord.Message) for the given interaction token and message ID
func (c *Client) DeleteFollowupMessage(interactionToken string, messageID discord.Snowflake) error {
	return DoEmptyRequest(c, "DELETE", endpoints.FollowupMessage(c.selfUser.ID, interactionToken, messageID), nil, nil, 1)
}
