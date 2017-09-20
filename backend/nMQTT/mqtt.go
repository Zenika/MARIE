package nMQTT

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/Zenika/MARIE/backend/config"
	"github.com/Zenika/MARIE/backend/nWS"

	"github.com/Zenika/MARIE/backend/thing"

	"github.com/gomqtt/client"
	"github.com/gomqtt/packet"
)

// MqttConnection represents a connection with two chans for actions and getters
type MqttConnection struct {
	mqtt *client.Client
	get  chan []byte
	do   chan []byte
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
				// Subscribe to real time
				_, err = mqtt.Subscribe("value/"+v.Name, 0)
				if err != nil {
					log.Fatal(err)
				}
				// Subscribe to recorded values
				_, err = mqtt.Subscribe("record/"+v.Name, 0)
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
	mqtt.Subscribe("return", 0)
	mqttConn = MqttConnection{
		get:  make(chan []byte),
		do:   make(chan []byte),
		mqtt: mqtt,
	}
	go mqttRoutine(mqttConn.do)
	go mqttRoutine(mqttConn.get)
	log.Println("MQTT client started")
}

// GetConnection returns the mqtt connection
func GetConnection() MqttConnection {
	return mqttConn
}

// AddGetSubscription add subscribtion on a topic for getters
func (c MqttConnection) AddGetSubscription(topic string) {
	log.Println(topic)
	c.mqtt.Subscribe("record/"+topic, 0)
	c.mqtt.Subscribe("value/"+topic, 0)
}

// Type returns the type of the connection
func (c MqttConnection) Type() string {
	return "MQTT"
}

// Handle the request on mqtt
func handle(msg *packet.Message, err error) {
	if msg == nil {
		return
	}
	if msg.Topic == "register" {
		cfg := config.Load()
		register(msg.Payload)
		heartbeatMsg := make(map[string]int)
		heartbeatMsg["time"] = cfg.HeartbeatTimer
		heartbeatMsgByt, err := json.Marshal(heartbeatMsg)
		if err != nil {
			log.Println(err)
			return
		}
		mqttConn.mqtt.Publish("heartbeat_time", heartbeatMsgByt, 0, false)
		return
	}

	if msg.Topic == "heartbeat" {
		heartbeat(msg.Payload)
		return
	}

	// See if topic begins with value
	match, _ := regexp.MatchString("^value", msg.Topic)
	if match {
		getterValueHandler(msg.Payload)
		return
	}

	match, _ = regexp.MatchString("^return", msg.Topic)
	if match {
		returnCodeHandler(msg.Payload)
		return
	}

	match, _ = regexp.MatchString("^record", msg.Topic)
	if match {
		getter := strings.Split(msg.Topic, "/")
		recordHandler(getter[1], msg.Payload)
		return
	}

	log.Println("Topic not recognized")
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
		t.State = true
		err = t.SetState(true)
		if err != nil {
			log.Println(err)
		}
		msg := make(map[string]interface{})
		msg["topic"] = "state-on"
		msg["macaddress"] = t.MacAddress
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

	// Add subscriptions to getters
	for _, t := range t.Getters {
		mqttConn.AddGetSubscription(t.Name)
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

	// Broadcast the thing creation
	nWS.BroadcastJSON(msg)
}

func mqttRoutine(c chan []byte) {
	for {
		message := <-c
		nWS.Broadcast(message)
	}
}
