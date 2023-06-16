package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/ptr"
	"testing"
)

func TestAutoModerationRule(t *testing.T) {
	rule, err := testingRestClient.CreateAutoModerationRule(testingGuildID, discord.CreateAutoModerationRule{
		Name:        "No Java",
		EventType:   discord.AutoModerationEventTypeMessageSend,
		TriggerType: discord.AutoModerationTriggerTypeKeyword,
		TriggerMetadata: &discord.AutoModerationTriggerMetadata{
			KeywordFilter: []string{"*java*"},
		},
		Actions: []discord.AutoModerationAction{{
			Type: discord.AutoModerationActionTypeBlockMessage,
			Metadata: &discord.AutoModerationActionMetadata{
				CustomMessage: "Java is not allowed here!",
			},
		}},
		AuditLogReason: "It should not be talked about Java here!",
	})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = testingRestClient.ModifyAutoModerationRule(testingGuildID, rule.ID, discord.ModifyAutoModerationRule{
		Name:           ptr.New("Java Not Allowed"),
		AuditLogReason: "Wrong name given",
	})
	if err != nil {
		t.Error(err)
		return
	}

	rule, err = testingRestClient.GetAutoModerationRule(testingGuildID, rule.ID)
	if err != nil {
		t.Error(err)
		return
	}
	expected := "Java Not Allowed"
	if rule.Name != expected {
		t.Errorf("Got invalid auto moderation rule name (expected: %s, got: %s)", expected, rule.Name)
		return
	}

	err = testingRestClient.DeleteAutoModerationRule(rule.GuildID, rule.ID, "Rule no longer needed")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Successfully created, edited, got and deleted an auto moderation rule")
}
