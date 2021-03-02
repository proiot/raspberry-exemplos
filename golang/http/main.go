package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
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

func main() {
	deviceEUI := os.Getenv("DEVICE_EUI")
	if len(deviceEUI) == 0 {
		panic("Device EUI inválido")
	}

	deviceToken := os.Getenv("DEVICE_TOKEN")
	if len(deviceToken) == 0 {
		panic("Token inválido")
	}

	url := fmt.Sprintf("https://things.conn.proiot.network/%s", deviceEUI)

	payload := getPayload()

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		panic(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", deviceToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusOK {
		err = errors.New("statusCode: " + http.StatusText(res.StatusCode))
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result:")
	fmt.Println(string(body))
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
