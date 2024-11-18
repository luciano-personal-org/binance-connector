package main

import (
	"context"
	"fmt"
	"log"

	binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
	CurrentOpenOCOExample()
}

func CurrentOpenOCOExample() {
	client := binance_connector.NewWebsocketAPIClient("api_key", "secret_key", "wss://ws-api.testnet.binance.vision/ws-api/v3")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	response, err := client.NewCurrentOpenOCOService().Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance_connector.PrettyPrint(response))

	client.WaitForCloseSignal()
}
