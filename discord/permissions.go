package discord

import "github.com/kkrypt0nn/centauri/utils/flags"

type Permissions uint64

const PermissionsNone Permissions = 0
const (
	PermissionsCreateInstantInvite Permissions = 1 << iota
	PermissionsKickMembers
	PermissionsBanMembers
	PermissionsAdministrator
	PermissionsManageChannels
	PermissionsManageGuild
	PermissionsAddReactions
	PermissionsViewAuditLog
	PermissionsPrioritySpeaker
	PermissionsStream
	PermissionsViewChannel
	PermissionsSendMessages
	PermissionsSendTTSMessages
	PermissionsManageMessages
	PermissionsEmbedLinks
	PermissionsAttachFiles
	PermissionsReadMessageHistory
	PermissionsMentionEveryone
	PermissionsUseExternalEmojis
	PermissionsViewGuildInsights
	PermissionsConnect
	PermissionsSpeak
	PermissionsMuteMembers
	PermissionsDeafenMembers
	PermissionsMoveMembers
	PermissionsUseVoiceActivityDetection
	PermissionsChangeNickname
	PermissionsManageNicknames
	PermissionsManageRoles
	PermissionsManageWebhooks
	PermissionsManageExpressions
	PermissionsUseApplicationCommands
	PermissionsRequestToSpeak
	PermissionsManageEvents
	PermissionsManageThreads
	PermissionsCreatePublicThreads
	PermissionsCreatePrivateThreads
	PermissionsUseExternalStickers
	PermissionsSendMessagesInThreads
	PermissionsUseEmbeddedActivities
	PermissionsModerateMembers
	PermissionsViewCreatorMonetizationAnalytics
	PermissionsUseSoundboard
	PermissionsCreateExpressions
	_
	_
	PermissionsSendVoiceMessages
)

// Compute creates a new permissions structure (discord.Permissions) from the given permissions
func (p Permissions) Compute(permissions ...Permissions) Permissions {
	return flags.Compute(permissions...)
}

// Add adds the given permissions (discord.Permissions) to the current permissions
func (p Permissions) Add(permissions ...Permissions) Permissions {
	return flags.Add(p, permissions...)
}

// Remove removes the given permissions (discord.Permissions) from the current permissions
func (p Permissions) Remove(permissions ...Permissions) Permissions {
	return flags.Remove(p, permissions...)
}

// Toggle toggles the given permissions (discord.Permissions) in the current permissions
func (p Permissions) Toggle(permissions ...Permissions) Permissions {
	return flags.Toggle(p, permissions...)
}

// Has checks if the given permissions (discord.Permissions) are the current permissions
func (p Permissions) Has(permissions ...Permissions) bool {
	return flags.Has(p, permissions...)
}

// HasNot checks if the given permissions (discord.Permissions) are not in the current permissions
func (p Permissions) HasNot(permissions ...Permissions) bool {
	return flags.HasNot(p, permissions...)
}
