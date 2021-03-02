import os
import http.client
import json
import random

deviceEUI = os.getenv("DEVICE_EUI")
deviceToken = os.getenv("DEVICE_TOKEN")

if deviceEUI is None:
    print("Device EUI invalido")
    exit()


if deviceToken is None:
    print("Token invalido")
    exit()

def get_payload():
    data = {
        'data': [
            {
                'name': 'temp',
                'value': random.randint(20, 35)
            },
            {
                'name': 'bat',
                'value': random.randint(0, 100)
            }
        ]
    }

    return data

try:
    conn = http.client.HTTPSConnection("things.conn.proiot.network")

    headers = {
        'Content-type': 'application/json',
        'authorization': deviceToken
    }

    payload = get_payload()

    conn.request("POST", "/" + deviceEUI, json.dumps(payload), headers)

    response = conn.getresponse()
    body = response.read()

    print("Status", response.status)
    print(body.decode("utf-8"))

except:
    raise
