package mqtt

import (
	"encoding/json"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/Zenika/MARIE/backend/config"
	"github.com/Zenika/MARIE/backend/record"

	"github.com/Zenika/MARIE/backend/thing"

	"github.com/gomqtt/client"
	"github.com/gomqtt/packet"
	"gopkg.in/mgo.v2"
)

var mqtt *client.Client

// Init MQTT client
func Init() {
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

	mqtt.Callback = handle

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
	mqtt.Subscribe("register", 0)
	log.Println("MQTT client started")
}

// AddSubscription add subscribtion on a specific topic
func AddSubscription(topic string) {
	mqtt.Subscribe(topic, 0)
}

// Do something on the thing
func Do(mac string, name string, params map[string]interface{}) {
	// paramStr, err := json.Marshal(params)
	// if err != nil {
	// 	log.Println("Error while parsing JSON")
	// 	return
	// }
	req := map[string]string{"macaddress": mac}
	reqStr, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	mqtt.Publish(name, []byte(reqStr), 0, false)
}

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
		} else if err == nil {
			log.Fatal(err)
		}
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
