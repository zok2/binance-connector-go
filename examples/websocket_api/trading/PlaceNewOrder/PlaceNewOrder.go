package main

import (
	"context"
	"fmt"
	"log"

	binance_connector "github.com/zok2/binance-connector-go"
)

func main() {
	PlaceNewOrderExample()
}

func PlaceNewOrderExample() {
	client := binance_connector.NewWebsocketAPIClient("api_key", "secret_key", "wss://ws-api.testnet.binance.vision/ws-api/v3")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	response, err := client.NewPlaceNewOrderService().Symbol("BTCUSDT").Side("BUY").OrderType("MARKET").Quantity(0.01).
		Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance_connector.PrettyPrint(response))

	client.WaitForCloseSignal()
}
