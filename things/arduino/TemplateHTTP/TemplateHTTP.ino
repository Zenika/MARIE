#include <BridgeClient.h>
#include <HttpClient.h>
#include <Process.h>
#include <ArduinoJson.h>


BridgeClient net;
String macAddr;
StaticJsonBuffer<200> jsonBuffer;
IPAddress server(10,0,10,45);

void setup() {
  Bridge.begin();
  getMACAddress();
  HttpClient client;

  String message = "{\"name\":\"Template\","
                     "\"macaddress\":\"" + macAddr + "\","
                     "\"location\":\"couloir\","
                     "\"type\":\"light\","
                     "\"actions\":["
                     "{\"name\":\"on\"},"
                     "{\"name\":\"off\"}"
                     "]"
                     "}";
  int len = message.length() + 1;
  char msg[len];
  message.toCharArray(msg, len);
  client.post("http://10.0.10.145:8081/api/things/register", msg);
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

void loop() {
  // put your main code here, to run repeatedly:

}
