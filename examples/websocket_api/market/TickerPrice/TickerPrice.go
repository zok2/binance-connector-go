package main

import (
	"context"
	"fmt"
	"log"

	binance_connector "github.com/zok2/binance-connector-go"
)

func main() {
	TickerPrice()
}

func TickerPrice() {
	client := binance_connector.NewWebsocketAPIClient("", "", "wss://ws-api.testnet.binance.vision/ws-api/v3")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	response, err := client.NewTickerPriceService().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance_connector.PrettyPrint(response))

	client.WaitForCloseSignal()
}
