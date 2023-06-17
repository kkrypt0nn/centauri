package main

import (
	"fmt"
	"github.com/kkrypt0nn/centauri"
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/gateway"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	botClient := centauri.NewGatewayClient(os.Getenv("TOKEN"), discord.IntentsGuildMessages|discord.IntentsMessageContent)

	botClient.On(gateway.EventTypeReady, func(c *gateway.Client, ready *gateway.Ready) {
		botClient.Logger.Info(fmt.Sprintf("We are now logged in as %s", ready.User.Username))
	})
	botClient.On(gateway.EventTypeMessageCreate, func(c *gateway.Client, message *gateway.MessageCreate) {
		botClient.Logger.Info(fmt.Sprintf("Got a new message from %s: %s", message.Author.Username, message.Content))
	})

	err := botClient.Login()
	if err != nil {
		log.Fatal(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	botClient.Close()
}
