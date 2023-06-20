package discord

import "github.com/kkrypt0nn/centauri/utils/flags"

// Activity represents a user's activity on Discord
// https://discord.com/developers/docs/topics/gateway-events#activity-object
type Activity struct {
	Name          string             `json:"name"`
	Type          ActivityType       `json:"type"`
	URL           string             `json:"url,omitempty"`
	CreatedAt     int                `json:"created_at"`
	Timestamps    ActivityTimeStamps `json:"timestamps"`
	ApplicationID Snowflake          `json:"application_id"`
	Details       string             `json:"details,omitempty"`
	State         string             `json:"state,omitempty"`
	Emoji         Emoji              `json:"emoji,omitempty"`
	Party         ActivityParty      `json:"party,omitempty"`
	Assets        ActivityAssets     `json:"assets,omitempty"`
	Secrets       ActivitySecrets    `json:"secrets"`
	Instance      bool               `json:"instance"`
	Flags         ActivityFlags      `json:"flags"`
	Buttons       []ActivityButton   `json:"buttons"`
}

// ActivityType represents the type of activity (discord.Activity)
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-types
type ActivityType int

const (
	ActivityTypePlaying ActivityType = iota
	ActivityTypeStreaming
	ActivityTypeListening
	ActivityTypeWatching
	ActivityTypeCustom
	ActivityTypeCompeting
)

// ActivityTimeStamps represents the unix timestamps for the start and/or end of the activity (discord.Activity)
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-timestamps
type ActivityTimeStamps struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

// ActivityParty represents a party of a player
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-party
type ActivityParty struct {
	ID   string `json:"id"`
	Size [2]int `json:"size"`
}

// ActivityAssets represents the images for the presence and their hover texts
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-assets
type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

// ActivitySecrets represents the secrets for joining and spectating
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-secrets
type ActivitySecrets struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

// ActivityFlags represents the various flags an activity can have
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-flags
type ActivityFlags uint64

const ActivityFlagNone ActivityFlags = 0
const (
	ActivityFlagInstance ActivityFlags = 1 << iota
	ActivityFlagJoin
	ActivityFlagSpectate
	ActivityFlagJoinRequest
	ActivityFlagSync
	ActivityFlagPlay
	ActivityFlagPartyPrivacyFriends
	ActivityFlagPartyPrivacyVoiceChannel
	ActivityFlagEmbedded
)

// ComputeActivityFlags creates a new activity flags structure (discord.ActivityFlags) from the given activity flags
func ComputeActivityFlags(activityFlags ...ActivityFlags) ActivityFlags {
	return flags.Compute(activityFlags...)
}

// Add adds the given activity flags (discord.ActivityFlags)
func (f ActivityFlags) Add(activityFlags ...ActivityFlags) ActivityFlags {
	return flags.Add(f, activityFlags...)
}

// Remove removes the given activity flags (discord.ActivityFlags)
func (f ActivityFlags) Remove(activityFlags ...ActivityFlags) ActivityFlags {
	return flags.Remove(f, activityFlags...)
}

// Toggle toggles the given activity flags (discord.ActivityFlags)
func (f ActivityFlags) Toggle(activityFlags ...ActivityFlags) ActivityFlags {
	return flags.Toggle(f, activityFlags...)
}

// Has checks if all the given activity flags (discord.ActivityFlags) are set
func (f ActivityFlags) Has(activityFlags ...ActivityFlags) bool {
	return flags.Has(f, activityFlags...)
}

// HasAny checks if any of the given activity flags (discord.ActivityFlags) is set
func (f ActivityFlags) HasAny(activityFlags ...ActivityFlags) bool {
	return flags.HasAny(f, activityFlags...)
}

// HasNot checks if all the given activity flags (discord.ActivityFlags) are not set
func (f ActivityFlags) HasNot(activityFlags ...ActivityFlags) bool {
	return flags.HasNot(f, activityFlags...)
}

// HasNotAny checks if any of the given activity flags (discord.ActivityFlags) is not set
func (f ActivityFlags) HasNotAny(activityFlags ...ActivityFlags) bool {
	return flags.HasNotAny(f, activityFlags...)
}

// ActivityButton represents buttons in an activity (discord.Activity)
// https://discord.com/developers/docs/topics/gateway-events#activity-object-activity-buttons
type ActivityButton struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// ClientStatus represents the active sessions of a user on the various platforms
// https://discord.com/developers/docs/topics/gateway-events#client-status-object
type ClientStatus struct {
	Desktop string `json:"desktop,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
	Web     string `json:"web,omitempty"`
}

// StatusType represents the status of a user
// https://discord.com/developers/docs/topics/gateway-events#update-presence-status-types
type StatusType string

const (
	StatusTypeOnline    StatusType = "online"
	StatusTypeDND       StatusType = "dnd"
	StatusTypeIdle      StatusType = "idle"
	StatusTypeInvisible StatusType = "invisible"
	StatusTypeOffline   StatusType = "offline"
)
