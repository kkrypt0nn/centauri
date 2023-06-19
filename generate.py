import os
import sys

event_name = ""

if len(sys.argv) == 2:
    event_name = sys.argv[1]
else:
    print("Please give an event name")
    exit(1337)

os.system("cls")

print("In events.go:")
print(f"EventType{event_name.replace(' ', '')} EventType = \"{event_name.upper().replace(' ', '_')}\"")

print("\n" + "-"*25 + "\n")

print("In events_receive.go:")
print(f"// {event_name.replace(' ', '')} is the data of the {event_name} event")
print(f"// https://discord.com/developers/docs/topics/gateway-events#{event_name.lower().replace(' ', '-')}")
print(f"type {event_name.replace(' ', '')} struct {{")
print(f"\t// Stuff here")
print("}")

print("\n" + "-"*25 + "\n")

print("In event_handlers.go:")
event_handler = f"{event_name[0].lower() + event_name[1:].replace(' ', '')}EventHandler"
print(f"// {event_handler} is the event handler for the {event_name} event (gateway.{event_name.replace(' ', '')})")
print(f"type {event_handler} func(*Client, *{event_name.replace(' ', '')})")
print("")
print(f"// Handle handles the {event_name} event (gateway.{event_name.replace(' ', '')})")
print(f"func (h {event_handler}) Handle(c *Client, i interface{{}}) {{")
print(f"\tif e, ok := i.(*{event_name.replace(' ', '')}); ok {{")
print(f"\t\th(c, e)")
print("\t}")
print("}")
print("")
print(f"// New returns the {event_name} event (gateway.{event_name.replace(' ', '')})")
print(f"func (h {event_handler}) New() interface{{}} {{")
print(f"\treturn &{event_name.replace(' ', '')}{{}}")
print("}")
print("")
print(f"case func(*Client, *{event_name.replace(' ', '')}):")
print(f"\treturn {event_handler}(t)")

print("\n" + "-"*25 + "\n")

print("In gateway.go:")
print(f"EventType{event_name.replace(' ', '')}: {event_handler}(nil),")
