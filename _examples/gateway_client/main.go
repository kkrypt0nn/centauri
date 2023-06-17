package main

import (
	"github.com/kkrypt0nn/centauri"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	botClient := centauri.NewGatewayClient("Bot MTAyNDM1Njk3Njk0MTA4ODg1OQ.G7h4xH.kNLn5uIhTfBDZWyelAgTxgONnZKgCtFlmHV_DM")
	botClient.Debug = true
	err := botClient.Login()
	if err != nil {
		log.Fatal(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	botClient.Close()
}
