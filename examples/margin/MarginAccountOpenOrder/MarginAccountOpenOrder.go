package main

import (
	"context"
	"fmt"

	binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
	MarginAccountOpenOrder()
}

func MarginAccountOpenOrder() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountOpenOrderService - /sapi/v1/margin/openOrders
	marginAccountOpenOrder, err := client.NewMarginAccountOpenOrderService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountOpenOrder))
}
