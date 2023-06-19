package gateway

// EventHandler is an interface for all the event handlers
type EventHandler interface {
	Handle(*Client, interface{})
}

// readyEventHandler is the event handler for the Ready event (gateway.Ready)
type readyEventHandler func(*Client, *Ready)

// Handle handles the Ready event (gateway.Ready)
func (h readyEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*Ready); ok {
		h(c, e)
	}
}

// New returns the Ready event (gateway.Ready) structure
func (h readyEventHandler) New() interface{} {
	return &Ready{}
}

// messageCreateEventHandler is the event handler for the Message Create event (gateway.MessageCreate)
type messageCreateEventHandler func(*Client, *MessageCreate)

// Handle handles the Message Create event (gateway.MessageCreate)
func (h messageCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*MessageCreate); ok {
		h(c, e)
	}
}

// New returns the Message Create event (gateway.MessageCreate) structure
func (h messageCreateEventHandler) New() interface{} {
	return &MessageCreate{}
}

// applicationCommandPermissionsUpdateEventHandler is the event handler for the Application Command Permissions Update event (gateway.ApplicationCommandPermissionsUpdate)
type applicationCommandPermissionsUpdateEventHandler func(*Client, *ApplicationCommandPermissionsUpdate)

// Handle handles the Application Command Permissions Update event (gateway.ApplicationCommandPermissionsUpdate)
func (h applicationCommandPermissionsUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*ApplicationCommandPermissionsUpdate); ok {
		h(c, e)
	}
}

// New returns the Application Command Permissions Update event (gateway.ApplicationCommandPermissionsUpdate) structure
func (h applicationCommandPermissionsUpdateEventHandler) New() interface{} {
	return &MessageCreate{}
}

// autoModerationRuleCreateEventHandler is the event handler for the Auto Moderation Rule Create event (gateway.AutoModerationRuleCreate)
type autoModerationRuleCreateEventHandler func(*Client, *AutoModerationRuleCreate)

// Handle handles the Auto Moderation Rule Create event (gateway.AutoModerationRuleCreate)
func (h autoModerationRuleCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*AutoModerationRuleCreate); ok {
		h(c, e)
	}
}

// New returns the Auto Moderation Rule Create event (gateway.AutoModerationRuleCreate)
func (h autoModerationRuleCreateEventHandler) New() interface{} {
	return &AutoModerationRuleCreate{}
}

// autoModerationRuleUpdateEventHandler is the event handler for the Auto Moderation Rule Update event (gateway.AutoModerationRuleUpdate)
type autoModerationRuleUpdateEventHandler func(*Client, *AutoModerationRuleUpdate)

// Handle handles the Auto Moderation Rule Update event (gateway.AutoModerationRuleUpdate)
func (h autoModerationRuleUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*AutoModerationRuleUpdate); ok {
		h(c, e)
	}
}

// New returns the Auto Moderation Rule Update event (gateway.AutoModerationRuleUpdate)
func (h autoModerationRuleUpdateEventHandler) New() interface{} {
	return &AutoModerationRuleUpdate{}
}

// autoModerationRuleDeleteEventHandler is the event handler for the Auto Moderation Rule Delete event (gateway.AutoModerationRuleDelete)
type autoModerationRuleDeleteEventHandler func(*Client, *AutoModerationRuleDelete)

// Handle handles the Auto Moderation Rule Delete event (gateway.AutoModerationRuleDelete)
func (h autoModerationRuleDeleteEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*AutoModerationRuleDelete); ok {
		h(c, e)
	}
}

// New returns the Auto Moderation Rule Delete event (gateway.AutoModerationRuleDelete)
func (h autoModerationRuleDeleteEventHandler) New() interface{} {
	return &AutoModerationRuleDelete{}
}

// autoModerationActionExecutionEventHandler is the event handler for the Auto Moderation Action Execution event (gateway.AutoModerationActionExecution)
type autoModerationActionExecutionEventHandler func(*Client, *AutoModerationActionExecution)

// Handle handles the Auto Moderation Action Execution event (gateway.AutoModerationActionExecution)
func (h autoModerationActionExecutionEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*AutoModerationActionExecution); ok {
		h(c, e)
	}
}

// New returns the Auto Moderation Action Execution event (gateway.AutoModerationActionExecution)
func (h autoModerationActionExecutionEventHandler) New() interface{} {
	return &AutoModerationActionExecution{}
}

// channelCreateEventHandler is the event handler for the Channel Create event (gateway.ChannelCreate)
type channelCreateEventHandler func(*Client, *ChannelCreate)

// Handle handles the Channel Create event (gateway.ChannelCreate)
func (h channelCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*ChannelCreate); ok {
		h(c, e)
	}
}

// New returns the Channel Create event (gateway.ChannelCreate)
func (h channelCreateEventHandler) New() interface{} {
	return &ChannelCreate{}
}

// channelUpdateEventHandler is the event handler for the Channel Update event (gateway.ChannelUpdate)
type channelUpdateEventHandler func(*Client, *ChannelUpdate)

