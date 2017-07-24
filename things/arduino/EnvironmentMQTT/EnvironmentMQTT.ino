#include <MQTTClient.h>
#include <BridgeClient.h>
#include <Process.h>
#include <ArduinoJson.h>
#include <Adafruit_Sensor.h>
#include <DHT.h>

#define DHTTYPE DHT22
#define DHTPIN 5

BridgeClient net;
MQTTClient client;
String macAddr;
StaticJsonBuffer<200> jsonBuffer;
DHT dht(DHTPIN, DHTTYPE);

void setup () {
  Bridge.begin();
  client.begin("broker.shiftr.io", net);
  getMACAddress();
    
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
}

void connect () {
  while (!client.connect("marie_env", "4eabe27f", "c5e68ac27238e781")) {
  }

  client.subscribe("/get_temperature");
  client.subscribe("/get_humidity");
  client.publish("/register", String("{\"macaddress\":\"" + macAddr + "\","
                                     "\"type\": \"environment\","
                                     "\"getters\":["
                                     "{\"name\":\"temperature\","
                                     "\"type\":\"number\"}"
                                     "]}"));
}

// the loop function runs over and over again forever
void loop() {
  client.loop();

  if (!client.connected()) {
    connect();
  }
}

void messageReceived(String topic, String payload, char * bytes, unsigned int length) {
  JsonObject& root = jsonBuffer.parseObject(payload);
  String requiredMacAddress = root["macaddress"];
  if (requiredMacAddress == macAddr) {
    if (topic == "/get_temperature") {
      float t = dht.readTemperature();
      if (isnan(t)) {
        client.publish("/temperature_value", "{\"error\":\"NaN\"}");
      } else {
        client.publish("/temperature_value", "{\"value\": " + String(t) + "}");
      }

    } else {
      float h = dht.readHumidity();
      if (isnan(h)) {
        client.publish("/humidity_value", "{\"error\":\"NaN\"}");
      } else {
        client.publish("/humidity_value", "{\"value\": " + String(h) + "}");
      }
    }


  }
  jsonBuffer.clear();
}
