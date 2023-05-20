package main

import (
	"fmt"
	"github.com/kkrypt0nn/centauri"
	"github.com/kkrypt0nn/centauri/discord"
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
		fmt.Printf("Successfully created the rule '%s'", rule.Name)
	}
}
