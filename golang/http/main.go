package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	deviceToken := os.Getenv("DEVICE_TOKEN")
	if len(deviceToken) == 0 {
		panic("Token inválido")
	}

	deviceID := "dev-01"
	variableID := "01"

	value := "25.5"

	url := fmt.Sprintf("https://things.proiot.network/stream/device/%s/variable/%s/%v", deviceID, variableID, value)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("authorization", deviceToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New("statusCode: " + http.StatusText(res.StatusCode))
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
