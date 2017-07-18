package network

import (
	"github.com/Zenika/MARIE/backend/mqtt"
	"github.com/Zenika/MARIE/backend/thing"
	"gopkg.in/mgo.v2/bson"
)

// Do something on a thing
func Do(id bson.ObjectId, name string, params map[string]interface{}) {
	t, err := thing.Read(id)
	if err != nil {
		return
	}
	switch t.Protocol {
	case "MQTT":
		mqtt.Do(id, name, params)
		break
	}
}
