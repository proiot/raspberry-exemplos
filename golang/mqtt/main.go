package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	deviceID = "5ef20c1f47f0e40019a54876"
)

func connectionLostHandler(client mqtt.Client, err error) {
	fmt.Println("Conexão perdida")
}

func onConnectHandler(client mqtt.Client) {
	fmt.Println("Conectado")
	fmt.Println("Enviando mensagem...")

	token := client.Publish(fmt.Sprintf("device/%s", deviceID), 0, false, "ok")
	token.Wait()

	fmt.Println("Mensagem enviada")
}

func main() {
	deviceToken := os.Getenv("DEVICE_TOKEN")
	if len(deviceToken) == 0 {
		panic("Token inválido")
	}

	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s", "mqtt.proiot.network:1883"))
	opts.AutoReconnect = true
	opts.Username = deviceToken
	opts.SetConnectionLostHandler(connectionLostHandler)
	opts.SetOnConnectHandler(onConnectHandler)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	time.Sleep(5 * time.Second)
}
