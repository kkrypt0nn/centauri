# Centauri

[![Go Reference Badge](https://pkg.go.dev/badge/github.com/kkrypt0nn/centauri.svg)](https://pkg.go.dev/github.com/kkrypt0nn/centauri)
[![CI Badge](https://github.com/kkrypt0nn/centauri/actions/workflows/ci.yml/badge.svg)](https://github.com/kkrypt0nn/centauri/actions)
[![Discord Server Badge](https://img.shields.io/discord/1095412499416891482?logo=discord)](https://discord.gg/feA6ZGRgpw)
[![Last Commit Badge](https://img.shields.io/github/last-commit/kkrypt0nn/centauri)](https://github.com/kkrypt0nn/centauri/commits/main)
[![Conventional Commits Badge](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org/en/v1.0.0/)

> **Note** Due to my military service from **July 2023** until **April 2024** I will not be able to work on this
> project a lot, most likely only during weekends if I have time and motivation ^-^

> **Warning** This is a WIP library that may be **very unstable** and eventually **not fully optimised**, use at your
> own
> care! This README will also be reworked.

A Discord API wrapper written in Go with the goal of being easily understandable and simple to use, even for newcomers.

## Features

The plan for Centauri would be for it to cover Discord's APIs in their entirety. A small list of features that will be,
or are already in Centauri are:

- ANSI Colors
- Background Tasks
- Caching
- Easy Commands Creation
- Gateway API
- Interactions (Buttons, Modals, etc.)
- OAuth 2.0
- Rate Limiter
- REST API
- RPC
- Sharding
- Webhooks

## Getting Started

### Installation

To get started you will simply need to install the library in your project by executing the following command.

```bash
go get github.com/kkrypt0nn/centauri
```

### Example Usage

#### REST Client

If you just want to interact with the REST API, you may use the following code.

```go
package main

import (
	"fmt"

	"github.com/kkrypt0nn/centauri"
)

func main() {
	botClient := centauri.NewRestClient("Bot BOT_TOKEN")
	bot, err := botClient.GetCurrentUser()
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(fmt.Sprintf("Got bot %s with ID %d", bot.Username, bot.ID))
	}
}
```

#### Gateway Client

If you want to handle events coming from Discord, you may use the following code.

```go
package main

import (
	"fmt"
	"github.com/kkrypt0nn/centauri"
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/gateway"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	botClient := centauri.NewGatewayClient("Bot BOT_TOKEN", discord.IntentsGuildMessages|discord.IntentsMessageContent)

	botClient.On(gateway.EventTypeReady, func(c *gateway.Client, ready *gateway.Ready) {
		botClient.Logger.Info(fmt.Sprintf("We are now logged in as %s", ready.User.Username))
	})
	botClient.On(gateway.EventTypeMessageCreate, func(c *gateway.Client, message *gateway.MessageCreate) {
		botClient.Logger.Info(fmt.Sprintf("Got a new message from %s: %s", message.Author.Username, message.Content))
	})

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
}
```

More examples are available [here](_examples)

## Documentation

The [documentation](documentation) is available for reading [here](https://centauri.krypton.ninja). There is also the Go reference page available [here](https://pkg.go.dev/github.com/kkrypt0nn/centauri)

## Troubleshooting

Some troubleshooting articles and guides may be found in the [documentation](https://centauri.krypton.ninja/docs/category/troubleshooting/) mentioned above.

## Contributing

Before the version **1.0.0** of the project, contributions are mostly closed as I'd like to lay down the foundations of
the library. Afterward contributions will be more than welcome if following
the [Contributing Guidelines](CONTRIBUTING.md).

## License

This library was made with 💜 by Krypton and is under the [MIT License](LICENSE.md).