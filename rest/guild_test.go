package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/ptr"
	"testing"
)

func TestGuild(t *testing.T) {
	guild, err := testingRestClient.CreateGuild(discord.CreateGuild{
		Name: "Testing Guild",
	})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = testingRestClient.ModifyGuild(guild.ID, discord.ModifyGuild{
		Name: ptr.New("New Testing Guild"),
	})
	if err != nil {
		t.Error(err)
		return
	}

	guild, err = testingRestClient.GetGuild(guild.ID, false)
	if err != nil {
		t.Error(err)
		return
	}
	expected := "New Testing Guild"
	if guild.Name != expected {
		t.Errorf("Got invalid guild name (expected: %s, got: %s)", expected, guild.Name)
		return
	}

	err = testingRestClient.DeleteGuild(guild.ID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully created, modified, got and deleted a guild")
}

func TestGuildMember(t *testing.T) {
	memberByID, err := testingRestClient.GetGuildMember(testingGuildID, testingUserID)
	if err != nil {
		t.Error(err)
		return
	}

	memberSearch, err := testingRestClient.SearchGuildMember(testingGuildID, "*-*", 5)
	if err != nil {
		t.Error(err)
		return
	}
	if len(memberSearch) != 1 {
		t.Errorf("Got invalid number of members (expected: 1, got: %d)", len(memberSearch))
		return
	}

	if memberByID.EffectiveName() != memberSearch[0].EffectiveName() {
		t.Errorf("Got two different member effective names (memberByID: %s, memberSearch: %s)", memberByID.EffectiveName(), memberSearch[0].EffectiveName())
		return
	}
	t.Log("Successfully got and searched a guild member")
}

func TestGuildBan(t *testing.T) {
	// Thank you Nyvil for being the test-subject ^-^
	var userID discord.Snowflake = 480294716764848128

	err := testingRestClient.CreateGuildBan(testingGuildID, discord.CreateGuildBan{
		UserID: userID,
	})
	if err != nil {
		t.Error(err)
		return
	}

	guildBan, err := testingRestClient.GetGuildBan(testingGuildID, userID)
	if err != nil {
		t.Error(err)
		return
	}
	expected := "nyvil"
	if guildBan.User.Username != expected {
		t.Errorf("Got invalid guild ban username (expected: %s, got: %s)", expected, guildBan.User.Username)
		return
	}

	err = testingRestClient.RemoveGuildBan(testingGuildID, userID, "No longer needed")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully banned, got and unbanned a user from a guild")
}

func TestGuildRole(t *testing.T) {
	role, err := testingRestClient.CreateGuildRole(testingGuildID, discord.CreateGuildRole{
		Name:        ptr.New("#BEBEFE is a good color"),
		Permissions: ptr.New(discord.PermissionsSendMessages | discord.PermissionsManageMessages),
		Color:       ptr.New(12500734),
		Hoist:       ptr.New(true),
	})
	if err != nil {
		t.Error(err)
		return
	}

	role, err = testingRestClient.ModifyGuildRole(testingGuildID, role.ID, discord.ModifyGuildRole{
		Name: ptr.New("#BEBEFE is the best color"),
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = testingRestClient.AddGuildMemberRole(testingGuildID, testingUserID, role.ID, "Best color")
	if err != nil {
		t.Error(err)
		return
	}

	err = testingRestClient.DeleteGuildRole(testingGuildID, role.ID, "No longer needed")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully created, modified, added to a member and deleted a guild role")
}
