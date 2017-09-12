package nMQTT

import (
	"encoding/json"
	"log"

	"github.com/Zenika/MARIE/backend/thing"
)

// DoAll do an action on all things
func (c MqttConnection) DoAll(id string, thingType string, action string, params []thing.Parameter) {
	bytParams, err := parseParams(id, params)
	if err != nil {
		log.Println(err)
		return
	}
	c.mqtt.Publish("type/"+thingType+"/action/"+action, bytParams, 0, false)
}

// DoLocation do an action on all things with a specific location
func (c MqttConnection) DoLocation(id string, thingType string, location string, action string, params []thing.Parameter) {
	bytParams, err := parseParams(id, params)
	if err != nil {
		log.Println(err)
		return
	}
	c.mqtt.Publish("type/"+thingType+"/location/"+location+"/action/"+action, bytParams, 0, false)
}

// DoMacAddress do an action on a particular thing
func (c MqttConnection) DoMacAddress(id string, macaddress string, action string, params []thing.Parameter) {
	bytParams, err := parseParams(id, params)
	if err != nil {
		log.Println(err)
		return
	}
	c.mqtt.Publish("macaddress/"+macaddress+"/action/"+action, bytParams, 0, false)
}

func parseParams(id string, params []thing.Parameter) ([]byte, error) {
	res, err := json.Marshal(params)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	parsed := make(map[string]interface{})
	parsed["id"] = id
	parsed["params"] = string(res)

	bytParsed, err := json.Marshal(parsed)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return bytParsed, nil
}
