package main

import (
	"context"
	"fmt"

	binance_connector "github.com/zok2/binance-connector-go"
)

func main() {
	AccountInfo()
}

func AccountInfo() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// AccountInfoService - /sapi/v1/account/apiTradingStatus
	accountInfo, err := client.NewAccountInfoService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(accountInfo))
}
