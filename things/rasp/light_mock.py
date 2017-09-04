import paho.mqtt.client as mqtt
import sys
from utils import return_code, respond, isAction, getAction, isGetter, register, heartbeat

on = 0

def on_connect(mqttc, obj, flags, rc):
    print("Connected")

def on_message(mqttc, obj, msg):
    print(msg.topic + " " + str(msg.payload))
    global on
    if isAction(msg.topic, "on"):
      on = 1
    elif isAction(msg.topic, "off"):
      on = 0
    return_code(mqttc, getAction(msg.topic), 0)
    if isGetter(msg.topic, "state"):
      respond(mqttc, "state", on)


mqttc = mqtt.Client()
mqttc.on_message = on_message
mqttc.on_connect = on_connect
mqttc.connect("10.0.10.3", 1883, 60)

actions = [
  { "name": "on"},
  { "name": "off"}
]
getters = [
  { 
    "name": "state",
    "type": "boolean"
  }
]

register(mqttc, "Lumiere_mock", "light", "couloir", actions, getters)

timer = heartbeat([mqttc])

rc = 0
try:
  while rc == 0:
    rc = mqttc.loop()
except KeyboardInterrupt:
  timer.cancel()
  sys.exit()