#include <MQTTClient.h>
#include <BridgeClient.h>
#include <Process.h>
#include <ArduinoJson.h>

BridgeClient net;
MQTTClient client;
boolean on;
String macAddr;
StaticJsonBuffer<200> jsonBuffer;


void setup () {
  Bridge.begin();
  client.begin("broker.shiftr.io", net);
  pinMode(13, OUTPUT);
  getMACAddress();
  connect();
  on = false;
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
  while (!client.connect("marie_light", "4eabe27f", "c5e68ac27238e781")) {
    Serial.print(".");
  }

  client.subscribe("/on");
  client.subscribe("/off");
  client.publish("/register", String("{\"macaddress\":\"" + macAddr + "\","
                                     "\"location\": \"couloir\","
                                     "\"type\": \"light\","
                                     "\"actions\":["
                                     "{\"name\":\"on\"},"
                                     "{\"name\":\"off\"}"
                                     "]}"));
}

// the loop function runs over and over again forever
void loop() {
  client.loop();

  if (!client.connected()) {
    connect();
  }

  if (on) {
    digitalWrite(13, HIGH);
  } else {
    digitalWrite(13, LOW);
  }

}

void messageReceived(String topic, String payload, char * bytes, unsigned int length) {
  JsonObject& root = jsonBuffer.parseObject(payload);
  String requiredMacAddress = root["macaddress"];
  if (requiredMacAddress == macAddr) {
    if (topic == "/on") {
      on = true;
    } else if (topic == "/off") {
      on = false;
    }
  }
  jsonBuffer.clear();
}