// Handle handles the Channel Update event (gateway.ChannelUpdate)
func (h channelUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*ChannelUpdate); ok {
		h(c, e)
	}
}

// New returns the Channel Update event (gateway.ChannelUpdate)
func (h channelUpdateEventHandler) New() interface{} {
	return &ChannelUpdate{}
}

// channelDeleteEventHandler is the event handler for the Channel Delete event (gateway.ChannelDelete)
type channelDeleteEventHandler func(*Client, *ChannelDelete)

// Handle handles the Channel Delete event (gateway.ChannelDelete)
func (h channelDeleteEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*ChannelDelete); ok {
		h(c, e)
	}
}

// New returns the Channel Delete event (gateway.ChannelDelete)
func (h channelDeleteEventHandler) New() interface{} {
	return &ChannelDelete{}
}

// channelPinsUpdateEventHandler is the event handler for the Channel Pins Update event (gateway.ChannelPinsUpdate)
type channelPinsUpdateEventHandler func(*Client, *ChannelPinsUpdate)

// Handle handles the Channel Pins Update event (gateway.ChannelPinsUpdate)
func (h channelPinsUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*ChannelPinsUpdate); ok {
		h(c, e)
	}
}

// New returns the Channel Pins Update event (gateway.ChannelPinsUpdate)
func (h channelPinsUpdateEventHandler) New() interface{} {
	return &ChannelPinsUpdate{}
}

// threadCreateEventHandler is the event handler for the Thread Create event (gateway.ThreadCreate)
type threadCreateEventHandler func(*Client, *ThreadCreate)

// Handle handles the Thread Create event (gateway.ThreadCreate)
func (h threadCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*ThreadCreate); ok {
		h(c, e)
	}
}

// New returns the Thread Create event (gateway.ThreadCreate)
func (h threadCreateEventHandler) New() interface{} {
	return &ThreadCreate{}
}

// threadUpdateEventHandler is the event handler for the Thread Update event (gateway.ThreadUpdate)
type threadUpdateEventHandler func(*Client, *ThreadUpdate)

// Handle handles the Thread Update event (gateway.ThreadUpdate)
func (h threadUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*ThreadUpdate); ok {
		h(c, e)
	}
}

// New returns the Thread Update event (gateway.ThreadUpdate)
func (h threadUpdateEventHandler) New() interface{} {
	return &ThreadUpdate{}
}

// threadDeleteEventHandler is the event handler for the Thread Delete event (gateway.ThreadDelete)
type threadDeleteEventHandler func(*Client, *ThreadDelete)

// Handle handles the Thread Delete event (gateway.ThreadDelete)
func (h threadDeleteEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*ThreadDelete); ok {
		h(c, e)
	}
}

// New returns the Thread Delete event (gateway.ThreadDelete)
func (h threadDeleteEventHandler) New() interface{} {
	return &ThreadDelete{}
}

// threadListSyncEventHandler is the event handler for the Thread List Sync event (gateway.ThreadListSync)
type threadListSyncEventHandler func(*Client, *ThreadListSync)

// Handle handles the Thread List Sync event (gateway.ThreadListSync)
func (h threadListSyncEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*ThreadListSync); ok {
		h(c, e)
	}
}

// New returns the Thread List Sync event (gateway.ThreadListSync)
func (h threadListSyncEventHandler) New() interface{} {
	return &ThreadListSync{}
}

// threadMemberUpdateEventHandler is the event handler for the Thread Member Update event (gateway.ThreadMemberUpdate)
type threadMemberUpdateEventHandler func(*Client, *ThreadMemberUpdate)

// Handle handles the Thread Member Update event (gateway.ThreadMemberUpdate)
func (h threadMemberUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*ThreadMemberUpdate); ok {
		h(c, e)
	}
}

// New returns the Thread Member Update event (gateway.ThreadMemberUpdate)
func (h threadMemberUpdateEventHandler) New() interface{} {
	return &ThreadMemberUpdate{}
}

// threadMembersUpdateEventHandler is the event handler for the Thread Members Update event (gateway.ThreadMembersUpdate)
type threadMembersUpdateEventHandler func(*Client, *ThreadMembersUpdate)

// Handle handles the Thread Members Update event (gateway.ThreadMembersUpdate)
func (h threadMembersUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*ThreadMembersUpdate); ok {
		h(c, e)
	}
}

// New returns the Thread Members Update event (gateway.ThreadMembersUpdate)
func (h threadMembersUpdateEventHandler) New() interface{} {
	return &ThreadMembersUpdate{}
}

// guildCreateEventHandler is the event handler for the Guild Create event (gateway.GuildCreate)
type guildCreateEventHandler func(*Client, *GuildCreate)

// Handle handles the Guild Create event (gateway.GuildCreate)
func (h guildCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildCreate); ok {
		h(c, e)
	}
}

// New returns the Guild Create event (gateway.GuildCreate)
func (h guildCreateEventHandler) New() interface{} {
	return &GuildCreate{}
}

