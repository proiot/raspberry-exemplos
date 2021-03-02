package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Payload payload
type Payload struct {
	Data []Variable `json:"data"`
}

// Variable variable
type Variable struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

var (
	deviceEUI     string
	deviceHashAPP string
)

func main() {
	deviceHashAPP = os.Getenv("DEVICE_HASH_APP")
	if len(deviceHashAPP) == 0 {
		panic("Device Hash App inválido")
	}

	deviceEUI = os.Getenv("DEVICE_EUI")
	if len(deviceEUI) == 0 {
		panic("Device EUI inválido")
	}

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

func connectionLostHandler(client mqtt.Client, err error) {
	fmt.Println("Conexão perdida")
}

func onConnectHandler(client mqtt.Client) {
	fmt.Println("Conectado")
	fmt.Println("Enviando mensagem...")

	payload := getPayload()

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	token := client.Publish(fmt.Sprintf("proiot/%s/%s/tx", deviceHashAPP, deviceEUI), 0, false, string(payloadBytes))
	token.Wait()

	fmt.Println("Mensagem enviada")
}

func getPayload() Payload {
	temp := Variable{
		Name:  "temp",
		Value: randomVariable(25, 45),
	}

	battery := Variable{
		Name:  "bat",
		Value: randomVariable(0, 100),
	}

	data := []Variable{}
	data = append(data, temp)
	data = append(data, battery)

	payload := Payload{
		Data: data,
	}

	return payload
}

func randomVariable(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
