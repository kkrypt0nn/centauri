package main

import (
	"fmt"
	"net/http"

	"github.com/kkrypt0nn/centauri"
	"github.com/kkrypt0nn/centauri/rest"
)

type MyInterceptor struct {
	rest.DefaultInterceptor
}

func NewCustomInterceptor() *MyInterceptor {
	return &MyInterceptor{}
}

func (i *MyInterceptor) ModifyRequest(r *http.Request) {
	r.Header.Set("Authorization", "Bot OTHER_BOT_TOKEN")
}

func main() {
	botClient := centauri.NewRestClient("Bot BOT_TOKEN")
	botClient.Debug = true

	botClient.HttpClient.SetInterceptor(NewCustomInterceptor())

	bot, err := botClient.GetCurrentUser()
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(fmt.Sprintf("Got bot %s with ID %s", bot.Username, bot.ID))
	}
}
