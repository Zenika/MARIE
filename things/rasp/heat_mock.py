import paho.mqtt.client as mqtt
import json
from time import sleep
from utils import getParam, respond, return_code, isAction, isGetter, register, getAction, heartbeat

def on_connect(mqttc, obj, flags, rc):
    print("Connected")

def on_message(mqttc, obj, msg):
    payload = str(msg.payload)
    print(msg.topic + " " + payload)
    if isAction(msg.topic, "heat"):
        print(getParam(msg.payload, "temperature"))
        return_code(mqttc, payload, 0)

mqttc = mqtt.Client()
mqttc.on_message = on_message
mqttc.on_connect = on_connect
mqttc.connect("10.0.10.3", 1883, 60)

actions = [
    {
        "name": "heat",
        "parameters": [
            {
                "name": "temperature",
                "type": "number"
            }
        ]
    }
]
getters = [ ]

register(mqttc, "Chaudiere", "boiler", "couloir", actions, getters)

heartbeat([mqttc])

rc = 0
while rc == 0:
  rc = mqttc.loop()
