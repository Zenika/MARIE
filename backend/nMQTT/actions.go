package nMQTT

// DoAll do an action on all things
func (c MqttConnection) DoAll(id string, thingType string, action string, params map[string]interface{}) {
	go mqttRoutine(c.do, id)
	c.mqtt.Publish("type/"+thingType+"/action/"+action, []byte(""), 0, false)
}

// DoLocation do an action on all things with a specific location
func (c MqttConnection) DoLocation(id string, thingType string, location string, action string, params map[string]interface{}) {
	go mqttRoutine(c.do, id)
	c.mqtt.Publish("type/"+thingType+"/location/"+location+"/action/"+action, []byte(""), 0, false)
}

// DoMacAddress do an action on a particular thing
func (c MqttConnection) DoMacAddress(id string, macaddress string, action string, params map[string]interface{}) {
	go mqttRoutine(c.do, id)
	c.mqtt.Publish("macaddress/"+macaddress+"/action/"+action, []byte(""), 0, false)
}
