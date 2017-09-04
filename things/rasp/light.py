import paho.mqtt.client as mqtt
from gpiozero import LED
from utils import return_code, respond, isAction, getAction, isGetter, register, heartbeat

led = LED(17)

on = 0

def on_connect(mqttc, obj, flags, rc):
    print("Connected")

def on_message(mqttc, obj, msg):
    print(msg.topic + " " + str(msg.payload))
    global on
    if isAction(msg.topic, "on"):
      on = 1
    else:
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

register(mqttc, "Lumiere", "light", "couloir", actions, getters)

heartbeat([mqttc])

rc = 0
while rc == 0:
  if on == 1:
    led.on()
  if on == 0:
    led.off()
  rc = mqttc.loop()
