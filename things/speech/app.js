const snowboy = require('snowboy');
const models = new snowboy.Models();
const record = require('node-record-lpcm16')
const Speech = require('@google-cloud/speech')
const W3CWebSocket = require('websocket').w3cwebsocket;

const client = new W3CWebSocket('ws://localhost:8081/ws');

const mqtt = require('mqtt')
const mqttc = mqtt.connect('mqtt://10.0.10.3:1883')

let mac = '';



mqttc.on('connect', () => {
  console.log('MQTT connected')
  require('getmac').getMac((err, macaddr) => {
    if (err) {
      throw err
    }
    mac = macaddr
    mqttc.publish('register', JSON.stringify({
      "name": "Speech",
      "type": "speech",
      "macaddress": mac,
      "location": "couloir"
    }), {}, () => {
      heartbeat()
    })
  })
})

mqttc.on('message', (topic, message) => {
  console.log(topic)
})

function heartbeat() {
  mqttc.publish('heartbeat', JSON.stringify({
    "macaddress": mac
  }))
  setTimeout(() => {heartbeat()}, 15000)
}

client.onerror = err => {
  console.log('Connection error ')
  throw JSON.stringify(err);
}

client.onopen = function () {
  console.log('Client connected')
}

// The encoding of the audio file, e.g. 'LINEAR16'
const encoding = 'LINEAR16';

// The sample rate of the audio file in hertz, e.g. 16000
const sampleRateHertz = 16000;

// The BCP-47 language code to use, e.g. 'en-US'
const languageCode = 'fr-FR';

models.add({
  file: 'resources/snowboy.umdl',
  sensitivity: '0.6',
  hotwords: 'snowboy'
})


const detector = new snowboy.Detector({
  resource: 'resources/common.res',
  audioGain: 2.0,
  models
})

detector.on('error', function () {
  console.log('error');
});

detector.on('hotword', (index, hotword) => {
  mqttc.publish('start_speech')
  const speechClient = Speech({
    projectId: require('./config').projectId,
    keyFilename: './keyfile.json'
  })

  const request = {
    config: {
      encoding: encoding,
      sampleRateHertz: sampleRateHertz,
      languageCode: languageCode
    },
    interimResults: false // If you want interim results, set this to true
  };

  // Create a recognize stream
  const recognizeStream = speechClient.streamingRecognize(request)
    .on('error', console.error)
    .on('data', (data) => {
      if (data.results[0] && data.results[0].alternatives[0]) {
        console.log(data.results[0].alternatives[0].transcript)
        const message = {type: 'speech', message: data.results[0].alternatives[0].transcript}
        mqttc.publish('speech', JSON.stringify({'message': message}))
        client.send(JSON.stringify({type: 'speech', message: message}).toString())
      } else {
        console.log('Reached transcription time limit, press Ctrl+C')
      }
        recognizeStream.end();
        console.log('Recognize ended');
    });

  console.log('Recognize started');
  mic.pipe(recognizeStream)
});

const mic = record.start( {
  threshold: 0,
  // verbose: true,
  // silence: '10.0'
})
mic.pipe(detector)