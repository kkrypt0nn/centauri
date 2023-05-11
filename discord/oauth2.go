package discord

type OAuthScopes string

const (
	OAuthScopeActivitiesRead                        OAuthScopes = "activities.read"
	OAuthScopeActivitiesWrite                       OAuthScopes = "activities.write"
	OAuthScopeApplicationsBuildsRead                OAuthScopes = ""
	OAuthScopeApplicationsBuildsUpload              OAuthScopes = ""
	OAuthScopeApplicationsCommands                  OAuthScopes = ""
	OAuthScopeApplicationsCommandsUpdate            OAuthScopes = ""
	OAuthScopeApplicationsCommandsPermissionsUpdate OAuthScopes = ""
	OAuthScopeApplicationsEntitlements              OAuthScopes = ""
	OAuthScopeApplicationsStoreUpdate               OAuthScopes = ""
	OAuthScopeBot                                   OAuthScopes = ""
	OAuthScopeConnections                           OAuthScopes = ""
	OAuthScopeDMChannelsRead                        OAuthScopes = ""
	OAuthScopeEmail                                 OAuthScopes = ""
	OAuthScopeGroupDMJoin                           OAuthScopes = ""
	OAuthScopeGuilds                                OAuthScopes = ""
	OAuthScopeGuildsJoin                            OAuthScopes = ""
	OAuthScopeGuildsMembersRead                     OAuthScopes = ""
	OAuthScopeIdentify                              OAuthScopes = ""
	OAuthScopeMessagesRead                          OAuthScopes = ""
	OAuthScopeRelationshipsRead                     OAuthScopes = ""
	OAuthScopeRoleConnectionsWrite                  OAuthScopes = ""
	OAuthScopeRPC                                   OAuthScopes = ""
	OAuthScopeRPCActivitiesRead                     OAuthScopes = ""
	OAuthScopeRPCNotificationsRead                  OAuthScopes = ""
	OAuthScopeRPCVoiceRead                          OAuthScopes = ""
	OAuthScopeRPCVoiceWrite                         OAuthScopes = ""
	OAuthScopeVoice                                 OAuthScopes = ""
	OAuthScopeWebhookIncoming                       OAuthScopes = ""
)
