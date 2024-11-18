package main

import (
	"context"
	"fmt"

	binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
	AvgPrice()
}

func AvgPrice() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "", baseURL)

	// AvgPrice
	avgPrice, err := client.NewAvgPriceService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(avgPrice))
}
