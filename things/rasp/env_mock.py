import paho.mqtt.client as mqtt
import random
from utils import setHeartbeat, respond, isAction, isGetter, register, heartbeat

def on_connect(mqttc, obj, flags, rc):
    print("Connected")

def on_message(mqttc, obj, msg):
    payload = str(msg.payload)
    print(msg.topic + " " + payload)
    if isGetter(msg.topic, "humidity"):
      humidity = random.randint(0, 100)
      respond(mqttc, payload, "humidity", humidity)
    if isGetter(msg.topic, "temperature"):
      temperature = random.randint(0, 30)
      respond(mqttc, payload, "temperature", temperature)
    elif msg.topic == "heartbeat_time":
      setHeartbeat(msg.payload)
      heartbeat([mqttc])

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
  },
  {
    "name": "temperature",
    "type": "number"
  }
]

register(mqttc, "Environnement_mock", "env", "couloir", actions, getters)

heartbeat([mqttc])

rc = 0
while rc == 0:
  rc = mqttc.loop()
