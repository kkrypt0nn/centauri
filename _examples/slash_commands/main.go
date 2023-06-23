package main

import (
	"fmt"
	"github.com/kkrypt0nn/centauri"
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/gateway"
	"github.com/kkrypt0nn/centauri/ptr"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	botClient := centauri.NewGatewayClient("Bot BOT_TOKEN", 0)

	botClient.On(gateway.EventTypeReady, func(c *gateway.Client, ready *gateway.Ready) {
		botClient.Logger.Info(fmt.Sprintf("We are now logged in as %s", ready.User.Username))
	})
	botClient.On(gateway.EventTypeInteractionCreate, func(c *gateway.Client, interaction *gateway.InteractionCreate) {
		if interaction.Interaction.Data.InteractionType() == discord.InteractionTypeApplicationCommand {
			if interaction.Interaction.Data.(*discord.InteractionApplicationCommandData).Name == "test" {
				err := botClient.Rest().CreateInteractionResponse(interaction.ID, interaction.Token, discord.CreateInteractionResponse{
					Type: discord.InteractionResponseTypeMessage,
					Data: discord.MessageInteractionResponse{
						Content: ptr.New("Hello world!"),
					},
				})
				if err != nil {
					c.Logger.Error(err.Error())
				}
			}
		}
	})

	_, err := botClient.Rest().CreateGlobalApplicationCommand(discord.CreateGlobalApplicationCommand{
		Name:        "test",
		Description: ptr.New("test"),
	})
	if err != nil {
		botClient.Logger.Error(err.Error())
	}

	err = botClient.Login()
	if err != nil {
		botClient.Logger.Error(err.Error())
		botClient.Close()
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	botClient.Close()
}
