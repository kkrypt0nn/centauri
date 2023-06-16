package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/ptr"
	"testing"
)

func TestCurrentUser(t *testing.T) {
	user, err := testingRestClient.GetCurrentUser()
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Centauri Testing"
	if user.Username != expected {
		t.Errorf("Got invalid current user username (expected: %s, got: %s)", expected, user.Username)
		return
	}

	guilds, err := testingRestClient.GetCurrentUserGuilds(0, 0, 5)
	if err != nil {
		t.Error(err)
		return
	}
	expectedLengthMin := 1
	expectedLengthMax := 2 // In case it gets ran while the bot created a new guild
	if len(guilds) < expectedLengthMin || len(guilds) > expectedLengthMax {
		t.Errorf("Got invalid length of guilds (expected: %d<=len(guilds)<=%d, got: %d)", expectedLengthMin, expectedLengthMax, len(guilds))
		return
	}

	t.Log("Successfully got the current user and got its guilds")
}

func TestUser(t *testing.T) {
	user, err := testingRestClient.GetUser(testingUserID)
	if err != nil {
		t.Error(err)
		return
	}
	expected := "*-*"
	if user.GlobalName != expected {
		t.Errorf("Got invalid global name (expected: %s, got: %s)", expected, user.GlobalName)
		return
	}

	dmChannel, err := testingRestClient.CreateDM(user.ID)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = testingRestClient.CreateMessage(dmChannel.ID, discord.CreateMessage{
		Content: ptr.New("Centauri Testing"),
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully got a user and send a private message")
}
