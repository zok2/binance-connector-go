package main

import (
	"context"
	"fmt"

	binance_connector "github.com/zok2/binance-connector-go"
)

func main() {
	GetFiatPaymentHistory()
}

func GetFiatPaymentHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// GetFiatPaymentHistoryService - /sapi/v1/fiat/payments
	getFiatPaymentHistory, err := client.NewGetFiatPaymentHistoryService().
		TransactionType("0").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(getFiatPaymentHistory))
}
