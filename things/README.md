# MARIE things

MARIE can manage any sort of things, and you can find scripts for them in this folder.

## Getting Started

### Arduino

You can develop new things that run on arduino. Due to limitations, it can be impossible to properly register thing attributes with MQTT. You'll have to only register the MAC address and then modify it on the web interface.

### Raspberry Pi

The scripts for Rapsberry Pi are developed in Python. You dispose of a script names utils which provides lot of useful functions.

### MQTT

Things can communicate with the MQTT Protocol. They have to register to the backend to be used. The JSON to send to this is :

```json
{
    "name": "ThingName",
    "type": "ThingType",
    "location": "ThingLocation",
    "macaddress": "00:11:22:33:44:55",
    "actions": [
        {
            "name": "action1"
        }
    ],
    "getters": [
        {
            "name": "getter1",
            "type": "number"
        }
    ]
}
```

If the thing want to do something, it has to subscribe to these topics:

```text
/type/:thingType/action/:action
/type/:thingType/location/:location/action/:action
/macaddress/:macaddress/action/:action
```

If you want to send a return code to notify the application, you have to send this message :

```json
// Topic
/return

// Content
{
    "code": 0
}
```

And if you want to fetch data from them, you have to subscribe to:

```text
/getter/:getter
/macaddress/:macaddress/getter/:getter
/location/:location/getter/:getter
```

The thing will be requested on those topics if the user want some data. To answer, you have to send the data like this:

```json
// Topic
/value/:getter

//Content
{
    "value": ":value"
}
```

### Env

To use the DHT sensor with the Raspberry Pi, you have to do :

```shell
git clone https://github.com/adafruit/Adafruit_Python_DHT.git
cd Adafruit_Python_DHT
sudo apt-get update
sudo apt-get install build-essential python-dev python-openssl
sudo python setup.py install
```

### Speech

The speech thing is a Node.JS server. You have to configure it with the Google Project Id and the JSON keyfile to enable speech recognition. You'll also have to change the model you want to use if you want another hotword than Snowboy.