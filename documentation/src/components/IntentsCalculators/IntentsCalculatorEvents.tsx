import React, { ChangeEvent, useState } from "react";
import { Intent, defaultEvents, intentsMap } from "../../data/intents";
import "../IntentsCalculator.css";

export default function IntentsCalculatorEvents() {
  const [intentValue, setIntentValue] = useState(0);
  const [intents, setIntents] = useState(
    new Map<string, number>([
      ["GUILDS", 0],
      ["GUILD_MEMBERS", 0],
      ["GUILD_MODERATION", 0],
      ["GUILD_EMOJIS_AND_STICKERS", 0],
      ["GUILD_INTEGRATIONS", 0],
      ["GUILD_WEBHOOKS", 0],
      ["GUILD_INVITES", 0],
      ["GUILD_VOICE_STATES", 0],
      ["GUILD_PRESENCES", 0],
      ["GUILD_MESSAGES", 0],
      ["GUILD_MESSAGE_REACTIONS", 0],
      ["GUILD_MESSAGE_TYPING", 0],
      ["DIRECT_MESSAGES", 0],
      ["DIRECT_MESSAGE_REACTIONS", 0],
      ["DIRECT_MESSAGE_TYPING", 0],
      ["MESSAGE_CONTENT", 0],
      ["GUILD_SCHEDULED_EVENTS", 0],
      ["AUTO_MODERATION_CONFIGURATION", 0],
      ["AUTO_MODERATION_EXECUTION", 0],
    ])
  );

  function handleCheckbox(event: ChangeEvent<HTMLInputElement>) {
    const eventName: string = event.target.id;
    const checked: boolean = event.target.checked;

    let intent: Intent = undefined;
    Array.from(intentsMap.keys()).map((key: string) => {
      const iIntent: Intent = intentsMap.get(key);
      if (iIntent.events.some((event) => event === eventName)) {
        intent = iIntent;
      }
    });
    if (intent === undefined) {
      return;
    }

    const totalEventsNeeded = intents.get(intent.name);
    if (totalEventsNeeded === 0 || totalEventsNeeded === 1) {
      setIntentValue(intentValue ^ (1 << intent.shift));
    }
    intents.set(
      intent.name,
      checked ? totalEventsNeeded + 1 : totalEventsNeeded - 1
    );
  }

  return (
    <>
      <div>
        <span className="intent-value__badge">Intents: {intentValue}</span>
      </div>
      <div className="intents__table padding-top--md">
        <div className="intents__column intents__list">
          <h3>Gateway events you need</h3>
          {defaultEvents.map((event: string) => {
            return (
              <div key={event}>
                <label className="intent__label event-default__label">
                  <input
                    type="checkbox"
                    id={event}
                    checked={true}
                    disabled={true}
                  />
                  {event}
                </label>
              </div>
            );
          })}
          {Array.from(intentsMap.keys()).map((key: string) => {
            const intent: Intent = intentsMap.get(key);
            return intent.events.map((event: string) => {
              return (
                <div key={event}>
                  <label
                    className={
                      "intent__label" +
                      (intent.privileged ? " intent-privileged__label" : "")
                    }
                  >
                    <input
                      type="checkbox"
                      id={event}
                      onChange={handleCheckbox}
                    />
                    {event}
                  </label>
                </div>
              );
            });
          })}
        </div>
        <div className="intents__column">
          <h3>Intents you will need</h3>
          {intentValue === 0 ? (
            <p>
              <i>None</i>
            </p>
          ) : (
            <ul>
              {Array.from(intents.keys()).map((key: string) => {
                const intent: Intent = intentsMap.get(key);
                const flag: number = 1 << intent.shift;
                if ((intentValue & flag) !== flag) {
                  return;
                }
                return (
                  <li>
                    <code
                      className={
                        intent.privileged ? "intent-privileged__label" : ""
                      }
                    >
                      {intent.name} (1 &lt;&lt; {intent.shift})
                    </code>
                  </li>
                );
              })}
            </ul>
          )}
        </div>
      </div>
    </>
  );
}
