package discord

// RateLimitExceeded represents content sent by Discord once a rate limit has been exceeded
// https://discord.com/developers/docs/topics/rate-limits#exceeding-a-rate-limit-example-exceeded-user-rate-limit-response
type RateLimitExceeded struct {
	Message    string  `json:"message"`
	RetryAfter float64 `json:"retry_after"`
	Global     bool    `json:"global"`
	Code       int     `json:"code,omitempty"`
}
