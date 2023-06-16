package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"testing"
)

func TestApplicationRoleConnectionMetadataRecords(t *testing.T) {
	_, err := testingRestClient.UpdateApplicationRoleConnectionMetadataRecords(testingBotID, []discord.ApplicationRoleConnectionMetadata{
		{
			Type:        discord.ApplicationRoleConnectionMetadataTypeIntegerEqual,
			Key:         "test1",
			Name:        "test1",
			Description: "This is a first test",
		}, {
			Type:        discord.ApplicationRoleConnectionMetadataTypeBooleanNotEqual,
			Key:         "test2",
			Name:        "test2",
			Description: "This is a second test",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	records, err := testingRestClient.GetApplicationRoleConnectionMetadataRecords(testingBotID)
	if err != nil {
		t.Error(err)
		return
	}
	if len(records) != 2 {
		t.Errorf("Got invalid number of application role connection metadata records (expected: 2, got: %d)", len(records))
		return
	}

	t.Log("Successfully updated and got the new application role connection metadata records")

	_, _ = testingRestClient.UpdateApplicationRoleConnectionMetadataRecords(testingBotID, []discord.ApplicationRoleConnectionMetadata{})
}
