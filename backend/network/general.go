package network

import (
	"github.com/Zenika/MARIE/backend/mqtt"
	"github.com/Zenika/MARIE/backend/thing"
)

// Do something on all things that match action and room
func Do(thingType string, action string, params map[string]interface{}, location string) {
	things := thing.ReadActionName(action)
	for _, t := range things {
		if t.Type == thingType {
			if location == "" || t.Location == location {
				DoSomething(t, action, params)
			}
		}
	}
}

// DoSomething on a thing
func DoSomething(t thing.Thing, action string, params map[string]interface{}) {
	switch t.Protocol {
	case "MQTT":
		mqtt.Do(t.MacAddress, action, params)
		break
	}
}
