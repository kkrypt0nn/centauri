package oauth2

type Scope string

const (
	ScopeActivitiesRead                        Scope = "activities.read"
	ScopeActivitiesWrite                       Scope = "activities.write"
	ScopeApplicationsBuildsRead                Scope = ""
	ScopeApplicationsBuildsUpload              Scope = ""
	ScopeApplicationsCommands                  Scope = ""
	ScopeApplicationsCommandsUpdate            Scope = ""
	ScopeApplicationsCommandsPermissionsUpdate Scope = ""
	ScopeApplicationsEntitlements              Scope = ""
	ScopeApplicationsStoreUpdate               Scope = ""
	ScopeBot                                   Scope = ""
	ScopeConnections                           Scope = ""
	ScopeDMChannelsRead                        Scope = ""
	ScopeEmail                                 Scope = ""
	ScopeGroupDMJoin                           Scope = ""
	ScopeGuilds                                Scope = ""
	ScopeGuildsJoin                            Scope = ""
	ScopeGuildsMembersRead                     Scope = ""
	ScopeIdentify                              Scope = ""
	ScopeMessagesRead                          Scope = ""
	ScopeRelationshipsRead                     Scope = ""
	ScopeRoleConnectionsWrite                  Scope = ""
	ScopeRPC                                   Scope = ""
	ScopeRPCActivitiesRead                     Scope = ""
	ScopeRPCNotificationsRead                  Scope = ""
	ScopeRPCVoiceRead                          Scope = ""
	ScopeRPCVoiceWrite                         Scope = ""
	ScopeVoice                                 Scope = ""
	ScopeWebhookIncoming                       Scope = ""
)
