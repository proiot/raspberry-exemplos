const http = require("https");

const deviceToken = process.env.DEVICE_TOKEN
if (!deviceToken) {
  console.error("Token inválido")
  return
}

const deviceID = "dev-01"
const variableID = "01"

const value = "25.5"

const path = `/stream/device/${deviceID}/variable/${variableID}/${value}`

const options = {
  "method": "POST",
  "hostname": "things.proiot.network",
  "path": path,
  "headers": {
    "authorization": deviceToken
  }
};

const req = http.request(options, (res) => {
  const chunks = [];

  if (res.statusCode !== 200) {
    console.error("statusCode:", res.statusMessage)
    return
  } 

  res.on("data",  (chunk) => {
    chunks.push(chunk);
  });

  res.on("end", () => {
    const body = Buffer.concat(chunks);
    console.log(body.toString());
  });
});

req.end();