// guildUpdateEventHandler is the event handler for the Guild Update event (gateway.GuildUpdate)
type guildUpdateEventHandler func(*Client, *GuildUpdate)

// Handle handles the Guild Update event (gateway.GuildUpdate)
func (h guildUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildUpdate); ok {
		h(c, e)
	}
}

// New returns the Guild Update event (gateway.GuildUpdate)
func (h guildUpdateEventHandler) New() interface{} {
	return &GuildUpdate{}
}

// guildDeleteEventHandler is the event handler for the Guild Delete event (gateway.GuildDelete)
type guildDeleteEventHandler func(*Client, *GuildDelete)

// Handle handles the Guild Delete event (gateway.GuildDelete)
func (h guildDeleteEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildDelete); ok {
		h(c, e)
	}
}

// New returns the Guild Delete event (gateway.GuildDelete)
func (h guildDeleteEventHandler) New() interface{} {
	return &GuildDelete{}
}

// guildAuditLogEntryCreateEventHandler is the event handler for the Guild Audit Log Entry Create event (gateway.GuildAuditLogEntryCreate)
type guildAuditLogEntryCreateEventHandler func(*Client, *GuildAuditLogEntryCreate)

// Handle handles the Guild Audit Log Entry Create event (gateway.GuildAuditLogEntryCreate)
func (h guildAuditLogEntryCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildAuditLogEntryCreate); ok {
		h(c, e)
	}
}

// New returns the Guild Audit Log Entry Create event (gateway.GuildAuditLogEntryCreate)
func (h guildAuditLogEntryCreateEventHandler) New() interface{} {
	return &GuildAuditLogEntryCreate{}
}

// guildBanAddEventHandler is the event handler for the Guild Ban Add event (gateway.GuildBanAdd)
type guildBanAddEventHandler func(*Client, *GuildBanAdd)

// Handle handles the Guild Ban Add event (gateway.GuildBanAdd)
func (h guildBanAddEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildBanAdd); ok {
		h(c, e)
	}
}

// New returns the Guild Ban Add event (gateway.GuildBanAdd)
func (h guildBanAddEventHandler) New() interface{} {
	return &GuildBanAdd{}
}

// guildBanRemoveEventHandler is the event handler for the Guild Ban Remove event (gateway.GuildBanRemove)
type guildBanRemoveEventHandler func(*Client, *GuildBanRemove)

// Handle handles the Guild Ban Remove event (gateway.GuildBanRemove)
func (h guildBanRemoveEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildBanRemove); ok {
		h(c, e)
	}
}

// New returns the Guild Ban Remove event (gateway.GuildBanRemove)
func (h guildBanRemoveEventHandler) New() interface{} {
	return &GuildBanRemove{}
}

// guildEmojisUpdateEventHandler is the event handler for the Guild Emojis Update event (gateway.GuildEmojisUpdate)
type guildEmojisUpdateEventHandler func(*Client, *GuildEmojisUpdate)

// Handle handles the Guild Emojis Update event (gateway.GuildEmojisUpdate)
func (h guildEmojisUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildEmojisUpdate); ok {
		h(c, e)
	}
}

// New returns the Guild Emojis Update event (gateway.GuildEmojisUpdate)
func (h guildEmojisUpdateEventHandler) New() interface{} {
	return &GuildEmojisUpdate{}
}

// guildStickersUpdateEventHandler is the event handler for the Guild Stickers Update event (gateway.GuildStickersUpdate)
type guildStickersUpdateEventHandler func(*Client, *GuildStickersUpdate)

// Handle handles the Guild Stickers Update event (gateway.GuildStickersUpdate)
func (h guildStickersUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildStickersUpdate); ok {
		h(c, e)
	}
}

// New returns the Guild Stickers Update event (gateway.GuildStickersUpdate)
func (h guildStickersUpdateEventHandler) New() interface{} {
	return &GuildStickersUpdate{}
}

// guildIntegrationsUpdateEventHandler is the event handler for the Guild Integrations Update event (gateway.GuildIntegrationsUpdate)
type guildIntegrationsUpdateEventHandler func(*Client, *GuildIntegrationsUpdate)

// Handle handles the Guild Integrations Update event (gateway.GuildIntegrationsUpdate)
func (h guildIntegrationsUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildIntegrationsUpdate); ok {
		h(c, e)
	}
}

// New returns the Guild Integrations Update event (gateway.GuildIntegrationsUpdate)
func (h guildIntegrationsUpdateEventHandler) New() interface{} {
	return &GuildIntegrationsUpdate{}
}

// guildMemberAddEventHandler is the event handler for the Guild Member Add event (gateway.GuildMemberAdd)
type guildMemberAddEventHandler func(*Client, *GuildMemberAdd)

// Handle handles the Guild Member Add event (gateway.GuildMemberAdd)
func (h guildMemberAddEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildMemberAdd); ok {
		h(c, e)
	}
}

