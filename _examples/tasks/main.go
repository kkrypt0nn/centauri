package main

import (
	"fmt"
	"github.com/kkrypt0nn/centauri"
	"time"
)

func main() {
	botClient := centauri.NewRestClient("Bot BOT_TOKEN")
	botClient.Debug = true

	// Create a new task
	getCurrentUserTask := botClient.TaskManager.NewTask(func() {
		bot, err := botClient.GetCurrentUser()
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println(fmt.Sprintf("Got bot %s with ID %s", bot.Username, bot.ID))
	}, 5*time.Second)
	// Start the task
	getCurrentUserTask.Start()

	// Sleep for 11 seconds and stop the running task
	time.Sleep(11 * time.Second)
	getCurrentUserTask.Stop()
}
