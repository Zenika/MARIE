import paho.mqtt.client as mqtt
from utils import record, setHeartbeat, respond, isAction, isGetter, register, heartbeat
import Adafruit_DHT

sensor = 22
pin = 4

def on_connect(mqttc, obj, flags, rc):
    print("Connected")

def on_message(mqttc, obj, msg):
    payload = str(msg.payload)
    print(msg.topic + " " + payload)
    if isGetter(msg.topic, "humidity"):
      humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
      if humidity is not None:
        respond(mqttc, payload, "humidity", humidity)
    if isGetter(msg.topic, "temperature"):
      humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
      if temperature is not None:
        respond(mqttc, payload, "temperature", temperature)
    elif msg.topic == "heartbeat_time":
      setHeartbeat(msg.payload)
      heartbeat([mqttc])
      record(mqttc, "humidity", humid)
      record(mqttc, "temperature", temp)

mqttc = mqtt.Client()
mqttc.on_connect = on_connect
mqttc.on_message = on_message
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

def humid():
  humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
  return humidity
def temp():
  humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
  return temperature

rc = 0
while rc == 0:
  rc = mqttc.loop()