// New returns the Guild Member Add event (gateway.GuildMemberAdd)
func (h guildMemberAddEventHandler) New() interface{} {
	return &GuildMemberAdd{}
}

// guildMemberRemoveEventHandler is the event handler for the Guild Member Remove event (gateway.GuildMemberRemove)
type guildMemberRemoveEventHandler func(*Client, *GuildMemberRemove)

// Handle handles the Guild Member Remove event (gateway.GuildMemberRemove)
func (h guildMemberRemoveEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildMemberRemove); ok {
		h(c, e)
	}
}

// New returns the Guild Member Remove event (gateway.GuildMemberRemove)
func (h guildMemberRemoveEventHandler) New() interface{} {
	return &GuildMemberRemove{}
}

// guildMemberUpdateEventHandler is the event handler for the Guild Member Update event (gateway.GuildMemberUpdate)
type guildMemberUpdateEventHandler func(*Client, *GuildMemberUpdate)

// Handle handles the Guild Member Update event (gateway.GuildMemberUpdate)
func (h guildMemberUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildMemberUpdate); ok {
		h(c, e)
	}
}

// New returns the Guild Member Update event (gateway.GuildMemberUpdate)
func (h guildMemberUpdateEventHandler) New() interface{} {
	return &GuildMemberUpdate{}
}

// guildMembersChunkEventHandler is the event handler for the Guild Members Chunk event (gateway.GuildMembersChunk)
type guildMembersChunkEventHandler func(*Client, *GuildMembersChunk)

// Handle handles the Guild Members Chunk event (gateway.GuildMembersChunk)
func (h guildMembersChunkEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildMembersChunk); ok {
		h(c, e)
	}
}

// New returns the Guild Members Chunk event (gateway.GuildMembersChunk)
func (h guildMembersChunkEventHandler) New() interface{} {
	return &GuildMembersChunk{}
}

// guildRoleCreateEventHandler is the event handler for the Guild Role Create event (gateway.GuildRoleCreate)
type guildRoleCreateEventHandler func(*Client, *GuildRoleCreate)

// Handle handles the Guild Role Create event (gateway.GuildRoleCreate)
func (h guildRoleCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildRoleCreate); ok {
		h(c, e)
	}
}

// New returns the Guild Role Create event (gateway.GuildRoleCreate)
func (h guildRoleCreateEventHandler) New() interface{} {
	return &GuildRoleCreate{}
}

// guildRoleUpdateEventHandler is the event handler for the Guild Role Update event (gateway.GuildRoleUpdate)
type guildRoleUpdateEventHandler func(*Client, *GuildRoleUpdate)

// Handle handles the Guild Role Update event (gateway.GuildRoleUpdate)
func (h guildRoleUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildRoleUpdate); ok {
		h(c, e)
	}
}

// New returns the Guild Role Update event (gateway.GuildRoleUpdate)
func (h guildRoleUpdateEventHandler) New() interface{} {
	return &GuildRoleUpdate{}
}

// guildRoleDeleteEventHandler is the event handler for the Guild Role Delete event (gateway.GuildRoleDelete)
type guildRoleDeleteEventHandler func(*Client, *GuildRoleDelete)

// Handle handles the Guild Role Delete event (gateway.GuildRoleDelete)
func (h guildRoleDeleteEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildRoleDelete); ok {
		h(c, e)
	}
}

// New returns the Guild Role Delete event (gateway.GuildRoleDelete)
func (h guildRoleDeleteEventHandler) New() interface{} {
	return &GuildRoleDelete{}
}

// guildScheduledEventCreateEventHandler is the event handler for the Guild Scheduled Event Create event (gateway.GuildScheduledEventCreate)
type guildScheduledEventCreateEventHandler func(*Client, *GuildScheduledEventCreate)

// Handle handles the Guild Scheduled Event Create event (gateway.GuildScheduledEventCreate)
func (h guildScheduledEventCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildScheduledEventCreate); ok {
		h(c, e)
	}
}

// New returns the Guild Scheduled Event Create event (gateway.GuildScheduledEventCreate)
func (h guildScheduledEventCreateEventHandler) New() interface{} {
	return &GuildScheduledEventCreate{}
}

// guildScheduledEventUpdateEventHandler is the event handler for the Guild Scheduled Event Update event (gateway.GuildScheduledEventUpdate)
type guildScheduledEventUpdateEventHandler func(*Client, *GuildScheduledEventUpdate)

// Handle handles the Guild Scheduled Event Update event (gateway.GuildScheduledEventUpdate)
func (h guildScheduledEventUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildScheduledEventUpdate); ok {
		h(c, e)
	}
}

// New returns the Guild Scheduled Event Update event (gateway.GuildScheduledEventUpdate)
func (h guildScheduledEventUpdateEventHandler) New() interface{} {
	return &GuildScheduledEventUpdate{}
}

// guildScheduledEventDeleteEventHandler is the event handler for the Guild Scheduled Event Delete event (gateway.GuildScheduledEventDelete)
type guildScheduledEventDeleteEventHandler func(*Client, *GuildScheduledEventDelete)

