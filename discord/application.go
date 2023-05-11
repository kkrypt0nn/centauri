package discord

import "github.com/kkrypt0nn/centauri/oauth2"

type ApplicationRoleConnectionMetadata struct {
	Type                     ApplicationRoleConnectionMetadataType `json:"type"`
	Key                      string                                `json:"key"`
	Name                     string                                `json:"name"`
	NameLocalizations        map[Locale]string                     `json:"name_localizations,omitempty"`
	Description              string                                `json:"description"`
	DescriptionLocalizations map[Locale]string                     `json:"description_localizations,omitempty"`
}

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

type Team struct {
	Icon        string       `json:"icon,omitempty"`
	ID          string       `json:"id"`
	Members     []TeamMember `json:"members"`
	Name        string       `json:"name"`
	OwnerUserID string       `json:"owner_user_id"`
}

type TeamMember struct {
	MembershipState MembershipState `json:"membership_state"`
	Permissions     []string        `json:"permissions"`
	TeamID          string          `json:"team_id"`
	User            *User           `json:"user"`
}

type MembershipState int

const (
	MembershipStateInvited MembershipState = 1 + iota
	MembershipStateAccepted
)

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

type InstallParams struct {
	Scopes      []oauth2.Scope `json:"scopes"`
	Permissions string         `json:"permissions"`
}
