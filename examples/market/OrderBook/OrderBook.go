package main

import (
	"context"
	"fmt"

	binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
	OrderBook()
}

func OrderBook() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "", baseURL)

	// OrderBook
	orderBook, err := client.NewOrderBookService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(orderBook))
}
