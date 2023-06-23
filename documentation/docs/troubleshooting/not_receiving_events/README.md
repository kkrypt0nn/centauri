---
title: Not receiving events
description: If you do not set the intents correctly when using the library, you will not receive events you expeted to receive.
---

## Missing intents

If you do not set the intents correctly when using the library, you will not receive events you expeted to receive.

Make sure that whenever you've created your Gateway Client you've specified the right intents. You can read about intents [here](/docs/guides/intents) and on how to use them properly. Here is an example to receive events when messages are sent in guilds and have access to their content:

```go
botClient := centauri.NewGatewayClient("Bot BOT_TOKEN", discord.IntentsGuildMessages|discord.IntentsMessageContent)
```

## More issues

If the explanations from this page do not help you with your issue(s), please join our [Discord server](https://discord.gg/feA6ZGRgpw) to get further assisted help.
