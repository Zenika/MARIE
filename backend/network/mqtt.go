package network

import (
	"encoding/json"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/Zenika/MARIE/backend/config"

	"github.com/Zenika/MARIE/backend/thing"

	"github.com/gomqtt/client"
	"github.com/gomqtt/packet"
)

var mqtt *client.Client

// InitMQTT client
func InitMQTT() {
	cfg := config.Load()

	// Create mqtt client and connect
	mqtt = client.New()
	mqttCfg := client.NewConfigWithClientID(cfg.MQTTUrl, cfg.MQTTId)
	connectFuture, err := mqtt.Connect(mqttCfg)
	if err != nil {
		log.Fatal(err)
	}
	err = connectFuture.Wait(1 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	mqtt.Callback = handleGetter

	// Subscribe to all getters
	things := thing.ReadAll()

	for _, t := range things {
		if t.Protocol == "MQTT" {
			for _, v := range t.Getters {
				_, err = mqtt.Subscribe(v.Name, 0)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	log.Println("MQTT client started")
}

// DoMQTT something on the thing
func DoMQTT(id bson.ObjectId, name string, params map[string]interface{}) {
	paramStr, err := json.Marshal(params)
	if err != nil {
		log.Println("Error while parsing JSON")
		return
	}
	mqtt.Publish(name, []byte(paramStr), 0, false)
}

func handleGetter(msg *packet.Message, err error) {
	var r = thing.Record{}
	err = json.Unmarshal(msg.Payload, &r)
	if err != nil {
		log.Println("Unmarshal error MQTT")
		return
	}
	r.Name = msg.Topic
	thing.SaveRecord(r)
}
