package main

import (
	"context"
	"fmt"

	binance_connector "github.com/zok2/binance-connector-go"
)

func main() {
	RecentTradesList()
}

func RecentTradesList() {
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient("", "", baseURL)

	// RecentTradesList
	recentTradesList, err := client.NewRecentTradesListService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(recentTradesList))
}
