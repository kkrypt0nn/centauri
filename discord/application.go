package discord

import "github.com/kkrypt0nn/centauri/oauth2"

// ApplicationRoleConnectionMetadata represents a role connection metadata for an application (discord.Application)
// https://discord.com/developers/docs/resources/application-role-connection-metadata#application-role-connection-metadata-object-application-role-connection-metadata-structure
type ApplicationRoleConnectionMetadata struct {
	Type                     ApplicationRoleConnectionMetadataType `json:"type"`
	Key                      string                                `json:"key"`
	Name                     string                                `json:"name"`
	NameLocalizations        map[Locale]string                     `json:"name_localizations,omitempty"`
	Description              string                                `json:"description"`
	DescriptionLocalizations map[Locale]string                     `json:"description_localizations,omitempty"`
}

// ApplicationRoleConnectionMetadataType represents the type of metadata value
// https://discord.com/developers/docs/resources/application-role-connection-metadata#application-role-connection-metadata-object-application-role-connection-metadata-type
type ApplicationRoleConnectionMetadataType int

const (
	ApplicationRoleConnectionMetadataTypeIntegerLessThanOrEqual ApplicationRoleConnectionMetadataType = 1 + iota
	ApplicationRoleConnectionMetadataTypeIntegerGreaterThanOrEqual
	ApplicationRoleConnectionMetadataTypeIntegerEqual
	ApplicationRoleConnectionMetadataTypeIntegerNotEqual
	ApplicationRoleConnectionMetadataTypeDatetimeLessThanOrEqual
	ApplicationRoleConnectionMetadataTypeDatetimeGreaterThanOrEqual
	ApplicationRoleConnectionMetadataTypeBooleanEqual
	ApplicationRoleConnectionMetadataTypeBooleanNotEqual
)

// Application represents a Discord application
// https://discord.com/developers/docs/resources/application#application-object-application-structure
type Application struct {
	ID                             string           `json:"id"`
	Name                           string           `json:"name"`
	Icon                           string           `json:"icon,omitempty"`
	Description                    string           `json:"description"`
	RPCOrigins                     []string         `json:"rpc_origins,omitempty"`
	BotPublic                      bool             `json:"bot_public"`
	BotRequireCodeGrant            bool             `json:"bot_require_code_grant"`
	TermsOfServiceURL              string           `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL               string           `json:"privacy_policy_url,omitempty"`
	Owner                          *User            `json:"owner,omitempty"`
	VerifyKey                      string           `json:"verify_key"`
	Team                           *Team            `json:"team,omitempty"`
	GuildID                        string           `json:"guild_id,omitempty"`
	PrimarySKUID                   string           `json:"primary_sku_id,omitempty"`
	Slug                           string           `json:"slug,omitempty"`
	CoverImage                     string           `json:"cover_image,omitempty"`
	Flags                          ApplicationFlags `json:"flags,omitempty"`
	Tags                           []string         `json:"tags,omitempty"`
	InstallParams                  InstallParams    `json:"install_params,omitempty"`
	CustomInstallURL               string           `json:"custom_install_url,omitempty"`
	RoleConnectionsVerificationURL string           `json:"role_connections_verification_url,omitempty"`
}

// Team represents a group of developers on Discord who want to collaborate on applications
// https://discord.com/developers/docs/topics/teams#data-models-team-object
type Team struct {
	Icon        string       `json:"icon,omitempty"`
	ID          string       `json:"id"`
	Members     []TeamMember `json:"members"`
	Name        string       `json:"name"`
	OwnerUserID string       `json:"owner_user_id"`
}

// TeamMember represents a member of a Discord team (discord.Team)
// https://discord.com/developers/docs/topics/teams#data-models-team-member-object
type TeamMember struct {
	MembershipState MembershipState `json:"membership_state"`
	Permissions     []string        `json:"permissions"`
	TeamID          string          `json:"team_id"`
	User            *User           `json:"user"`
}

// MembershipState represents the membership state of a team member (discord.TeamMember)
// https://discord.com/developers/docs/topics/teams#data-models-membership-state-enum
type MembershipState int

const (
	MembershipStateInvited MembershipState = 1 + iota
	MembershipStateAccepted
)

// ApplicationFlags represents the list of public flags of an application (discord.Application)
// https://discord.com/developers/docs/resources/application#application-object-application-flags
type ApplicationFlags uint64

const ApplicationFlagsNone ApplicationFlags = 0
const (
	_ ApplicationFlags = 1 << iota
	_
	_
	_
	_
	_
	ApplicationFlagApplicationAutoModerationRuleCreateBadge
	_
	_
	_
	_
	_
	ApplicationFlagGatewayPresence
	ApplicationFlagGatewayPresenceLimited
	ApplicationFlagGatewayGuildMembers
	ApplicationFlagGatewayGuildMembersLimited
	ApplicationFlagVerificationPendingGuildLimit
	ApplicationFlagEmbedded
	ApplicationFlagGatewayMessageContent
	ApplicationFlagGatewayMessageContentLimited
	_
	_
	_
	ApplicationFlagApplicationCommandBadge
)

// InstallParams represents the application's default in-app authorization link
// https://discord.com/developers/docs/resources/application#install-params-object-install-params-structure
type InstallParams struct {
	Scopes      []oauth2.Scope `json:"scopes"`
	Permissions string         `json:"permissions"`
}