// Handle handles the Guild Scheduled Event Delete event (gateway.GuildScheduledEventDelete)
func (h guildScheduledEventDeleteEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildScheduledEventDelete); ok {
		h(c, e)
	}
}

// New returns the Guild Scheduled Event Delete event (gateway.GuildScheduledEventDelete)
func (h guildScheduledEventDeleteEventHandler) New() interface{} {
	return &GuildScheduledEventDelete{}
}

// guildScheduledEventUserAddEventHandler is the event handler for the Guild Scheduled Event User Add event (gateway.GuildScheduledEventUserAdd)
type guildScheduledEventUserAddEventHandler func(*Client, *GuildScheduledEventUserAdd)

// Handle handles the Guild Scheduled Event User Add event (gateway.GuildScheduledEventUserAdd)
func (h guildScheduledEventUserAddEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildScheduledEventUserAdd); ok {
		h(c, e)
	}
}

// New returns the Guild Scheduled Event User Add event (gateway.GuildScheduledEventUserAdd)
func (h guildScheduledEventUserAddEventHandler) New() interface{} {
	return &GuildScheduledEventUserAdd{}
}

// guildScheduledEventUserRemoveEventHandler is the event handler for the Guild Scheduled Event User Remove event (gateway.GuildScheduledEventUserRemove)
type guildScheduledEventUserRemoveEventHandler func(*Client, *GuildScheduledEventUserRemove)

// Handle handles the Guild Scheduled Event User Remove event (gateway.GuildScheduledEventUserRemove)
func (h guildScheduledEventUserRemoveEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*GuildScheduledEventUserRemove); ok {
		h(c, e)
	}
}

// New returns the Guild Scheduled Event User Remove event (gateway.GuildScheduledEventUserRemove)
func (h guildScheduledEventUserRemoveEventHandler) New() interface{} {
	return &GuildScheduledEventUserRemove{}
}

// integrationCreateEventHandler is the event handler for the Integration Create event (gateway.IntegrationCreate)
type integrationCreateEventHandler func(*Client, *IntegrationCreate)

// Handle handles the Integration Create event (gateway.IntegrationCreate)
func (h integrationCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*IntegrationCreate); ok {
		h(c, e)
	}
}

// New returns the Integration Create event (gateway.IntegrationCreate)
func (h integrationCreateEventHandler) New() interface{} {
	return &IntegrationCreate{}
}

// integrationUpdateEventHandler is the event handler for the Integration Update event (gateway.IntegrationUpdate)
type integrationUpdateEventHandler func(*Client, *IntegrationUpdate)

// Handle handles the Integration Update event (gateway.IntegrationUpdate)
func (h integrationUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*IntegrationUpdate); ok {
		h(c, e)
	}
}

// New returns the Integration Update event (gateway.IntegrationUpdate)
func (h integrationUpdateEventHandler) New() interface{} {
	return &IntegrationUpdate{}
}

// integrationDeleteEventHandler is the event handler for the Integration Delete event (gateway.IntegrationDelete)
type integrationDeleteEventHandler func(*Client, *IntegrationDelete)

// Handle handles the Integration Delete event (gateway.IntegrationDelete)
func (h integrationDeleteEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*IntegrationDelete); ok {
		h(c, e)
	}
}

// New returns the Integration Delete event (gateway.IntegrationDelete)
func (h integrationDeleteEventHandler) New() interface{} {
	return &IntegrationDelete{}
}

// interactionCreateEventHandler is the event handler for the Interaction Create event (gateway.InteractionCreate)
type interactionCreateEventHandler func(*Client, *InteractionCreate)

// Handle handles the Interaction Create event (gateway.InteractionCreate)
func (h interactionCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*InteractionCreate); ok {
		h(c, e)
	}
}

// New returns the Interaction Create event (gateway.InteractionCreate)
func (h interactionCreateEventHandler) New() interface{} {
	return &InteractionCreate{}
}

// inviteCreateEventHandler is the event handler for the Invite Create event (gateway.InviteCreate)
type inviteCreateEventHandler func(*Client, *InviteCreate)

// Handle handles the Invite Create event (gateway.InviteCreate)
func (h inviteCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*InviteCreate); ok {
		h(c, e)
	}
}

// New returns the Invite Create event (gateway.InviteCreate)
func (h inviteCreateEventHandler) New() interface{} {
	return &InviteCreate{}
}

// inviteDeleteEventHandler is the event handler for the Invite Delete event (gateway.InviteDelete)
type inviteDeleteEventHandler func(*Client, *InviteDelete)

// Handle handles the Invite Delete event (gateway.InviteDelete)
func (h inviteDeleteEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*InviteDelete); ok {
		h(c, e)
	}
}

// New returns the Invite Delete event (gateway.InviteDelete)
func (h inviteDeleteEventHandler) New() interface{} {
	return &InviteDelete{}
}

