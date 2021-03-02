import os
import sys
import json
import time
import ssl
import random
import paho.mqtt.client as mqtt

deviceHashApp = os.getenv("DEVICE_HASH_APP")
deviceEUI = os.getenv("DEVICE_EUI")
deviceToken = os.getenv("DEVICE_TOKEN")

if deviceHashApp is None:
    print("Device Hash App invalido")
    exit()

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

def on_connect(client, userdata, flags, rc):
    print("Conectado: "+str(rc))

try:
    client = mqtt.Client()

    client.tls_set(None,
               None,
               None, 
               cert_reqs=ssl.CERT_NONE, 
               tls_version=ssl.PROTOCOL_TLSv1, 
               ciphers=None)

    client.username_pw_set(username=deviceToken, password=None)
    client.on_connect = on_connect

    client.connect("mqtt.proiot.network", 8883, 60)

    while True:
        payload = get_payload()

        print("Enviando...")
        client.publish("proiot/" + deviceHashApp + "/" + deviceEUI + "/tx", json.dumps(payload))

        time.sleep(5)

except KeyboardInterrupt:
    sys.exit(0)