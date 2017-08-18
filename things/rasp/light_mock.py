import paho.mqtt.client as mqtt
from utils import return_code, isAction, getAction, isGetter, register

on = 0

def on_connect(mqttc, obj, flags, rc):
    print("Connected")

def on_message(mqttc, obj, msg):
    print(msg.topic + " " + str(msg.payload))
    global on
    on = 0
    if isAction(msg.topic, "on"):
      on = 1
    return_code(mqttc, getAction(msg.topic), 0)


mqttc = mqtt.Client()
mqttc.on_message = on_message
mqttc.on_connect = on_connect
mqttc.connect("10.0.10.3", 1883, 60)

actions = [
  { "name": "on"},
  { "name": "off"}
]
getters = []

register(mqttc, "Lumiere_mock", "light", "couloir", actions, getters)


rc = 0
while rc == 0:
  rc = mqttc.loop()
