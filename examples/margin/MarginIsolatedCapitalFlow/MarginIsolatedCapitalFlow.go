package main

import (
	"context"
	"fmt"

	binance_connector "github.com/zok2/binance-connector-go"
)

func main() {
	MarginIsolatedCapitalFlow()
}

func MarginIsolatedCapitalFlow() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// MarginIsolatedCapitalFlowService - /sapi/v1/margin/capital-flow
	marginIsolatedCapitalFlow, err := client.NewMarginIsolatedCapitalFlowService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(marginIsolatedCapitalFlow))
}
