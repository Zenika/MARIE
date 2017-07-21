package network

import (
	"github.com/Zenika/MARIE/backend/apiai"
	"github.com/Zenika/MARIE/backend/thing"
	uuid "github.com/satori/go.uuid"
)

// Do something on all things that match action and room
func Do(thingType string, action string, params map[string]interface{}, location string) {
	things := thing.ReadActionName(action)
	for _, t := range things {
		if t.Type == thingType {
			if location == "" || t.Location == location {
				switch t.Protocol {
				case "MQTT":
					mqttConn.Do(t.MacAddress, action, params)
					break
				}
			}
		}
	}
}

// Analyze query and returns response
func Analyze(req string) map[string]interface{} {
	res := apiai.Analyze(req)
	if res.Metadata.IntentName == "Get" {
		id := uuid.NewV4()
		return map[string]interface{}{"executing": id, "count": Get(id.String(), res.Parameters["variable-name"], res.Parameters["location"])}
	}

	if res.Metadata.IntentName == "Do" {
		Do(res.Parameters["thing"], res.Parameters["action"], nil, res.Parameters["location"])
		return map[string]interface{}{"doing": res.Parameters["action"], "on": res.Parameters["thing"], "in": res.Parameters["location"]}
	}
	return nil
}

// Broadcast message to all connected devices
func Broadcast(m []byte) {
	BroadcastWS(m)
}

// Get some value
func Get(id string, name string, location string) int {
	things := thing.ReadGetterName(name)
	sum := 0
	for _, t := range things {
		if location == "" || location == t.Location {
			sum = sum + 1
			switch t.Protocol {
			case "MQTT":
				mqttConn.Get(id, name, t.MacAddress)
				break
			}
		}
	}
	return sum
}
