package main

import (
	"fmt"
	"github.com/kkrypt0nn/centauri"
)

func main() {
	userClient := centauri.NewRestClient("USER_TOKEN", false)
	userClient.Debug = true
	user, err := userClient.GetSelfUser()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Got User:", user.Username+"#"+user.Discriminator)
	}

	botClient := centauri.NewRestClient("BOT_TOKEN", true)
	botClient.Debug = true
	bot, err := botClient.GetSelfUser()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Got Bot:", bot.Username+"#"+bot.Discriminator)
	}
}
