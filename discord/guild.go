package discord

// PartialGuild represents a discord.Guild with only a few fields
// https://discord.com/developers/docs/resources/user#get-current-user-guilds-example-partial-guild
type PartialGuild struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Icon        string         `json:"icon"`
	Owner       bool           `json:"owner"`
	Permissions uint64         `json:"permissions,string"`
	Features    []GuildFeature `json:"features"`
}

// GuildFeature represents the different features a guild may have activated
// https://discord.com/developers/docs/resources/guild#guild-object-guild-features
type GuildFeature string

const (
	AnimatedBanner                        GuildFeature = "ANIMATED_BANNER"
	AnimatedIcon                          GuildFeature = "ANIMATED_ICON"
	ApplicationCommandPermissionsV2       GuildFeature = "APPLICATION_COMMAND_PERMISSIONS_V2"
	AutoModeration                        GuildFeature = "AUTO_MODERATION"
	Banner                                GuildFeature = "BANNER"
	Community                             GuildFeature = "COMMUNITY"
	CreatorMonetizableProvisional         GuildFeature = "CREATOR_MONETIZABLE_PROVISIONAL"
	CreatorStorePage                      GuildFeature = "CREATOR_STORE_PAGE"
	DeveloperSupportServer                GuildFeature = "DEVELOPER_SUPPORT_SERVER"
	Discoverable                          GuildFeature = "DISCOVERABLE"
	Featurable                            GuildFeature = "FEATURABLE"
	InvitesDisabled                       GuildFeature = "INVITES_DISABLED"
	InviteSplash                          GuildFeature = "INVITE_SPLASH"
	MemberVerificationGateEnabled         GuildFeature = "MEMBER_VERIFICATION_GATE_ENABLED"
	MoreStickers                          GuildFeature = "MORE_STICKERS"
	News                                  GuildFeature = "NEWS"
	Partnered                             GuildFeature = "PARTNERED"
	PreviewEnabled                        GuildFeature = "PREVIEW_ENABLED"
	RoleIcons                             GuildFeature = "ROLE_ICONS"
	RoleSubscriptionsAvailableForPurchase GuildFeature = "ROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE"
	RoleSubscriptionsEnabled              GuildFeature = "ROLE_SUBSCRIPTIONS_ENABLED"
	TicketedEventsEnabled                 GuildFeature = "TICKETED_EVENTS_ENABLED"
	VanityURL                             GuildFeature = "VANITY_URL"
	Verified                              GuildFeature = "VERIFIED"
	VIPRegions                            GuildFeature = "VIP_REGIONS"
	WelcomeScreenEnabled                  GuildFeature = "WELCOME_SCREEN_ENABLED"
)
