package main

import (
	"context"
	"fmt"

	binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
	MarginIsolatedMarginTier()
}

func MarginIsolatedMarginTier() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginIsolatedMarginTierService - /sapi/v1/margin/isolatedMarginTier
	marginIsolatedMarginTier, err := client.NewMarginIsolatedMarginTierService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginIsolatedMarginTier))
}
