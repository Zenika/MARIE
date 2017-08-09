#include <MQTTClient.h>
#include <HttpClient.h>
#include <BridgeClient.h>
#include <Process.h>
#include <ArduinoJson.h>

#define MQTTSERVER "10.0.10.3"
#define HTTPSERVER "10.0.10.145:8081"
#define REGISTERSUFFIX "/api/things/register"
#define ACTIONSSUFFIX "/api/things/actions"
#define GETTERSSUFFIX "/api/things/getters"
#define MQTTNAME "marie_template"

BridgeClient net;
MQTTClient client;  
String macAddr;
StaticJsonBuffer<128> jsonBuffer;
HttpClient c;

void setup () {
  Bridge.begin();
  getMACAddress();
  registerThing();
  client.begin(MQTTSERVER, 1883, net);
  connect();
}

// the loop function runs over and over again forever
void loop() {
  client.loop();

  if (!client.connected()) {
    connect();
  }
}

void registerThing() {
  char message[256];
  message[0] = '\0';
  strcat(message, "{\"macaddress\":\"");
  strcat(message, macAddr.c_str());
  strcat(message, "\",");
  strcat(message, "\"name\":\"Template\","
                  "\"location\":\"couloir\","
                  "\"protocol\":\"MQTT\","
                  "\"type\":\"template\""
                  "}");
  sendRequest(REGISTERSUFFIX, message);

  message[0] = '\0';
  strcat(message, "{\"macaddress\":\"");
  strcat(message, macAddr.c_str());
  strcat(message, "\",");
  strcat(message, "\"actions\":["
                  "{\"name\":\"template\"},"
                  "{\"name\":\"template2\"}"
                  "]}");

  sendRequest(ACTIONSSUFFIX, message);

  message[0] = '\0';
  strcat(message, "{\"macaddress\":\"");
  strcat(message, macAddr.c_str());
  strcat(message, "\",");
  strcat(message, "\"getters\":["
                  "{\"name\":\"template\",\"type\":\"number\"}"
                  "]}");

  sendRequest(GETTERSSUFFIX, message);
}

void messageReceived(String topic, String payload, char * bytes, unsigned int length) {
  JsonObject& root = jsonBuffer.parseObject(payload);
  String requiredMacAddress = root["macaddress"];
  if (requiredMacAddress == macAddr) {

  }
  jsonBuffer.clear();
}

void connect () {
  while (!client.connect(MQTTNAME)) {
  }
}

void sendRequest (const char suffix[], char *message) {
  char url[128];
  url[0] = '\0';
  strcat(url, "http://");
  strcat(url, HTTPSERVER);
  strcat(url, suffix);
  c.post(url, message);
  while (c.available()) {
    char res = c.read();
  }
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
