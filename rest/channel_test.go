package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/ptr"
	"testing"
)

func TestChannel(t *testing.T) {
	channel, err := testingRestClient.CreateGuildChannel(testingGuildID, discord.CreateGuildChannel{
		Name:  "test",
		Topic: ptr.New("A testing channel"),
	})
	if err != nil {
		t.Error(err)
		return
	}

	channel, err = testingRestClient.ModifyChannel(channel.ID, discord.ModifyChannel{
		Name: ptr.New("not-general"),
	})
	if err != nil {
		t.Error(err)
		return
	}

	channel, err = testingRestClient.GetChannel(channel.ID)
	if err != nil {
		t.Error(err)
		return
	}
	expected := "not-general"
	if channel.Name != expected {
		t.Errorf("Got invalid channel name (expected: %s, got: %s)", expected, channel.Name)
		return
	}

	err = testingRestClient.DeleteChannel(channel.ID, "No longer needed :D")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully created, modified and deleted a channel")
}

func TestMessage(t *testing.T) {
	message, err := testingRestClient.CreateMessage(testingChannelID, discord.CreateMessage{
		Content: ptr.New("Hello world!"),
		Embeds: []discord.Embed{
			{
				Description: "Hello world!",
			},
		},
		Components: []discord.Component{
			&discord.ActionRow{Components: []discord.Component{
				&discord.Button{
					Style:    discord.ButtonStylePrimary,
					Label:    "Hello world!",
					CustomID: "hello-world",
				},
			}},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = testingRestClient.EditMessage(testingChannelID, message.ID, discord.EditMessage{
		Content: ptr.New("Hello brand new world!"),
	})
	if err != nil {
		t.Error(err)
		return
	}

	message, err = testingRestClient.GetChannelMessage(testingChannelID, message.ID)
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Hello brand new world!"
	if message.Content != expected {
		t.Errorf("Got invalid message content (expected: %s, got: %s)", expected, message.Content)
		return
	}

	err = testingRestClient.CreateReaction(testingChannelID, message.ID, "🚑")
	if err != nil {
		t.Error(err)
		return
	}

	err = testingRestClient.DeleteAllReactions(testingChannelID, message.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = testingRestClient.DeleteMessage(testingChannelID, message.ID, "No longer needed")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully created, edited, got and deleted a message, as well as added a reaction and removed all reactions from that message")
}

func TestThreads(t *testing.T) {
	thread, err := testingRestClient.StartThreadWithoutMessage(testingChannelID, discord.StartThreadWithoutMessage{
		Name: "Testing thread",
	})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = testingRestClient.CreateMessage(thread.ID, discord.CreateMessage{
		Content: ptr.New("Hello world!"),
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = testingRestClient.AddThreadMember(thread.ID, testingUserID)
	if err != nil {
		t.Error(err)
		return
	}

	err = testingRestClient.RemoveThreadMember(thread.ID, testingUserID)
	if err != nil {
		t.Error(err)
		return
	}

	err = testingRestClient.DeleteChannel(thread.ID, "No longer needed")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully created a thread, sent a message in it, added and removed a user from it and deleted the thread")
}
