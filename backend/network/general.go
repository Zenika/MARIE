package network

import "gopkg.in/mgo.v2/bson"
import "github.com/Zenika/MARIE/backend/thing"

// Do something on a thing
func Do(id bson.ObjectId, name string, params map[string]interface{}) {
	t, err := thing.Read(id)
	if err != nil {
		return
	}
	switch t.Protocol {
	case "MQTT":
		DoMQTT(id, name, params)
		break
	}
}
