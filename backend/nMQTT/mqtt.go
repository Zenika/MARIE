package nMQTT

import (
	"encoding/json"
	"log"
	"regexp"
	"time"

	"github.com/Zenika/MARIE/backend/config"
	"github.com/Zenika/MARIE/backend/nWS"
	"github.com/Zenika/MARIE/backend/record"

	"github.com/Zenika/MARIE/backend/thing"

	"github.com/gomqtt/client"
	"github.com/gomqtt/packet"
)

// MqttConnection represents a connection with two chans for actions and getters
type MqttConnection struct {
	mqtt *client.Client
	get  chan string
	do   chan string
}

var mqttConn MqttConnection

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
				_, err = mqtt.Subscribe("value/"+v.Name, 0)
				if err != nil {
					log.Fatal(err)
				}
				// Subscribe to real time values
				_, err = mqtt.Subscribe(v.Name+"_value", 0)
				if err != nil {
					log.Fatal(err)
				}
			}

			for _, v := range t.Actions {
				// Subscribe to return code
				_, err = mqtt.Subscribe("return/"+v.Name, 0)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	mqtt.Subscribe("register", 0)
	mqtt.Subscribe("heartbeat", 0)
	mqttConn = MqttConnection{
		get:  make(chan string),
		do:   make(chan string),
		mqtt: mqtt,
	}
	log.Println("MQTT client started")
}

// GetConnection returns the mqtt connection
func GetConnection() MqttConnection {
	return mqttConn
}

// AddGetSubscription add subscribtion on a topic for getters
func (c MqttConnection) AddGetSubscription(topic string) {
	c.mqtt.Subscribe(topic, 0)
	c.mqtt.Subscribe(topic+"_value", 0)
	c.mqtt.Subscribe("value/"+topic, 0)
}

// AddDoSubscription add subscription on a topic for actions
func (c MqttConnection) AddDoSubscription(topic string) {
	c.mqtt.Subscribe("return/"+topic, 0)
}

// Type returns the type of the connection
func (c MqttConnection) Type() string {
	return "MQTT"
}

// Handle the request on mqtt
func handle(msg *packet.Message, err error) {
	if msg.Topic == "register" {
		register(msg.Payload)
		return
	}

	if msg.Topic == "heartbeat" {
		heartbeat(msg.Payload)
		return
	}

	// See if topic begins with value
	match, _ := regexp.MatchString("^value", msg.Topic)

	if match {
		mqttConn.get <- string(msg.Payload)
		return
	}

	match, _ = regexp.MatchString("^return", msg.Topic)

	if match {
		mqttConn.do <- string(msg.Payload)
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
	r.Save()
}

func heartbeat(payload []byte) {
	var t = thing.Thing{}
	err := json.Unmarshal(payload, &t)
	if err != nil {
		log.Println(err)
		return
	}
	t, err = thing.ReadMacAddress(t.MacAddress)
	if err != nil {
		log.Println(err)
		return
	}
	if t.State == false {
		err = t.SetState(true)
		if err != nil {
			log.Println(err)
		}
		msg := make(map[string]interface{})
		msg["state-on"] = t.MacAddress
		nWS.BroadcastJSON(msg)
	}
	t.UpdateHeartBeat()
}

func register(payload []byte) {
	// Read the thing in the payload
	var t = thing.Thing{}
	err := json.Unmarshal(payload, &t)
	if err != nil {
		log.Println(err)
		return
	}
	// Add Protocol
	if t.Protocol == "" {
		t.Protocol = "MQTT"
	}

	//Register thing
	t, err = thing.Register(t)
	if err != nil {
		log.Println(err)
		return
	}

	// Transform to JSON
	res, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
		return
	}

	var msg map[string]interface{}
	json.Unmarshal(res, &msg)
	msg["topic"] = "register"

	res, err = json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}

	// Broadcast the thing creation
	nWS.Broadcast(res)
}

func mqttRoutine(c chan string, id string) {
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
