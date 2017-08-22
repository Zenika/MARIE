#include <MQTTClient.h>
#include <BridgeClient.h> 
#include <Process.h>
#include <ArduinoJson.h>
#include <Adafruit_Sensor.h>
#include <DHT.h>

#define DHTTYPE DHT22
#define DHTPIN 5

#define PHOTOPIN A1

BridgeClient net;
MQTTClient client;
String macAddr;
StaticJsonBuffer<200> jsonBuffer;
DHT dht(DHTPIN, DHTTYPE);
int duration;
void setup () {
  Bridge.begin();
  client.begin("10.0.10.3", 1883, net);
  getMACAddress();

  pinMode(PHOTOPIN, INPUT);
  dht.begin();
  connect();
}

void getMACAddress () {
  Process p;
  p.runShellCommand("/sbin/ifconfig | grep HWaddr | grep wlan0 | awk '{print $5}'");
  while (p.running());

  macAddr = "";

  while (p.available() > 0) {
    char c = p.read();
    if (c > 10) {
      macAddr.concat(c);
    }
  }
  duration = micros();
  heartbeat();
}

void connect () {
  while (!client.connect("marie_env")) {
  }

  client.subscribe("getter/temperature");
  client.subscribe("location/couloir/getter/temperature");
  client.subscribe("macaddress/" + macAddr + "/getter/temperature");
  
  client.subscribe("getter/humidity");
  client.subscribe("location/couloir/getter/humidity");
  client.subscribe("macaddress/" + macAddr + "/getter/humidity");
 
  //client.subscribe("get_luminosity");
  client.publish("register", String("{\"macaddress\":\"" + macAddr + "\","
                                     "\"type\": \"environment\""
                                     "}"));
}

// the loop function runs over and over again forever
void loop() {
  if (micros() - duration > 15000) {
    heartbeat();
  }
  
  client.loop();

  if (!client.connected()) {
    connect();
  }
}

void messageReceived(String topic, String payload, char * bytes, unsigned int length) {
  JsonObject& root = jsonBuffer.parseObject(payload);
  if (topic == "getter/temperature" || topic == "macaddress/" + macAddr + "/getter/temperature" || topic == "location/couloir/getter/temperature") {
    float t = dht.readTemperature();
    if (isnan(t)) {
      client.publish("value/temperature", "{\"error\":\"NaN\"}");
    } else {
      client.publish("value/temperature", "{\"value\": " + String(t) + "}");
    }

  } else if (topic == "getter/humidity" || topic == "macaddress/" + macAddr + "/getter/humidity" || topic == "location/couloir/getter/humidity") {
    float h = dht.readHumidity();
    if (isnan(h)) {
      client.publish("value/humidity", "{\"error\":\"NaN\"}");
    } else {
      client.publish("value/humidity", "{\"value\": " + String(h) + "}");
    }
  } else if (topic == "get_luminosity") {
    int l = analogRead(PHOTOPIN);
    if (isnan(l)) {
      client.publish("luminosity_value", "{\"error\":\"NaN\"}");
    } else {
      client.publish("luminosity_value", "{\"value\": " + String(l) + "}");
    }
  }
  jsonBuffer.clear();
}

void heartbeat() {
  client.publish("heartbeat", "{\"macaddress\":\"" + macAddr + "\"}");
}

