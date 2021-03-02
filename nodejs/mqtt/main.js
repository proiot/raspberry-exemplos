const mqtt = require('mqtt')

const deviceHashApp = process.env.DEVICE_HASH_APP
if (!deviceHashApp) {
  console.error("Device HashApp inválido")
  return
}

const deviceEUI = process.env.DEVICE_EUI
if (!deviceEUI) {
  console.error("DeviceEUI inválido")
  return
}

const deviceToken = process.env.DEVICE_TOKEN
if (!deviceToken) {
  console.error("Token inválido")
  return
}

const options = {
  username: deviceToken
}

const client = mqtt.connect('mqtts://mqtt.proiot.network:8883', options)

client.on('connect', function () {
  console.log('Conectado')

  const payload = {
    data: [
      {
        "name": "temp",
        "value": (Math.random() * (45 - 20) + 10).toFixed(2)
      },
      {
        "name": "bat",
        "value": Math.floor(Math.random() * 100)
      }
    ]
  }

  client.publish(`proiot/${deviceHashApp}/${deviceEUI}/tx`, JSON.stringify(payload))

  client.end()
})
