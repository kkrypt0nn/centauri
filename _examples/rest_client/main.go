package main

import (
	"fmt"
	"github.com/kkrypt0nn/centauri"
)

func main() {
	botClient := centauri.NewRestClient("Bot BOT_TOKEN")
	botClient.Debug = true
	bot, err := botClient.GetCurrentUser()
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(fmt.Sprintf("Got bot %s with ID %s", bot.Username, bot.ID))
	}
}
