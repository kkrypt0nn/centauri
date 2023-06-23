package main

import (
	"github.com/kkrypt0nn/centauri"
	"github.com/kkrypt0nn/centauri/discord"
	"os"
)

func main() {
	botClient := centauri.NewRestClient("Bot BOT_TOKEN")

	logoFile, err := os.Open("./documentation/static/img/centauri.png")
	if err != nil {
		panic(err)
	}

	// Normal message
	_, err = botClient.CreateMessage(123456789, discord.CreateMessage{
		Files: []discord.File{
			{
				Name:   "centauri.png",
				Reader: logoFile,
			},
		},
	})
	if err != nil {
		botClient.Logger.Error(err.Error())
	}

	logoFile, err = os.Open("./documentation/static/img/centauri.png")
	if err != nil {
		panic(err)
	}

	// Embed
	_, err = botClient.CreateMessage(123456789, discord.CreateMessage{
		Embeds: []discord.Embed{
			discord.NewEmbedBuilder().SetImage("attachment://centauri.png"),
		},
		Files: []discord.File{
			{
				Name:   "centauri.png",
				Reader: logoFile,
			},
		},
	})
	if err != nil {
		botClient.Logger.Error(err.Error())
	}
}
