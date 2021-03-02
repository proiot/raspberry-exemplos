const http = require("https");

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
  "method": "POST",
  "hostname": "things.conn.proiot.network",
  "path": `/${deviceEUI}`,
  "headers": {
    "authorization": deviceToken,
    "content-type": "application/json"
  }
};

const req = http.request(options, (res) => {
  const chunks = [];

  if (res.statusCode !== 200) {
    console.error("statusCode:", res.statusMessage)
    return
  }

  res.on("data", (chunk) => {
    chunks.push(chunk);
  });

  res.on("end", () => {
    const body = Buffer.concat(chunks);
    console.log(body.toString());
  });
});

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

req.write(JSON.stringify(payload));

req.end();
