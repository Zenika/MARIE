import json
import threading

heartbeat_time = 1000

def setHeartbeat(msg):
  global heartbeat_time
  data = json.loads(msg)
  heartbeat = data["time"]

def heartbeat(mqttc):
  global heartbeat_time
  t = threading.Timer(heartbeat_time / 1000, heartbeat, [mqttc])
  t.daemon = True
  t.start()
  message = {"macaddress": getmac("wlan0")}
  mqttc[0].publish("heartbeat", json.dumps(message), qos=2)
  return t

def isAction(topic, action):
  if "action" not in topic:
    return 0
  strs = topic.split("/")
  name = strs[len(strs) - 1]
  return name == action

def getAction(topic):
  strs = topic.split("/")
  name = strs[len(strs) - 1]
  return name

def isGetter(topic, getter):
  if "getter" not in topic:
    return 0
  strs = topic.split("/")
  name = strs[len(strs) - 1]
  return name == getter

def doSubscribeActions(mqttc, thingType, location, macaddress, action):
  mqttc.subscribe("type/" + thingType + "/action/" + action)
  mqttc.subscribe("type/" + thingType + "/location/" + location + "/action/" + action)
  mqttc.subscribe("macaddress/" + macaddress + "/action/" + action)
  print("Subscribed to " + action + " action")

def doSubscribeGetters(mqttc, location, macaddress, getter):
  mqttc.subscribe("getter/" + getter)
  mqttc.subscribe("macaddress/" + macaddress + "/getter/" + getter)
  mqttc.subscribe("location/" + location + "/getter/" + getter)
  print("Subscribed to " + getter + " getter")

def record(mqttc, getter, value_fn):
  t = threading.Timer(1, record, [mqttc, getter, value_fn])
  t.daemon = True
  t.start()
  message = {"value": value_fn(), "macaddress": getmac("wlan0")}
  mqttc.publish("record/" + getter, json.dumps(message), qos=2)

def return_code(mqttc, msg, code):
  message = {"code": code, "id": getId(msg)}
  mqttc.publish("return", json.dumps(message), qos=2)

def respond(mqttc, msg, getter, value):
  message = {"value": value, "id": getId(msg)}
  mqttc.publish("value/" + getter, json.dumps(message), qos=2)

def getId(msg):
  data = json.loads(msg)
  return data["id"]

def getParam(msg, paramName):
  data = json.loads(msg)
  paramString = data["params"]
  params = json.loads(paramString)
  param = filter(lambda x: x["name"] == paramName, params)[0]
  return param["value"]

def register(mqttc, name, thingType, location, actions, getters):
  mqttc.subscribe("heartbeat_time")
  macaddress = getmac("wlan0")
  for action in actions:
    doSubscribeActions(mqttc, thingType, location, macaddress, action["name"])
  for getter in getters:
    doSubscribeGetters(mqttc, location, macaddress, getter["name"])

  message = ("{"
              "\"name\":\"" + name + "\","
              "\"type\":\"" + thingType + "\","
              "\"macaddress\":\"" + getmac("wlan0") + "\","
              "\"location\":\"" + location + "\","
              "\"actions\":" + json.dumps(actions) + ","
              "\"getters\":" + json.dumps(getters) + ""
           "}")
  mqttc.publish("register", message, qos=2)
  print("Registered")

def getmac(interface):
  try:
    mac = open('/sys/class/net/'+interface+'/address').readline()
  except:
    mac = "00:00:00:00:00:00"

  return mac[0:17]
