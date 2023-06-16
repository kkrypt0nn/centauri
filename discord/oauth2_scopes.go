package discord

// Scope represents the various OAuth 2 scopes available
// https://discord.com/developers/docs/topics/oauth2#shared-resources-oauth2-scopes
type Scope string

const (
	ScopeActivitiesRead                        Scope = "activities.read"
	ScopeActivitiesWrite                       Scope = "activities.write"
	ScopeApplicationsBuildsRead                Scope = "applications.builds.read"
	ScopeApplicationsBuildsUpload              Scope = "applications.builds.upload"
	ScopeApplicationsCommands                  Scope = "applications.commands"
	ScopeApplicationsCommandsUpdate            Scope = "applications.commands.update"
	ScopeApplicationsCommandsPermissionsUpdate Scope = "applications.commands.permissions.update"
	ScopeApplicationsEntitlements              Scope = "applications.entitlements"
	ScopeApplicationsStoreUpdate               Scope = "applications.store.update"
	ScopeBot                                   Scope = "bot"
	ScopeConnections                           Scope = "connections"
	ScopeDMChannelsRead                        Scope = "dm_channels.read"
	ScopeEmail                                 Scope = "email"
	ScopeGroupDMJoin                           Scope = "gdm.join"
	ScopeGuilds                                Scope = "guilds"
	ScopeGuildsJoin                            Scope = "guilds.join"
	ScopeGuildsMembersRead                     Scope = "guilds.members.read"
	ScopeIdentify                              Scope = "identify"
	ScopeMessagesRead                          Scope = "messages.read"
	ScopeRelationshipsRead                     Scope = "relationships.read"
	ScopeRoleConnectionsWrite                  Scope = "role_connections.write"
	ScopeRPC                                   Scope = "rpc"
	ScopeRPCActivitiesRead                     Scope = "rpc.activities.write"
	ScopeRPCNotificationsRead                  Scope = "rpc.notifications.read"
	ScopeRPCVoiceRead                          Scope = "rpc.voice.read"
	ScopeRPCVoiceWrite                         Scope = "rpc.voice.write"
	ScopeVoice                                 Scope = "voice"
	ScopeWebhookIncoming                       Scope = "webhook.incoming"
)