// messageUpdateEventHandler is the event handler for the Message Update event (gateway.MessageUpdate)
type messageUpdateEventHandler func(*Client, *MessageUpdate)

// Handle handles the Message Update event (gateway.MessageUpdate)
func (h messageUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*MessageUpdate); ok {
		h(c, e)
	}
}

// New returns the Message Update event (gateway.MessageUpdate)
func (h messageUpdateEventHandler) New() interface{} {
	return &MessageUpdate{}
}

// messageDeleteEventHandler is the event handler for the Message Delete event (gateway.MessageDelete)
type messageDeleteEventHandler func(*Client, *MessageDelete)

// Handle handles the Message Delete event (gateway.MessageDelete)
func (h messageDeleteEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*MessageDelete); ok {
		h(c, e)
	}
}

// New returns the Message Delete event (gateway.MessageDelete)
func (h messageDeleteEventHandler) New() interface{} {
	return &MessageDelete{}
}

// messageDeleteBulkEventHandler is the event handler for the Message Delete Bulk event (gateway.MessageDeleteBulk)
type messageDeleteBulkEventHandler func(*Client, *MessageDeleteBulk)

// Handle handles the Message Delete Bulk event (gateway.MessageDeleteBulk)
func (h messageDeleteBulkEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*MessageDeleteBulk); ok {
		h(c, e)
	}
}

// New returns the Message Delete Bulk event (gateway.MessageDeleteBulk)
func (h messageDeleteBulkEventHandler) New() interface{} {
	return &MessageDeleteBulk{}
}

// messageReactionAddEventHandler is the event handler for the Message Reaction Add event (gateway.MessageReactionAdd)
type messageReactionAddEventHandler func(*Client, *MessageReactionAdd)

// Handle handles the Message Reaction Add event (gateway.MessageReactionAdd)
func (h messageReactionAddEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*MessageReactionAdd); ok {
		h(c, e)
	}
}

// New returns the Message Reaction Add event (gateway.MessageReactionAdd)
func (h messageReactionAddEventHandler) New() interface{} {
	return &MessageReactionAdd{}
}

// messageReactionRemoveEventHandler is the event handler for the Message Reaction Remove event (gateway.MessageReactionRemove)
type messageReactionRemoveEventHandler func(*Client, *MessageReactionRemove)

// Handle handles the Message Reaction Remove event (gateway.MessageReactionRemove)
func (h messageReactionRemoveEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*MessageReactionRemove); ok {
		h(c, e)
	}
}

// New returns the Message Reaction Remove event (gateway.MessageReactionRemove)
func (h messageReactionRemoveEventHandler) New() interface{} {
	return &MessageReactionRemove{}
}

// messageReactionRemoveAllEventHandler is the event handler for the Message Reaction Remove All event (gateway.MessageReactionRemoveAll)
type messageReactionRemoveAllEventHandler func(*Client, *MessageReactionRemoveAll)

// Handle handles the Message Reaction Remove All event (gateway.MessageReactionRemoveAll)
func (h messageReactionRemoveAllEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*MessageReactionRemoveAll); ok {
		h(c, e)
	}
}

// New returns the Message Reaction Remove All event (gateway.MessageReactionRemoveAll)
func (h messageReactionRemoveAllEventHandler) New() interface{} {
	return &MessageReactionRemoveAll{}
}

// messageReactionRemoveEmojiEventHandler is the event handler for the Message Reaction Remove Emoji event (gateway.MessageReactionRemoveEmoji)
type messageReactionRemoveEmojiEventHandler func(*Client, *MessageReactionRemoveEmoji)

// Handle handles the Message Reaction Remove Emoji event (gateway.MessageReactionRemoveEmoji)
func (h messageReactionRemoveEmojiEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*MessageReactionRemoveEmoji); ok {
		h(c, e)
	}
}

// New returns the Message Reaction Remove Emoji event (gateway.MessageReactionRemoveEmoji)
func (h messageReactionRemoveEmojiEventHandler) New() interface{} {
	return &MessageReactionRemoveEmoji{}
}

type presenceUpdateEventHandler func(*Client, *PresenceUpdate)

// Handle handles the Presence Update event (gateway.PresenceUpdate)
func (h presenceUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*PresenceUpdate); ok {
		h(c, e)
	}
}

// New returns the Presence Update event (gateway.PresenceUpdate)
func (h presenceUpdateEventHandler) New() interface{} {
	return &PresenceUpdate{}
}

// stageInstanceCreateEventHandler is the event handler for the Stage Instance Create event (gateway.StageInstanceCreate)
type stageInstanceCreateEventHandler func(*Client, *StageInstanceCreate)

// Handle handles the Stage Instance Create event (gateway.StageInstanceCreate)
func (h stageInstanceCreateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*StageInstanceCreate); ok {
		h(c, e)
	}
}

// New returns the Stage Instance Create event (gateway.StageInstanceCreate)
func (h stageInstanceCreateEventHandler) New() interface{} {
	return &StageInstanceCreate{}
}

