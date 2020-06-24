const mqtt = require('mqtt')

const deviceToken = process.env.DEVICE_TOKEN
if (!deviceToken) {
  console.error("Token inválido")
  return
}

const options = {
    username: deviceToken
}

deviceID = "5ef20c1f47f0e40019a54876"

const client = mqtt.connect('mqtt://mqtt.proiot.network', options)

client.on('connect', function () {
    console.log('Conectado')

    const payload = {
      "01": "21"
    }

    client.publish(`device/${deviceID}`, JSON.stringify(payload))
    
    client.end()
})
