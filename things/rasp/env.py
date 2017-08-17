import paho.mqtt.client as mqtt
from time import sleep
from utils import respond, isAction, isGetter, register
import Adafruit_DHT

sensor = 22
pin = 4

def on_connect(mqttc, obj, flags, rc):
    print("Connected")

def on_message(mqttc, obj, msg):
    print(msg.topic + " " + str(msg.payload))
    humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
    if isGetter(msg.topic, "humidity"):
      if humidity is not None:
         respond(mqttc, "humidity", humidity)
    if isGetter(msg.topic, "temperature"):
      if temperature is not None:
        respond(mqttc, "temperature", temperature)


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

register(mqttc, "Environnement", "env", "couloir", actions, getters)

rc = 0
while rc == 0:
  rc = mqttc.loop()