// stageInstanceUpdateEventHandler is the event handler for the Stage Instance Update event (gateway.StageInstanceUpdate)
type stageInstanceUpdateEventHandler func(*Client, *StageInstanceUpdate)

// Handle handles the Stage Instance Update event (gateway.StageInstanceUpdate)
func (h stageInstanceUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*StageInstanceUpdate); ok {
		h(c, e)
	}
}

// New returns the Stage Instance Update event (gateway.StageInstanceUpdate)
func (h stageInstanceUpdateEventHandler) New() interface{} {
	return &StageInstanceUpdate{}
}

// stageInstanceDeleteEventHandler is the event handler for the Stage Instance Delete event (gateway.StageInstanceDelete)
type stageInstanceDeleteEventHandler func(*Client, *StageInstanceDelete)

// Handle handles the Stage Instance Delete event (gateway.StageInstanceDelete)
func (h stageInstanceDeleteEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*StageInstanceDelete); ok {
		h(c, e)
	}
}

// New returns the Stage Instance Delete event (gateway.StageInstanceDelete)
func (h stageInstanceDeleteEventHandler) New() interface{} {
	return &StageInstanceDelete{}
}

// typingStartEventHandler is the event handler for the Typing Start event (gateway.TypingStart)
type typingStartEventHandler func(*Client, *TypingStart)

// Handle handles the Typing Start event (gateway.TypingStart)
func (h typingStartEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*TypingStart); ok {
		h(c, e)
	}
}

// New returns the Typing Start event (gateway.TypingStart)
func (h typingStartEventHandler) New() interface{} {
	return &TypingStart{}
}

// userUpdateEventHandler is the event handler for the User Update event (gateway.UserUpdate)
type userUpdateEventHandler func(*Client, *UserUpdate)

// Handle handles the User Update event (gateway.UserUpdate)
func (h userUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*UserUpdate); ok {
		h(c, e)
	}
}

// New returns the User Update event (gateway.UserUpdate)
func (h userUpdateEventHandler) New() interface{} {
	return &UserUpdate{}
}

// voiceStateUpdateEventHandler is the event handler for the Voice State Update event (gateway.VoiceStateUpdate)
type voiceStateUpdateEventHandler func(*Client, *VoiceStateUpdate)

// Handle handles the Voice State Update event (gateway.VoiceStateUpdate)
func (h voiceStateUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*VoiceStateUpdate); ok {
		h(c, e)
	}
}

// New returns the Voice State Update event (gateway.VoiceStateUpdate)
func (h voiceStateUpdateEventHandler) New() interface{} {
	return &VoiceStateUpdate{}
}

// voiceServerUpdateEventHandler is the event handler for the Voice Server Update event (gateway.VoiceServerUpdate)
type voiceServerUpdateEventHandler func(*Client, *VoiceServerUpdate)

// Handle handles the Voice Server Update event (gateway.VoiceServerUpdate)
func (h voiceServerUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*VoiceServerUpdate); ok {
		h(c, e)
	}
}

// New returns the Voice Server Update event (gateway.VoiceServerUpdate)
func (h voiceServerUpdateEventHandler) New() interface{} {
	return &VoiceServerUpdate{}
}

// webhooksUpdateEventHandler is the event handler for the Webhooks Update event (gateway.WebhooksUpdate)
type webhooksUpdateEventHandler func(*Client, *WebhooksUpdate)

// Handle handles the Webhooks Update event (gateway.WebhooksUpdate)
func (h webhooksUpdateEventHandler) Handle(c *Client, i interface{}) {
	if e, ok := i.(*WebhooksUpdate); ok {
		h(c, e)
	}
}

// New returns the Webhooks Update event (gateway.WebhooksUpdate)
func (h webhooksUpdateEventHandler) New() interface{} {
	return &WebhooksUpdate{}
}

