package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/ptr"
	"testing"
)

func TestWebhook(t *testing.T) {
	webhook, err := testingRestClient.CreateWebhook(testingChannelID, discord.CreateWebhook{
		Name: "Centauri Testing",
	})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = testingRestClient.ModifyWebhook(webhook.ID, discord.ModifyWebhook{
		Name: ptr.New("Centauri Testing Webhook"),
	})
	if err != nil {
		t.Error(err)
		return
	}

	webhook, err = testingRestClient.GetWebhook(webhook.ID)
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Centauri Testing Webhook"
	if webhook.Name != expected {
		t.Errorf("Got invalid webhook name (expected: %s, got: %s)", expected, webhook.Name)
		return
	}

	message, err := testingRestClient.ExecuteWebhook(webhook.ID, webhook.Token, 0, true, discord.ExecuteWebhook{
		Content: ptr.New("Just testing around.."),
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = testingRestClient.DeleteMessage(webhook.ChannelID, message.ID, "No longer needed")
	if err != nil {
		t.Error(err)
		return
	}

	err = testingRestClient.DeleteWebhook(webhook.ID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully created, modified, executed and deleted a webhook")
}
