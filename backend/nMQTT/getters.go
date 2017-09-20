package nMQTT

import (
	"encoding/json"
	"log"

	"github.com/Zenika/MARIE/backend/record"
	"github.com/Zenika/MARIE/backend/thing"
)

// GetAll get values from every things
func (c MqttConnection) GetAll(id string, getter string) {
	c.mqtt.Publish("getter/"+getter, []byte("{\"id\":\""+id+"\"}"), 0, false)
}

// GetLocation get values from every things
func (c MqttConnection) GetLocation(id string, location string, getter string) {
	c.mqtt.Publish("location/"+location+"/getter/"+getter, []byte("{\"id\":\""+id+"\"}"), 0, false)
}

// GetMacAddress get values from every things
func (c MqttConnection) GetMacAddress(id string, macaddress string, getter string) {
	c.mqtt.Publish("macaddress/"+macaddress+"/getter/"+getter, []byte("{\"id\":\""+id+"\"}"), 0, false)
}

func getterValueHandler(payload []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(payload, &data)
	if err != nil {
		log.Println(err)
		return
	}
	data["topic"] = "getter-value"
	msgString, err := json.Marshal(data)
	mqttConn.get <- msgString
}

func recordHandler(getter string, payload []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(payload, &data)
	if err != nil {
		log.Println(err)
		return
	}
	t, err := thing.ReadMacAddress(data["macaddress"].(string))
	if err != nil {
		log.Println(err)
		return
	}
	r := record.Record{
		Name:    getter,
		ThingID: t.ID,
		Value:   data["value"],
	}

	r.Save()
}
