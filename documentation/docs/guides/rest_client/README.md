---
title: REST Client
description: This guide will explain how to use the REST Client and perform various actions.
---

The REST Client is a client that will facilitate the interaction with Discord's REST API to - for example - send messages or create new channels.

You can see an example of a REST Client [here](https://github.com/kkrypt0nn/centauri/blob/main/_examples/rest_client/main.go).

## Creating a REST Client

To create a REST Client you just need to call the `NewRestClient()` function on the `centauri` package like the following:

```go
botClient := centauri.NewRestClient("Bot BOT_TOKEN")
```

## Arguments

As argument you just need to pass your bot's token with the `Bot` prefix before.

## Using the client

To use the client let IntelliSense help you, simply write down `botClient.` and it will show existing methods you can use. If there are any arguments needed, you can check the [Go Reference documentation](https://pkg.go.dev/github.com/kkrypt0nn/centauri/rest#Client) to know what kind of arguments come where, otherwise you can also `Ctrl/Cmd+Left Click` the method in your favorite IDE.

:::tip

When you create or edit things like channels, messages, etc. the structure for these functions are *almost* always `discord.X` where `X` is the method you're calling. For example:

```go
botClient.CreateGlobalApplicationCommand(discord.CreateGlobalApplicationCommand{
    Name:        "test",
    Description: ptr.New("test"),
})
```

Here you can see the method is `CreateGlobalApplicationCommand` and the argument struture is `CreateGlobalApplicationCommand` as well, that makes it easier to remember.

:::