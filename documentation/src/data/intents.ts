export type Intent = {
  name: string;
  privileged: boolean;
  shift: number;
  events: string[];
};

export const intentsMap = new Map<string, Intent>([
  [
    "GUILDS",
    {
      name: "GUILDS",
      privileged: false,
      shift: 0,
      events: [
        "GUILD_CREATE",
        "GUILD_UPDATE",
        "GUILD_DELETE",
        "GUILD_ROLE_CREATE",
        "GUILD_ROLE_UPDATE",
        "GUILD_ROLE_DELETE",
        "CHANNEL_CREATE",
        "CHANNEL_UPDATE",
        "CHANNEL_DELETE",
        "CHANNEL_PINS_UPDATE",
        "THREAD_CREATE",
        "THREAD_UPDATE",
        "THREAD_DELETE",
        "THREAD_LIST_SYNC",
        "THREAD_MEMBER_UPDATE",
        "THREAD_MEMBERS_UPDATE",
        "STAGE_INSTANCE_CREATE",
        "STAGE_INSTANCE_UPDATE",
        "STAGE_INSTANCE_DELETE",
      ],
    },
  ],
  [
    "GUILD_MEMBERS",
    {
      name: "GUILD_MEMBERS",
      privileged: true,
      shift: 1,
      events: [
        "GUILD_MEMBER_ADD",
        "GUILD_MEMBER_UPDATE",
        "GUILD_MEMBER_REMOVE",
        "THREAD_MEMBERS_UPDATE",
      ],
    },
  ],
  [
    "GUILD_MODERATION",
    {
      name: "GUILD_MODERATION",
      privileged: false,
      shift: 2,
      events: [
        "GUILD_AUDIT_LOG_ENTRY_CREATE",
        "GUILD_BAN_ADD",
        "GUILD_BAN_REMOVE",
      ],
    },
  ],
  [
    "GUILD_EMOJIS_AND_STICKERS",
    {
      name: "GUILD_MEMBERS",
      privileged: false,
      shift: 3,
      events: ["GUILD_EMOJIS_UPDATE", "GUILD_STICKERS_UPDATE"],
    },
  ],
  [
    "GUILD_INTEGRATIONS",
    {
      name: "GUILD_INTEGRATIONS",
      privileged: false,
      shift: 4,
      events: [
        "GUILD_INTEGRATIONS_UPDATE",
        "INTEGRATION_CREATE",
        "INTEGRATION_UPDATE",
        "INTEGRATION_DELETE",
      ],
    },
  ],
  [
    "GUILD_WEBHOOKS",
    {
      name: "GUILD_WEBHOOKS",
      privileged: false,
      shift: 5,
      events: ["WEBHOOKS_UPDATE"],
    },
  ],
  [
    "GUILD_INVITES",
    {
      name: "GUILD_INVITES",
      privileged: false,
      shift: 6,
      events: ["INVITE_CREATE", "INVITE_DELETE"],
    },
  ],
  [
    "GUILD_VOICE_STATES",
    {
      name: "GUILD_VOICE_STATES",
      privileged: false,
      shift: 7,
      events: ["VOICE_STATE_UPDATE"],
    },
  ],
  [
    "GUILD_PRESENCES",
    {
      name: "GUILD_PRESENCES",
      privileged: true,
      shift: 8,
      events: ["PRESENCE_UPDATE"],
    },
  ],
  [
    "GUILD_MESSAGES",
    {
      name: "GUILD_MESSAGES",
      privileged: false,
      shift: 9,
      events: [
        "MESSAGE_CREATE",
        "MESSAGE_UPDATE",
        "MESSAGE_DELETE",
        "MESSAGE_DELETE_BULK",
      ],
    },
  ],
  [
    "GUILD_MESSAGE_REACTIONS",
    {
      name: "GUILD_MESSAGE_REACTIONS",
      privileged: false,
      shift: 10,
      events: [
        "MESSAGE_REACTION_ADD",
        "MESSAGE_REACTION_REMOVE",
        "MESSAGE_REACTION_REMOVE_ALL",
        "MESSAGE_REACTION_REMOVE_EMOJI",
      ],
    },
  ],
  [
    "GUILD_MESSAGE_TYPING",
    {
      name: "GUILD_MESSAGE_TYPING",
      privileged: false,
      shift: 11,
      events: ["TYPING_START"],
    },
  ],
  [
    "DIRECT_MESSAGES",
    {
      name: "DIRECT_MESSAGES",
      privileged: false,
      shift: 12,
      events: [
        "MESSAGE_CREATE",
        "MESSAGE_UPDATE",
        "MESSAGE_DELETE",
        "CHANNEL_PINS_UPDATE",
      ],
    },
  ],
  [
    "DIRECT_MESSAGE_REACTIONS",
    {
      name: "DIRECT_MESSAGE_REACTIONS",
      privileged: false,
      shift: 13,
      events: [
        "MESSAGE_REACTION_ADD",
        "MESSAGE_REACTION_REMOVE",
        "MESSAGE_REACTION_REMOVE_ALL",
        "MESSAGE_REACTION_REMOVE_EMOJI",
      ],
    },
  ],
  [
    "DIRECT_MESSAGE_TYPING",
    {
      name: "DIRECT_MESSAGE_TYPING",
      privileged: false,
      shift: 14,
      events: ["TYPING_START"],
    },
  ],
  [
    "MESSAGE_CONTENT",
    {
      name: "MESSAGE_CONTENT",
      privileged: true,
      shift: 15,
      events: [],
    },
  ],
  [
    "GUILD_SCHEDULED_EVENTS",
    {
      name: "GUILD_SCHEDULED_EVENTS",
      privileged: false,
      shift: 16,
      events: [
        "GUILD_SCHEDULED_EVENT_CREATE",
        "GUILD_SCHEDULED_EVENT_UPDATE",
        "GUILD_SCHEDULED_EVENT_DELETE",
        "GUILD_SCHEDULED_EVENT_USER_ADD",
        "GUILD_SCHEDULED_EVENT_USER_REMOVE",
      ],
    },
  ],
  [
    "AUTO_MODERATION_CONFIGURATION",
    {
      name: "AUTO_MODERATION_CONFIGURATION",
      privileged: false,
      shift: 20,
      events: [
        "AUTO_MODERATION_RULE_CREATE",
        "AUTO_MODERATION_RULE_UPDATE",
        "AUTO_MODERATION_RULE_DELETE",
      ],
    },
  ],
  [
    "AUTO_MODERATION_EXECUTION",
    {
      name: "AUTO_MODERATION_EXECUTION",
      privileged: false,
      shift: 21,
      events: ["AUTO_MODERATION_ACTION_EXECUTION"],
    },
  ],
]);
