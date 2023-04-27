import React, { ChangeEvent, useState } from "react";
import { Intent, defaultEvents, intentsMap } from "../../data/intents";
import "../IntentsCalculator.css";

export default function IntentsCalculatorIntents() {
  const [intentValue, setIntentValue] = useState(0);

  function handleCheckbox(event: ChangeEvent<HTMLInputElement>) {
    const intent: Intent = intentsMap.get(event.target.id);
    setIntentValue(intentValue ^ (1 << intent.shift));
  }

  return (
    <>
      <div>
        <span className="intent-value__badge">Intents: {intentValue}</span>
      </div>
      <div className="intents__table padding-top--md">
        <div className="intents__column intents__list">
          <h3>Intents you need</h3>
          {Array.from(intentsMap.keys()).map((key: string) => {
            const intent: Intent = intentsMap.get(key);
            return (
              <div key={key}>
                <label
                  className={
                    "intent__label" +
                    (intent.privileged ? " intent-privileged__label" : "")
                  }
                >
                  <input type="checkbox" id={key} onChange={handleCheckbox} />
                  {intent.name} (1 &lt;&lt; {intent.shift})
                </label>
              </div>
            );
          })}
        </div>
        <div className="intents__column">
          <h3>Gateway events you will receive</h3>
          <ul>
            {defaultEvents.map((event: string) => {
              return (
                <li key={event}>
                  <code className="event-default__label">{event}</code>
                </li>
              );
            })}
            {Array.from(intentsMap.keys()).map((key: string) => {
              const intent: Intent = intentsMap.get(key);
              const flag: number = 1 << intent.shift;
              if ((intentValue & flag) !== flag) {
                return;
              }
              return intent.events.map((event: string) => {
                return (
                  <li key={event}>
                    <code
                      className={
                        intent.privileged ? "intent-privileged__label" : ""
                      }
                    >
                      {event}
                    </code>
                  </li>
                );
              });
            })}
          </ul>
        </div>
      </div>
    </>
  );
}
