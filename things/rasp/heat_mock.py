import paho.mqtt.client as mqtt
import json
from time import sleep
from utils import respond, return_code, isAction, isGetter, register, getAction, heartbeat

def on_connect(mqttc, obj, flags, rc):
    print("Connected")

def on_message(mqttc, obj, msg):
    print(msg.topic + " " + str(msg.payload))
    if isAction(msg.topic, "heat"):
        data = json.loads(str(msg.payload))
        temperature = filter(lambda x: x["name"] == "temperature", data)[0]
        print(temperature["value"])
        return_code(mqttc, getAction(msg.topic), 0)

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
