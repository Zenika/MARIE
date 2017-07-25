# MARIE things

MARIE can manage any sort of things.

## Arduino

You can develop new things that run on arduino

### MQTT

You can find a template inside the arduino direct to use to communicate with MARIE using MQTT.
The thing will send a message on the register topic to let MARIE know that it exists.

Then MARIE can ask for data in an async way and broadcast this data to websockets.