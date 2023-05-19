package discord

// Reaction represents a reaction that has been added to a message (discord.Message)
// https://discord.com/developers/docs/resources/channel#reaction-object-reaction-structure
type Reaction struct {
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emoji `json:"emoji"`
}
