package discord

import (
	"fmt"
	"github.com/kkrypt0nn/centauri/utils/flags"
	"time"
)

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
	ID                             Snowflake        `json:"id"`
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
	GuildID                        Snowflake        `json:"guild_id,omitempty"`
	PrimarySKUID                   Snowflake        `json:"primary_sku_id,omitempty"`
	Slug                           string           `json:"slug,omitempty"`
	CoverImage                     string           `json:"cover_image,omitempty"`
	Flags                          ApplicationFlags `json:"flags,omitempty"`
	Tags                           []string         `json:"tags,omitempty"`
	InstallParams                  InstallParams    `json:"install_params,omitempty"`
	CustomInstallURL               string           `json:"custom_install_url,omitempty"`
	RoleConnectionsVerificationURL string           `json:"role_connections_verification_url,omitempty"`
}

// CreatedAt returns the creation time of the application (discord.Application)
func (a *Application) CreatedAt() time.Time {
	return a.ID.CreatedAt()
}

// IconURL returns the icon URL of the application (discord.Application)
func (a *Application) IconURL(asFormat ImageFormat) string {
	if a.Icon != "" {
		if asFormat == "" {
			asFormat = "png"
		}

		suffix := fmt.Sprintf("%s.%s", a.Icon, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/app-icons/%d/%s", a.ID, suffix)
	}
	return ""
}

// CoverURL returns the cover URL of the application (discord.Application)
func (a *Application) CoverURL(asFormat ImageFormat) string {
	if a.CoverImage != "" {
		if asFormat == "" {
			asFormat = "png"
		}

		suffix := fmt.Sprintf("%s.%s", a.CoverImage, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/app-icons/%d/%s.png", a.ID, suffix)
	}
	return ""
}

// Team represents a group of developers on Discord who want to collaborate on applications
// https://discord.com/developers/docs/topics/teams#data-models-team-object
type Team struct {
	ID          Snowflake    `json:"id"`
	Icon        string       `json:"icon,omitempty"`
	Members     []TeamMember `json:"members"`
	Name        string       `json:"name"`
	OwnerUserID string       `json:"owner_user_id"`
}

// CreatedAt returns the creation time of the team (discord.Team)
func (t *Team) CreatedAt() time.Time {
	return t.ID.CreatedAt()
}

// IconURL returns the icon URL of the team (discord.Team)
func (t *Team) IconURL(asFormat ImageFormat) string {
	if t.Icon != "" {
		if asFormat == "" {
			asFormat = "png"
		}

		suffix := fmt.Sprintf("%s.%s", t.Icon, asFormat)
		return fmt.Sprintf("https://cdn.discordapp.com/team-icons/%d/%s", t.ID, suffix)
	}
	return ""
}

// TeamMember represents a member of a Discord team (discord.Team)
// https://discord.com/developers/docs/topics/teams#data-models-team-member-object
type TeamMember struct {
	MembershipState MembershipState `json:"membership_state"`
	Permissions     []string        `json:"permissions"`
	TeamID          Snowflake       `json:"team_id"`
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

// ComputeApplicationFlags creates a new application flags structure (discord.ApplicationFlags) from the given application flags
func ComputeApplicationFlags(applicationFlags ...ApplicationFlags) ApplicationFlags {
	return flags.Compute(applicationFlags...)
}

// Add adds the given application flags (discord.ApplicationFlags)
func (f ApplicationFlags) Add(applicationFlags ...ApplicationFlags) ApplicationFlags {
	return flags.Add(f, applicationFlags...)
}

// Remove removes the given application flags (discord.ApplicationFlags)
func (f ApplicationFlags) Remove(applicationFlags ...ApplicationFlags) ApplicationFlags {
	return flags.Remove(f, applicationFlags...)
}

// Toggle toggles the given application flags (discord.ApplicationFlags)
func (f ApplicationFlags) Toggle(applicationFlags ...ApplicationFlags) ApplicationFlags {
	return flags.Toggle(f, applicationFlags...)
}

// Has checks if all the given application flags (discord.ApplicationFlags) are set
func (f ApplicationFlags) Has(applicationFlags ...ApplicationFlags) bool {
	return flags.Has(f, applicationFlags...)
}

// HasAny checks if any of the given application flags (discord.ApplicationFlags) is set
func (f ApplicationFlags) HasAny(applicationFlags ...ApplicationFlags) bool {
	return flags.HasAny(f, applicationFlags...)
}

// HasNot checks if all the given application flags (discord.ApplicationFlags) are not set
func (f ApplicationFlags) HasNot(applicationFlags ...ApplicationFlags) bool {
	return flags.HasNot(f, applicationFlags...)
}

// HasNotAny checks if any of the given application flags (discord.ApplicationFlags) is not set
func (f ApplicationFlags) HasNotAny(applicationFlags ...ApplicationFlags) bool {
	return flags.HasNotAny(f, applicationFlags...)
}

// InstallParams represents the application's default in-app authorization link
// https://discord.com/developers/docs/resources/application#install-params-object-install-params-structure
type InstallParams struct {
	Scopes      []Scope     `json:"scopes"`
	Permissions Permissions `json:"permissions"`
}
