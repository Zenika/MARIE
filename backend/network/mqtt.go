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
	err = connectFuture.Wait(5 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	mqtt.Callback = handle

	// Subscribe to all getters
	things, err := thing.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

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
	c.mqtt.Subscribe(topic+"_value", 0)
}

// DoMQTT something on the thing
func (c mqttConnection) Do(mac string, name string, params map[string]interface{}) {
	req := map[string]string{"macaddress": mac}
	reqStr, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return
	}
	c.mqtt.Publish(name, []byte(reqStr), 0, false)
}

// GetMQTT from a thing
func (c mqttConnection) Get(id string, name string, macaddress string) {

	req := map[string]string{"macaddress": macaddress}
	reqStr, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		return
	}
	go getRoutine(c.get, id)

	c.mqtt.Publish("get_"+name, []byte(reqStr), 0, false)
}

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
	hub.broadcast <- byt
}

// Handle the request on mqtt
func handle(msg *packet.Message, err error) {
	if msg.Topic == "register" {
		var t = thing.Thing{}
		err = json.Unmarshal(msg.Payload, &t)
		if err != nil {
			log.Println(err)
			return
		}

		_, err := thing.ReadMacAddress(t.MacAddress)
		if err == mgo.ErrNotFound {
			// Add missing attributes
			t.ID = bson.NewObjectId()
			t.Protocol = "MQTT"

			// Insert thing in database
			thing.Create(t)
			res, err := json.Marshal(t)
			if err != nil {
				log.Println(err)
				return
			}

			// Broadcast the thing creation
			Broadcast(res)
		} else if err != nil {
			log.Println(err)
		}
		return
	}

	// See if topic ends with _value
	match, _ := regexp.MatchString("_value$", msg.Topic)

	if match {
		mqttConn.get <- string(msg.Payload)
		return
	}

	// In other cases, record the value
	var r = record.Record{}
	err = json.Unmarshal(msg.Payload, &r)
	if err != nil {
		log.Println(err)
		return
	}
	r.Name = msg.Topic
	record.Save(r)
}
