# MARIE things

MARIE can manage any sort of things.

## Arduino

You can develop new things that run on arduino

### MQTT

You can find a template inside the arduino direct to use to communicate with MARIE using MQTT.
The thing will send a message on the register topic to let MARIE know that it exists.

Then MARIE can ask for data in an async way and broadcast this data to websockets.

### Env

To use the DHT sensor with the raspberry pi, you have to do : 

```shell
git clone https://github.com/adafruit/Adafruit_Python_DHT.git
cd Adafruit_Python_DHT
sudo apt-get update
sudo apt-get install build-essential python-dev python-openssl
sudo python setup.py install
```