package rest

import (
	"github.com/kkrypt0nn/centauri/discord"
	"github.com/kkrypt0nn/centauri/ext/tasks"
	"github.com/kkrypt0nn/logger.go"
	"net/http"
	"os"
	"time"
)

var (
	testingBotToken   = os.Getenv("CENTAURI_BOT_TOKEN")
	testingBotID      = discord.MustParseSnowflake(os.Getenv("CENTAURI_BOT_ID"))
	testingGuildID    = discord.MustParseSnowflake(os.Getenv("CENTAURI_GUILD_ID"))
	testingChannelID  = discord.MustParseSnowflake(os.Getenv("CENTAURI_CHANNEL_ID"))
	testingUserID     = discord.MustParseSnowflake(os.Getenv("CENTAURI_USER_ID"))
	testingRestClient = &Client{
		HttpClient: &HttpClient{
			Client: &http.Client{
				Timeout: 20 * time.Second,
			},
			Interceptor: nil,
		},
		Logger:              logger.NewLogger(),
		RateLimiter:         NewRateLimiter(),
		TaskManager:         tasks.NewTaskManager(),
		authorizationHeader: testingBotToken,
	}
)
