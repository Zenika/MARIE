import paho.mqtt.client as mqtt
from time import sleep
from utils import respond, isAction, isGetter, register

def on_connect(mqttc, obj, flags, rc):
    print("Connected")

def on_message(mqttc, obj, msg):
    global on
    if isGetter( msg.topic, "humidity"):
      respond(mqttc, "humidity", 12)
    print(msg.topic + " " + str(msg.payload))


mqttc = mqtt.Client()
mqttc.on_message = on_message
mqttc.on_connect = on_connect
mqttc.connect("10.0.10.3", 1883, 60)

actions = [
]
getters = [
  { 
    "name": "humidity",
    "type": "number"
  }
]

register(mqttc, "Environnement", "env", "couloir", actions, getters)

rc = 0
while rc == 0:
  rc = mqttc.loop()
