package network

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/Zenika/MARIE/backend/mqtt"
	"github.com/Zenika/MARIE/backend/thing"
)

// Do something on all things that match action and room
func Do(thingType string, action string, params map[string]interface{}, room string) {
	things := thing.ReadActionName(action)
	for _, t := range things {
		if t.Type == thingType {
			if room == "" || t.Location == room {
				DoID(t.ID, action, params)
			}
		}
	}
}

// DoID something on a thing with its id
func DoID(id bson.ObjectId, action string, params map[string]interface{}) {
	t, err := thing.Read(id)
	if err != nil {
		log.Fatal(err)
	}
	switch t.Protocol {
	case "MQTT":
		mqtt.Do(t.MacAddress, action, params)
		break
	}
}
