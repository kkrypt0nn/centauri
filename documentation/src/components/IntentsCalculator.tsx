import React, { ChangeEvent, useState } from "react";
import { Intent, intentsMap } from "../data/intents";
import "./IntentsCalculator.css";

export default function IntentsCalculator() {
  const [intentValue, setIntentValue] = useState(0);

  function handleCheckbox(event: ChangeEvent<HTMLInputElement>) {
    const intent: Intent = intentsMap.get(event.target.id);
    setIntentValue(intentValue ^ (1 << intent.shift));
  }

  return (
    <>
      <div className="padding-top--lg">
        <div>
          <span className="intent-value__badge">Intents: {intentValue}</span>
        </div>
        <div className="intents__table padding-top--md">
          <div className="intents__column intents__list">
            <h3>Intents you choose</h3>
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
              {Array.from(intentsMap.keys()).map((key: string) => {
                const intent: Intent = intentsMap.get(key);
                if ((intentValue & (1 << intent.shift)) !== 1 << intent.shift) {
                  return;
                }
                return intentsMap.get(key).events.map((event: string) => {
                  return (
                    <li>
                      <code>{event}</code>
                    </li>
                  );
                });
              })}
            </ul>
          </div>
        </div>
      </div>
    </>
  );
}
