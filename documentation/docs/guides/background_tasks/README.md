---
title: Background Tasks
description: This guide will explain how to run background tasks.
---

Background tasks are useful when you want to run something at every `X` interval. For that Centauri offers an extension.

## Example

To create a background task you just need to use the `TaskManager` of the REST Client and call the `NewTask()` method. This will let you pass as first argument a function callback to execute and as second argument the interval at which the task should be executed.

Once you've added the task, you can call `Start()` or `Stop()` on it to start and stop it.

The [example](https://github.com/kkrypt0nn/centauri/blob/main/_examples/tasks/main.go) below shows a task that will start and stop after 11 seconds a task to get the current user.

```go
// Create a new task
getCurrentUserTask := botClient.TaskManager.NewTask(func() {
	bot, err := botClient.GetCurrentUser()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(fmt.Sprintf("Got bot %s with ID %d", bot.Username, bot.ID))
}, 5*time.Second)
// Start the task
getCurrentUserTask.Start()

// Sleep for 11 seconds and stop the running task
time.Sleep(11 * time.Second)
getCurrentUserTask.Stop()
```