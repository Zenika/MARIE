package network

import (
	"log"

	"github.com/Zenika/MARIE/backend/apiai"
	"github.com/Zenika/MARIE/backend/thing"
	uuid "github.com/satori/go.uuid"
)

// Do something on all things that match action and room
func Do(thingType string, action string, params map[string]interface{}, location string) (int, error) {
	things, err := thing.ReadActionName(action)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, t := range things {
		if t.Type == thingType {
			if location == "" || t.Location == location {
				sum = sum + 1
				switch t.Protocol {
				case "MQTT":
					mqttConn.Do(t.MacAddress, action, params)
					break
				}
			}
		}
	}
	return sum, nil
}

// Analyze query and returns response
func Analyze(req string) map[string]interface{} {
	res := apiai.Analyze(req)
	// If the user wants to Get some data
	if res.Metadata.IntentName == "Get" {
		id := uuid.NewV4()
		count, err := Get(id.String(), res.Parameters["variable-name"], res.Parameters["location"])
		if err != nil {
			return map[string]interface{}{"error": err.Error()}
		}
		return map[string]interface{}{
			"executing": id,
			"count":     count,
			"message":   res.Fulfillment.Speech,
		}
	}

	// If the user wants to Do something
	if res.Metadata.IntentName == "Do" {
		count, err := Do(res.Parameters["thing"], res.Parameters["action"], nil, res.Parameters["location"])
		if err != nil {
			return map[string]interface{}{"error": err.Error()}
		}
		return map[string]interface{}{
			"doing":   res.Parameters["action"],
			"on":      res.Parameters["thing"],
			"in":      res.Parameters["location"],
			"message": res.Fulfillment.Speech,
			"count":   count,
		}
	}
	return map[string]interface{}{"message": res.Fulfillment.Speech}
}

// Broadcast message to all connected devices
func Broadcast(m []byte) {
	BroadcastWS(m)
}

// Get some value
func Get(id string, name string, location string) (int, error) {
	things, err := thing.ReadGetterName(name)
	if err != nil {
		log.Println(err)
		return 0, err
	}
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
	return sum, nil
}
