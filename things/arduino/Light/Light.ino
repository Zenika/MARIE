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
#define MQTTNAME "marie_light"

#define LIGHTPIN 13

BridgeClient net;
MQTTClient client;  
String macAddr;
StaticJsonBuffer<128> jsonBuffer;
HttpClient c;

boolean on;

void setup () {
  Bridge.begin();
  getMACAddress();
  registerThing();
  client.begin(MQTTSERVER, 1883, net);
  connect();
  
  client.subscribe("on");
  client.subscribe("off");
  pinMode(LIGHTPIN, OUTPUT);
  on = false;
}

// the loop function runs over and over again forever
void loop() {
  client.loop();

  if (!client.connected()) {
    connect();
  }

  if (on) {
    digitalWrite(LIGHTPIN, HIGH);
  } else {
    digitalWrite(LIGHTPIN, LOW);
  }
}

void registerThing() {
  char message[256];
  message[0] = '\0';
  strcat(message, "{\"macaddress\":\"");
  strcat(message, macAddr.c_str());
  strcat(message, "\",");
  strcat(message, "\"name\":\"Lumière\","
                  "\"location\":\"couloir\","
                  "\"protocol\":\"MQTT\","
                  "\"type\":\"light\""
                  "}");
  sendRequest(REGISTERSUFFIX, message);

  message[0] = '\0';
  strcat(message, "{\"macaddress\":\"");
  strcat(message, macAddr.c_str());
  strcat(message, "\",");
  strcat(message, "\"actions\":["
                  "{\"name\":\"on\"},"
                  "{\"name\":\"off\"}"
                  "]}");

  sendRequest(ACTIONSSUFFIX, message);
}

void messageReceived(String topic, String payload, char * bytes, unsigned int length) {
  JsonObject& root = jsonBuffer.parseObject(payload);
  String requiredMacAddress = root["macaddress"];
  if (requiredMacAddress == macAddr) {
    if (topic == "on") {
      on = true;
    } else if (topic == "off") {
      on = false;
    }
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
