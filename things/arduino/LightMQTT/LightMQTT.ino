#include <MQTTClient.h>
#include <BridgeClient.h>
#include <Bridge.h>

BridgeClient net;
MQTTClient client;
boolean on;

void setup () {
  Bridge.begin();
  client.begin("broker.shiftr.io", net);
  pinMode(13, OUTPUT);

  connect();
  on = false;
}

void connect () {
  while (!client.connect("marie_light", "4eabe27f", "c5e68ac27238e781")) {
    Serial.print(".");
  }

  client.subscribe("/on");
  client.subscribe("/off");
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
  if (topic == "/on") {
    on = true;
  } else if (topic == "/off") {
    on = false;
  }
}
