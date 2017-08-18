package nMQTT

// GetAll get values from every things
func (c MqttConnection) GetAll(id string, getter string) {
	go mqttRoutine(c.get, id)
	c.mqtt.Publish("getter/"+getter, []byte(""), 0, false)
}

// GetLocation get values from every things
func (c MqttConnection) GetLocation(id string, location string, getter string) {
	go mqttRoutine(c.get, id)
	c.mqtt.Publish("location/"+location+"/getter/"+getter, []byte(""), 0, false)
}

// GetMacAddress get values from every things
func (c MqttConnection) GetMacAddress(id string, macaddress string, getter string) {
	go mqttRoutine(c.get, id)
	c.mqtt.Publish("macaddress/"+macaddress+"/getter/"+getter, []byte(""), 0, false)
}
