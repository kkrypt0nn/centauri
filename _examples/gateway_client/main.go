package main

import (
	"fmt"
	"github.com/kkrypt0nn/centauri"
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/gateway"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	botClient := centauri.NewGatewayClient("Bot BOT_TOKEN", discord.IntentsGuildMessages|discord.IntentsMessageContent)

	botClient.On(gateway.EventTypeReady, func(c *gateway.Client, ready *gateway.Ready) {
		botClient.Logger.Info(fmt.Sprintf("We are now logged in as %s", ready.User.Username))
	})
	botClient.On(gateway.EventTypeMessageCreate, func(c *gateway.Client, message *gateway.MessageCreate) {
		botClient.Logger.Info(fmt.Sprintf("Got a new message from %s: %s", message.Author.Username, message.Content))
		if message.Content == "updateActivity" {
			err := botClient.SetActivity(discord.StatusTypeDND, discord.ActivityTypeWatching, "the series Dark")
			if err != nil {
				botClient.Logger.Error(err.Error())
			}
		}
	})

	err := botClient.Login()
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
