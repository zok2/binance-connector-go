package main

import (
	"context"
	"fmt"

	binance_connector "github.com/zok2/binance-connector-go"
)

func main() {
	MarginAccountQueryAllOCO()
}

func MarginAccountQueryAllOCO() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountQueryAllOCOService - /sapi/v1/margin/allOrderList
	marginAccountQueryAllOCO, err := client.NewMarginAccountQueryAllOCOService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginAccountQueryAllOCO))
}
