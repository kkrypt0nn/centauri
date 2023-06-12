package discord

// Role represents a set of permissions attached to a group of users
// https://discord.com/developers/docs/topics/permissions#role-object-role-structure
type Role struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Color        int         `json:"color"`
	Hoist        bool        `json:"hoist"`
	Icon         string      `json:"icon,omitempty"`
	UnicodeEmoji string      `json:"unicode_emoji"`
	Position     int         `json:"position"`
	Permissions  Permissions `json:"permissions,string"`
	Managed      bool        `json:"managed"`
	Mentionable  bool        `json:"mentionable"`
	Tags         *RoleTags   `json:"tags,omitempty"`
}

// RoleTags represents the tags the role (discord.Role) has
// https://discord.com/developers/docs/topics/permissions#role-object-role-tags-structure
type RoleTags struct {
	BotID                 string `json:"bot_id"`
	IntegrationID         string `json:"integration_id"`
	PremiumSubscriber     bool   `json:"premium_subscriber"`
	SubscriptionListingID string `json:"subscription_listing_id"`
	AvailableForPurchase  bool   `json:"available_for_purchase"`
	GuildConnections      bool   `json:"guild_connections"`
}

// RoleSubscriptionData represents the data of the role subscription purchase or renewal that prompted the message (discord.MessageTypeRoleSubscriptionPurchase)
// https://discord.com/developers/docs/resources/channel#role-subscription-data-object-role-subscription-data-object-structure
type RoleSubscriptionData struct {
	RoleSubscriptionListingID string `json:"role_subscription_listing_id"`
	TierName                  string `json:"tier_name"`
	TotalMonthsSubscribed     int    `json:"total_months_subscribed"`
	IsRenewal                 bool   `json:"is_renewal"`
}

// CreateGuildRole represents the payload to send to Discord to create a new role (discord.Role) in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#create-guild-role-json-params
type CreateGuildRole struct {
	Name         *string      `json:"name,omitempty"`
	Permissions  *Permissions `json:"permissions,string,omitempty"`
	Color        *int         `json:"color,omitempty"`
	Hoist        *bool        `json:"hoist,omitempty"`
	Icon         *string      `json:"icon,omitempty"`
	UnicodeEmoji *string      `json:"unicode_emoji,omitempty"`
	Mentionable  *bool        `json:"mentionable,omitempty"`

	AuditLogReason string `json:"-"`
}

// ModifyGuildRole represents the payload to send to Discord to modify an existing role (discord.Role) in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#modify-guild-role-json-params
type ModifyGuildRole struct {
	Name         *string      `json:"name,omitempty"`
	Permissions  *Permissions `json:"permissions,string,omitempty"`
	Color        *int         `json:"color,omitempty"`
	Hoist        *bool        `json:"hoist,omitempty"`
	Icon         *string      `json:"icon,omitempty"`
	UnicodeEmoji *string      `json:"unicode_emoji,omitempty"`
	Mentionable  *bool        `json:"mentionable,omitempty"`

	AuditLogReason string `json:"-"`
}

// GuildRolePosition represents an element of the array of the payload to send to Discord to modify the positions of roles (discord.Role) in a guild (discord.Guild)
// https://discord.com/developers/docs/resources/guild#modify-guild-role-positions-json-params
type GuildRolePosition struct {
	ID       string `json:"id"`
	Position *int   `json:"position,omitempty"`
}
