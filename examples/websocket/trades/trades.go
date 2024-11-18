package main

import (
	"fmt"
	"time"

	binance_connector "github.com/luciano-personal-org/binance-connector"
)

func main() {
	WsTradeExample()
}

func WsTradeExample() {
	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false)
	wsTradeHandler := func(event *binance_connector.WsTradeEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, stopCh, err := websocketStreamClient.WsTradeServe("LTCBTC", wsTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopCh to exit
	go func() {
		time.Sleep(10 * time.Second)
		stopCh <- struct{}{}
	}()
	<-doneCh
}
