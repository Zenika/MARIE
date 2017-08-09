import paho.mqtt.client as mqtt
from gpiozero import LED
from time import sleep

def getmac(interface):
  try:
    mac = open('/sys/class/net/'+interface+'/address').readline()
  except:
    mac = "00:00:00:00:00:00"

  return mac[0:17]

def on_connect(mqttc, obj, flags, rc):
    print("rc: " + str(rc))

def on_message(mqttc, obj, msg):
    global on
    if msg.topic == "on":
      on = 1
    else:
      on = 0
    print(msg.topic + " " + str(msg.qos) + " " + str(msg.payload))


def on_publish(mqttc, obj, mid):
    print("mid: " + str(mid))
    pass

def on_subscribe(mqttc, obj, mid, granted_qos):
    print("Subscribed: " + str(mid) + " " + str(granted_qos))

mqttc = mqtt.Client()
mqttc.on_message = on_message
mqttc.on_connect = on_connect
mqttc.on_publish = on_publish
mqttc.on_subscribe = on_subscribe
mqttc.connect("10.0.10.3", 1883, 60)

print("tuple")
message = ("{"
            "\"name\":\"Template\","
              "\"type\":\"template\","
              "\"macaddress\":\"" + getmac("wlan0") + "\","
              "\"location\":\"template\","
              "\"actions\":["
              "{\"name\":\"template\"}"
              "],"
              "\"getters\":["
              "{\"name\":\"template\", \"type\":\"number\"}"
              "]"
           "}")
(rc, mid) = mqttc.publish("register", message, qos=2)
mqttc.subscribe("template", 0)

rc = 0
while rc == 0:
  rc = mqttc.loop()
