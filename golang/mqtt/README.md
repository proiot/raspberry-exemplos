# Golang - MQTT

## Build

```bash
# Exportando o device eui
export DEVICE_EUI=xxx

# Exportando o device hash app
export DEVICE_HASH_APP=xxx

# Exportando o token de acesso
export DEVICE_TOKEN=xxx

# 1º Rodando de forma simples
go run main.go

# 2º Gerando o binário
go build -o bin/proiot-golang-mqtt

# 3º Rodando a aplicação
./bin/proiot-golang-mqtt
```
