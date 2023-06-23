---
title: Gateway Client
description: This guide will explain how to use the Gateway Client and listen to events.
---

When you want to get events coming from Discord, for example when a message is sent in a channel or when a user joined a guild, then you will need a so called **Gateway Client**.

You can see an example of a Gateway Client [here](https://github.com/kkrypt0nn/centauri/blob/main/_examples/gateway_client/main.go).

## Creating a Gateway Client

Creating and then using a Gateway Client is not rocket science, no really - it isn't. Similarly to the [REST Client](/docs/guides/rest_client) you will just need to call the `NewGatewayClient` function on the `centauri` package like the following.

```go
botClient := centauri.NewGatewayClient("Bot BOT_TOKEN", 0)
```

## Arguments

The arguments you **must** pass into the function are the token and the intents you want. If you don't know what intents are, check [this guide](/docs/guides/intents) that explains it.

Similarly to the REST Client, the token must be prefixed with `Bot`.

## Listening to events

To listen to events you just need to call the `On()` method on the Gateway Client variable you've created.

As first argument you pass the **type** of the event, which you can easily list with IntelliSense and typing `gateway.EventType` - one example is `EventTypeMessageCreate`.

As second argument you will pass the callback function that will get executed when the event is dispatched by Discord. It will always be like the following:

```go
func(c *gateway.Client, x *gateway.<EventType>) {
	// ...
}
```

Obviously you have to replace `<EventType>` with the event type you have, for example `gateway.MessageCreate`.

The entire nonsense above results in something like

```go
botClient.On(gateway.EventTypeMessageCreate, func(c *gateway.Client, message *gateway.MessageCreate) {
    botClient.Logger.Info(fmt.Sprintf("Got a new message from %s: %s", message.Author.Username, message.Content))
})
```

to listen to events when messages are created (sent).

## Login and close

After creating all of the above, you will have to call the `Login()` method, otherwise your bot will never login and come online. If there are some errors or signals coming that the program has ended, you will want to call the `Close()` method. For example

```go
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
```

## Sharding

Sharding is something that is coming soon, so stay tuned!

## Using REST Client with Gateway Client

Yes, you can [use the REST Client[(/docs/guides/rest_client/#using-the-client)] when just creating a Gateway Client and you won't need additional steps. All you have to do is call the `Rest()` method on the Gateway Client variable and you will then have access to everything needed.