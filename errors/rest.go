package errors

import (
	"errors"
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

type RateLimitExceeded struct {
	*discord.RateLimitExceeded
}

func (r *RateLimitExceeded) Error() string {
	return fmt.Sprintf("Rate limit exceeded, retry after %g", r.RetryAfter)
}

var (
	Unauthorized = errors.New("the HTTP request was unauthorized, please double check the given token")
)
