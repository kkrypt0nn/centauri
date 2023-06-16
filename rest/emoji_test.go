package rest

import (
	"encoding/base64"
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/ptr"
	"os"
	"testing"
)

func TestEmoji(t *testing.T) {
	bytes, err := os.ReadFile("../documentation/static/img/centauri.png")
	if err != nil {
		t.Error(err)
		return
	}

	base64encoded := "data:image/png;base64," + base64.StdEncoding.EncodeToString(bytes)

	emoji, err := testingRestClient.CreateGuildEmoji(testingGuildID, discord.CreateGuildEmoji{
		Name:  "centauri",
		Image: base64encoded,
	})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = testingRestClient.ModifyGuildEmoji(testingGuildID, emoji.ID, discord.ModifyGuildEmoji{
		Name: ptr.New("centauri_logo"),
	})
	if err != nil {
		t.Error(err)
		return
	}

	emoji, err = testingRestClient.GetGuildEmoji(testingGuildID, emoji.ID)
	if err != nil {
		t.Error(err)
		return
	}
	expected := "centauri_logo"
	if emoji.Name != expected {
		t.Errorf("Got invalid emoji name (expected: %s, got: %s)", expected, emoji.Name)
		return
	}

	err = testingRestClient.DeleteGuildEmoji(testingGuildID, emoji.ID, "No longer needed")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully created, updated, got and deleted a guild emoji")
}
