package gateway

// CloseCode represents a code sent when the connection is closed
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-close-event-codes
type CloseCode int

const (
	CloseCodeUnknownError CloseCode = 4000 + iota
	CloseCodeUnknownOpcode
	CloseCodeDecodeError
	CloseCodeNotAuthenticated
	CloseCodeAuthenticationFailed
	CloseCodeAlreadyAuthenticated
	_
	CloseCodeInvalidSequence
	CloseCodeRateLimited
	CloseCodeSessionTimedOut
	CloseCodeInvalidShard
	CloseCodeShardingRequired
	CloseCodeInvalidAPIVersion
	CloseCodeInvalidIntents
	CloseCodeDisallowedIntents
)

// ShouldReconnect returns whether the client should attempt a reconnection after the given close code (packets.CloseCode)
func (c CloseCode) ShouldReconnect() bool {
	switch c {
	case CloseCodeAuthenticationFailed, CloseCodeInvalidShard, CloseCodeShardingRequired, CloseCodeInvalidAPIVersion, CloseCodeInvalidIntents, CloseCodeDisallowedIntents:
		return false
	default:
		return true
	}
}
