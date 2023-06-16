package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/ptr"
	"testing"
)

func TestGlobalApplicationCommands(t *testing.T) {
	applicationCommand, err := testingRestClient.CreateGlobalApplicationCommand(testingBotID, discord.CreateGlobalApplicationCommand{
		Name:        "test",
		Description: ptr.New("Testing global application command"),
		Options: []discord.ApplicationCommandOption{
			&discord.StringCommandOption{
				Name:        "test_string",
				Description: "String option test",
				Required:    true,
				MinLength:   ptr.New(13),
				MaxLength:   ptr.New(37),
			},
			&discord.IntegerCommandOption{
				Name:        "test_int",
				Description: "Integer option test",
				Required:    true,
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = testingRestClient.EditGlobalApplicationCommand(testingBotID, applicationCommand.ID, discord.EditGlobalApplicationCommand{Name: ptr.New("test_global")})
	if err != nil {
		t.Error(err)
		return
	}

	applicationCommand, err = testingRestClient.GetGlobalApplicationCommand(testingBotID, applicationCommand.ID)
	if err != nil {
		t.Error(err)
		return
	}
	expected := "test_global"
	if applicationCommand.Name != expected {
		t.Errorf("Got an invalid global application command name (expected: %s, got: %s)", expected, applicationCommand.Name)
		return
	}

	err = testingRestClient.DeleteGlobalApplicationCommand(testingBotID, applicationCommand.ID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully created, updated, got and deleted a global application command")
}

func TestGuildApplicationCommands(t *testing.T) {
	applicationCommand, err := testingRestClient.CreateGuildApplicationCommand(testingBotID, testingGuildID, discord.CreateGuildApplicationCommand{
		Name:        "test",
		Description: ptr.New("Testing guild application command"),
		Options: []discord.ApplicationCommandOption{
			&discord.StringCommandOption{
				Name:        "test_string",
				Description: "String option test",
				Required:    true,
				MinLength:   ptr.New(13),
				MaxLength:   ptr.New(37),
			},
			&discord.IntegerCommandOption{
				Name:        "test_int",
				Description: "Integer option test",
				Required:    true,
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = testingRestClient.EditGuildApplicationCommand(testingBotID, testingGuildID, applicationCommand.ID, discord.EditGuildApplicationCommand{Name: ptr.New("test_guild")})
	if err != nil {
		t.Error(err)
		return
	}

	applicationCommand, err = testingRestClient.GetGuildApplicationCommand(testingBotID, testingGuildID, applicationCommand.ID)
	if err != nil {
		t.Error(err)
		return
	}
	expected := "test_guild"
	if applicationCommand.Name != expected {
		t.Errorf("Got an invalid global application command name (expected: %s, got: %s)", expected, applicationCommand.Name)
		return
	}

	err = testingRestClient.DeleteGuildApplicationCommand(testingBotID, testingGuildID, applicationCommand.ID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully created, updated, got and deleted a guild application command")
}
