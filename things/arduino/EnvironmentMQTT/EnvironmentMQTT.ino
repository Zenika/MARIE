#include <MQTTClient.h>
#include <BridgeClient.h>
#include <Process.h>
#include <ArduinoJson.h>

BridgeClient net;
MQTTClient client;
String macAddr;
StaticJsonBuffer<200> jsonBuffer;


void setup () {
  Bridge.begin();
  client.begin("broker.shiftr.io", net);
  getMACAddress();
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
    Serial.print(".");
  }

  client.subscribe("/get_temperature");
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
    client.publish("/temperature_value", String("{\"value\":\"15\"}"));
  }
  jsonBuffer.clear();
}
