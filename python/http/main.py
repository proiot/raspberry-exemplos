import os 
import http.client

deviceToken = os.getenv("DEVICE_TOKEN")

if deviceToken is None:
    print("Token invalido")
    exit()

deviceID = "dev-01"
variableID = "01"
value = "25.5"

conn = http.client.HTTPSConnection("things.proiot.network")
conn.request("POST", "/stream/device/" + deviceID + "/variable/" + variableID + "/" + value, None, { 'authorization': deviceToken })

res = conn.getresponse()
data = res.read()

print("Status", res.status)
print(data.decode("utf-8"))