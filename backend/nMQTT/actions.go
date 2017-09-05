package nMQTT

import (
	"encoding/json"
	"log"

	"github.com/Zenika/MARIE/backend/thing"
)

// DoAll do an action on all things
func (c MqttConnection) DoAll(id string, thingType string, action string, params []thing.Parameter) {
	go mqttRoutine(c.do, id)
	res, err := json.Marshal(params)
	if err != nil {
		log.Println(err)
		return
	}
	c.mqtt.Publish("type/"+thingType+"/action/"+action, res, 0, false)
}

// DoLocation do an action on all things with a specific location
func (c MqttConnection) DoLocation(id string, thingType string, location string, action string, params []thing.Parameter) {
	go mqttRoutine(c.do, id)
	res, err := json.Marshal(params)
	if err != nil {
		log.Println(err)
		return
	}
	c.mqtt.Publish("type/"+thingType+"/location/"+location+"/action/"+action, res, 0, false)
}

// DoMacAddress do an action on a particular thing
func (c MqttConnection) DoMacAddress(id string, macaddress string, action string, params []thing.Parameter) {
	go mqttRoutine(c.do, id)
	res, err := json.Marshal(params)
	if err != nil {
		log.Println(err)
		return
	}
	c.mqtt.Publish("macaddress/"+macaddress+"/action/"+action, res, 0, false)
}
