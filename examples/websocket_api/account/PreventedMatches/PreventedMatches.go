package main

import (
	"context"
	"fmt"
	"log"

	binance_connector "github.com/zok2/binance-connector-go"
)

func main() {
	PreventedMatchesExample()
}

func PreventedMatchesExample() {
	client := binance_connector.NewWebsocketAPIClient("api_key", "secret_key")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	response, err := client.NewAccountPreventedMatchesService().Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance_connector.PrettyPrint(response))

	client.WaitForCloseSignal()
}
