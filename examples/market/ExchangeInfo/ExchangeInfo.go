package main

import (
	"context"
	"fmt"

	binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
	ExchangeInfo()
}

func ExchangeInfo() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "", baseURL)

	// ExchangeInfo
	exchangeInfo, err := client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(exchangeInfo))
}