// handlerForFunction returns the event handler for the given function
func handlerForFunction(f interface{}) EventHandler {
	switch t := f.(type) {
	case func(*Client, *Ready):
		return readyEventHandler(t)
	case func(*Client, *ApplicationCommandPermissionsUpdate):
		return applicationCommandPermissionsUpdateEventHandler(t)
	case func(*Client, *AutoModerationRuleCreate):
		return autoModerationRuleCreateEventHandler(t)
	case func(*Client, *AutoModerationRuleUpdate):
		return autoModerationRuleUpdateEventHandler(t)
	case func(*Client, *AutoModerationRuleDelete):
		return autoModerationRuleDeleteEventHandler(t)
	case func(*Client, *AutoModerationActionExecution):
		return autoModerationActionExecutionEventHandler(t)
	case func(*Client, *ChannelCreate):
		return channelCreateEventHandler(t)
	case func(*Client, *ChannelUpdate):
		return channelUpdateEventHandler(t)
	case func(*Client, *ChannelDelete):
		return channelDeleteEventHandler(t)
	case func(*Client, *ChannelPinsUpdate):
		return channelPinsUpdateEventHandler(t)
	case func(*Client, *ThreadCreate):
		return threadCreateEventHandler(t)
	case func(*Client, *ThreadUpdate):
		return threadUpdateEventHandler(t)
	case func(*Client, *ThreadDelete):
		return threadDeleteEventHandler(t)
	case func(*Client, *ThreadListSync):
		return threadListSyncEventHandler(t)
	case func(*Client, *ThreadMemberUpdate):
		return threadMemberUpdateEventHandler(t)
	case func(*Client, *ThreadMembersUpdate):
		return threadMembersUpdateEventHandler(t)
	case func(*Client, *GuildCreate):
		return guildCreateEventHandler(t)
	case func(*Client, *GuildUpdate):
		return guildUpdateEventHandler(t)
	case func(*Client, *GuildDelete):
		return guildDeleteEventHandler(t)
	case func(*Client, *GuildAuditLogEntryCreate):
		return guildAuditLogEntryCreateEventHandler(t)
	case func(*Client, *GuildBanAdd):
		return guildBanAddEventHandler(t)
	case func(*Client, *GuildBanRemove):
		return guildBanRemoveEventHandler(t)
	case func(*Client, *GuildEmojisUpdate):
		return guildEmojisUpdateEventHandler(t)
	case func(*Client, *GuildStickersUpdate):
		return guildStickersUpdateEventHandler(t)
	case func(*Client, *GuildIntegrationsUpdate):
		return guildIntegrationsUpdateEventHandler(t)
	case func(*Client, *GuildMemberAdd):
		return guildMemberAddEventHandler(t)
	case func(*Client, *GuildMemberRemove):
		return guildMemberRemoveEventHandler(t)
	case func(*Client, *GuildMemberUpdate):
		return guildMemberUpdateEventHandler(t)
	case func(*Client, *GuildMembersChunk):
		return guildMembersChunkEventHandler(t)
	case func(*Client, *GuildRoleCreate):
		return guildRoleCreateEventHandler(t)
	case func(*Client, *GuildRoleUpdate):
		return guildRoleUpdateEventHandler(t)
	case func(*Client, *GuildRoleDelete):
		return guildRoleDeleteEventHandler(t)
	case func(*Client, *GuildScheduledEventCreate):
		return guildScheduledEventCreateEventHandler(t)
	case func(*Client, *GuildScheduledEventUpdate):
		return guildScheduledEventUpdateEventHandler(t)
	case func(*Client, *GuildScheduledEventDelete):
		return guildScheduledEventDeleteEventHandler(t)
	case func(*Client, *GuildScheduledEventUserAdd):
		return guildScheduledEventUserAddEventHandler(t)
	case func(*Client, *GuildScheduledEventUserRemove):
		return guildScheduledEventUserRemoveEventHandler(t)
	case func(*Client, *IntegrationCreate):
		return integrationCreateEventHandler(t)
	case func(*Client, *IntegrationUpdate):
		return integrationUpdateEventHandler(t)
	case func(*Client, *IntegrationDelete):
		return integrationDeleteEventHandler(t)
	case func(*Client, *InteractionCreate):
		return interactionCreateEventHandler(t)
	case func(*Client, *InviteCreate):
		return inviteCreateEventHandler(t)
	case func(*Client, *InviteDelete):
		return inviteDeleteEventHandler(t)
	case func(*Client, *MessageCreate):
		return messageCreateEventHandler(t)
	case func(*Client, *MessageUpdate):
		return messageUpdateEventHandler(t)
	case func(*Client, *MessageDelete):
		return messageDeleteEventHandler(t)
	case func(*Client, *MessageDeleteBulk):
		return messageDeleteBulkEventHandler(t)
	case func(*Client, *MessageReactionAdd):
		return messageReactionAddEventHandler(t)
	case func(*Client, *MessageReactionRemove):
		return messageReactionRemoveEventHandler(t)
	case func(*Client, *MessageReactionRemoveAll):
		return messageReactionRemoveAllEventHandler(t)
	case func(*Client, *MessageReactionRemoveEmoji):
		return messageReactionRemoveEmojiEventHandler(t)
	case func(*Client, *PresenceUpdate):
		return presenceUpdateEventHandler(t)
	case func(*Client, *StageInstanceCreate):
		return stageInstanceCreateEventHandler(t)
	case func(*Client, *StageInstanceUpdate):
		return stageInstanceUpdateEventHandler(t)
	case func(*Client, *StageInstanceDelete):
		return stageInstanceDeleteEventHandler(t)
	case func(*Client, *TypingStart):
		return typingStartEventHandler(t)
	case func(*Client, *UserUpdate):
		return userUpdateEventHandler(t)
	case func(*Client, *VoiceStateUpdate):
		return voiceStateUpdateEventHandler(t)
	case func(*Client, *VoiceServerUpdate):
		return voiceServerUpdateEventHandler(t)
	case func(*Client, *WebhooksUpdate):
		return webhooksUpdateEventHandler(t)
	default:
		return nil
	}
}
