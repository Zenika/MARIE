package network

import (
	"encoding/json"
	"log"
	"regexp"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/Zenika/MARIE/backend/config"
	"github.com/Zenika/MARIE/backend/record"

	"github.com/Zenika/MARIE/backend/thing"

	"github.com/gomqtt/client"
	"github.com/gomqtt/packet"
	"gopkg.in/mgo.v2"
)

type mqttConnection struct {
	mqtt *client.Client
	get  chan string
}

var mqttConn mqttConnection

// InitMQTT client
func InitMQTT() {

	cfg := config.Load()

	// Create mqtt client and connect
	mqtt := client.New()
	mqttCfg := client.NewConfigWithClientID(cfg.MQTTUrl, cfg.MQTTId)
	connectFuture, err := mqtt.Connect(mqttCfg)
	if err != nil {
		log.Fatal(err)
	}
	err = connectFuture.Wait(1 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	mqtt.Callback = handle

	// Subscribe to all getters
	things := thing.ReadAll()

	for _, t := range things {
		if t.Protocol == "MQTT" {
			for _, v := range t.Getters {
				// Subscribe to stored values
				_, err = mqtt.Subscribe(v.Name, 0)
				if err != nil {
					log.Fatal(err)
				}
				// Subscribe to real time values
				_, err = mqtt.Subscribe(v.Name+"_value", 0)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	mqtt.Subscribe("register", 0)
	mqttConn = mqttConnection{
		get:  make(chan string),
		mqtt: mqtt,
	}
	log.Println("MQTT client started")
}

// AddSubscription add subscribtion on a specific topic
func (c mqttConnection) AddSubscription(topic string) {
	c.mqtt.Subscribe(topic, 0)
}

// DoMQTT something on the thing
func (c mqttConnection) Do(mac string, name string, params map[string]interface{}) {
	req := map[string]string{"macaddress": mac}
	reqStr, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	c.mqtt.Publish(name, []byte(reqStr), 0, false)
}

// GetMQTT from a thing
func (c mqttConnection) Get(id string, name string, macaddress string) {
	go getRoutine(c.get, id)

	req := map[string]string{"macaddress": macaddress}
	reqStr, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	c.mqtt.Publish("get_"+name, []byte(reqStr), 0, false)
}

// Routine that  broadcast the results of getters
func getRoutine(c chan string, id string) {
	message := <-c
	var res map[string]interface{}
	err := json.Unmarshal([]byte(message), &res)
	if err != nil {
		log.Fatal(err)
	}
	res["id"] = id
	byt, err := json.Marshal(res)
	if err != nil {
		log.Fatal(byt)
	}
	hub.broadcast <- byt
}

// Handle the request on mqtt
func handle(msg *packet.Message, err error) {
	if msg.Topic == "register" {
		var t = thing.Thing{}
		err = json.Unmarshal(msg.Payload, &t)
		if err != nil {
			log.Fatal(err)
		}

		_, err := thing.ReadMacAddress(t.MacAddress)
		if err == mgo.ErrNotFound {
			t.ID = bson.NewObjectId()
			t.Protocol = "MQTT"
			thing.Create(t)
			res, err := json.Marshal(t)
			if err != nil {
				log.Fatal(err)
			}
			Broadcast(res)
		} else if err == nil {
			log.Fatal(err)
		}
		return
	}
	match, _ := regexp.MatchString("_value$", msg.Topic)
	if match {
		mqttConn.get <- string(msg.Payload)
		return
	}
	var r = record.Record{}
	err = json.Unmarshal(msg.Payload, &r)
	if err != nil {
		log.Fatal(err)
	}
	r.Name = msg.Topic
	record.Save(r)
}
