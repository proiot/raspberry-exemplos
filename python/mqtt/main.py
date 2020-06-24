import os
import sys
import paho.mqtt.client as mqtt

deviceToken = os.getenv("DEVICE_TOKEN")

if deviceToken is None:
    print("Token invalido")
    exit()

deviceID = "5ef20c1f47f0e40019a54876"

def on_connect(client, userdata, flags, rc):
    print("Conectado: "+str(rc))

    client.publish("device/" + deviceID, "{'01':'21'}")

try:
    client = mqtt.Client()
    client.username_pw_set(username=deviceToken, password=None)
    client.on_connect = on_connect

    client.connect("mqtt.proiot.network", 1883, 60)
    client.loop_forever()

except KeyboardInterrupt:
    sys.exit(0)