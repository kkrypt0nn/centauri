package discord

import (
	"github.com/kkrypt0nn/centauri/oauth2"
	"time"
)

// Integration represents an integration that is contained within a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#integration-object-integration-structure
type Integration struct {
	ID                string                    `json:"id"`
	Name              string                    `json:"name"`
	Type              string                    `json:"type"`
	Enabled           bool                      `json:"enabled"`
	Syncing           bool                      `json:"syncing,omitempty"`
	RoleID            string                    `json:"role_id,omitempty"`
	EnableEmoticons   bool                      `json:"enable_emoticons,omitempty"`
	ExpireBehavior    IntegrationExpireBehavior `json:"expire_behavior,omitempty"`
	ExpireGracePeriod int                       `json:"expire_grace_period,omitempty"`
	User              *User                     `json:"user,omitempty"`
	Account           IntegrationAccount        `json:"account"`
	SyncedAt          *time.Time                `json:"synced_at,omitempty"`
	SubscriberCount   int                       `json:"subscriber_count,omitempty"`
	Revoked           bool                      `json:"revoked,omitempty"`
	Application       IntegrationApplication    `json:"application,omitempty"`
	Scopes            []oauth2.Scope            `json:"scopes,omitempty"`
}

// IntegrationExpireBehavior represents the behavior of expiring subscribers
// https://discord.com/developers/docs/resources/guild#integration-object-integration-expire-behaviors
type IntegrationExpireBehavior int

const (
	IntegrationExpireBehaviorRemoveRole IntegrationExpireBehavior = iota
	IntegrationExpireBehaviorKick
)

// IntegrationAccount represents the integration account information
// https://discord.com/developers/docs/resources/guild#integration-account-object-integration-account-structure
type IntegrationAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// IntegrationApplication represents the bot/OAuth2 application for the integration (discord.Integration)
// https://discord.com/developers/docs/resources/guild#integration-application-object-integration-application-structure
type IntegrationApplication struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Bot         *User  `json:"bot,omitempty"`
}