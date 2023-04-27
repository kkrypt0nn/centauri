package main

import (
	"fmt"
	"github.com/kkrypt0nn/centauri"
)

func main() {
	userClient := centauri.NewRestClient("USER_TOKEN")
	userClient.Debug = true
	user, err := userClient.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Got User:", user.Username+"#"+user.Discriminator)
	}

	botClient := centauri.NewRestClient("Bot BOT_TOKEN")
	botClient.Debug = true
	bot, err := botClient.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Got Bot:", bot.Username+"#"+bot.Discriminator)
	}
}
