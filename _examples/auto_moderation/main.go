package main

import (
	"fmt"
	"github.com/kkrypt0nn/centauri"
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/ptr"
	"time"
)

func main() {
	botClient := centauri.NewRestClient("Bot BOT_TOKEN")
	botClient.Debug = true
	rule, err := botClient.CreateAutoModerationRule("GUILD_ID", discord.CreateAutoModerationRule{
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
		fmt.Println("err:", err)
	} else {
		fmt.Printf("Successfully created the rule '%s'\n", rule.Name)
	}

	time.Sleep(5 * time.Second)

	rule, err = botClient.ModifyAutoModerationRule(rule.GuildID, rule.ID, discord.ModifyAutoModerationRule{
		Name:           ptr.New("Java Not Allowed"),
		AuditLogReason: "Wrong name given",
	})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("Successfully updated the rule name to '%s'\n", rule.Name)
	}

	time.Sleep(5 * time.Second)

	err = botClient.DeleteAutoModerationRule(rule.GuildID, rule.ID, "Rule no longer needed")
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("Successfully deleted the rule '%s'", rule.Name)
	}
}
