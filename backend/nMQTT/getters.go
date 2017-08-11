package nMQTT

import (
	"encoding/json"
	"log"

	"github.com/Zenika/MARIE/backend/nWS"
)

// GetAll get values from every things
func (c MqttConnection) GetAll(id string, getter string) {
	go getRoutine(c.get, id)
	c.mqtt.Publish("getter/"+getter, []byte(""), 0, false)
}

// GetLocation get values from every things
func (c MqttConnection) GetLocation(id string, location string, getter string) {
	go getRoutine(c.get, id)
	c.mqtt.Publish("location/"+location+"/getter/"+getter, []byte(""), 0, false)
}

// GetMacAddress get values from every things
func (c MqttConnection) GetMacAddress(id string, macaddress string, getter string) {
	go getRoutine(c.get, id)
	c.mqtt.Publish("macaddress/"+macaddress+"/getter/"+getter, []byte(""), 0, false)
}

// Get from a thing
// func (c MqttConnection) Get(id string, name string, macaddress string) {
// 	req := map[string]string{"macaddress": macaddress}
// 	reqStr, err := json.Marshal(req)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	go getRoutine(c.get, id)

// 	c.mqtt.Publish("get_"+name, []byte(reqStr), 0, false)
// }

func getRoutine(c chan string, id string) {
	message := <-c

	var res map[string]interface{}
	err := json.Unmarshal([]byte(message), &res)
	if err != nil {
		log.Println(err)
		return
	}
	res["id"] = id
	byt, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}
	nWS.Broadcast(byt)
}
