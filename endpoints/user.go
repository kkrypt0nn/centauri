package endpoints

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

const (
	UsersEndpoint = Endpoint + "users"
)

func SelfUser() string {
	return fmt.Sprintf("%s/@me", UsersEndpoint)
}

func User(userID discord.Snowflake) string {
	return fmt.Sprintf("%s/%d", UsersEndpoint, userID)
}

func SelfUserGuilds() string {
	return fmt.Sprintf("%s/@me/guilds", UsersEndpoint)
}

func SelfUserChannels() string {
	return fmt.Sprintf("%s/@me/channels", UsersEndpoint)
}
