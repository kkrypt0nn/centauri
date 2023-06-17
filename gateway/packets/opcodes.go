package packets

// OpCode represents operation codes from the Gateway API
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-opcodes
type OpCode int

const (
	OpCodeDispatch OpCode = iota
	OpCodeHeartbeat
	OpCodeIdentify
	OpCodePresenceUpdate
	OpCodeVoiceStateUpdate
	_
	OpCodeResume
	OpCodeReconnect
	OpCodeRequestGuildMembers
	OpCodeInvalidSession
	OpCodeHello
	OpCodeHeartbeatACK
)
