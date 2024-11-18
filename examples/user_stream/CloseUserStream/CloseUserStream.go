package main

import (
	"context"
	"fmt"

	binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
	CloseUserStream()
}

func CloseUserStream() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	close := client.NewCloseUserStream().ListenKey("your_listen_key").
		Do(context.Background())
	fmt.Println(close)
}
