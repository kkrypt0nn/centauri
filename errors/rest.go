package errors

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/discord"
)

type RateLimitExceeded struct {
	*discord.RateLimitExceeded
}

func (r *RateLimitExceeded) Error() string {
	return fmt.Sprintf("rate limit exceeded, retry after %g", r.RetryAfter)
}
