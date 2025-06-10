package main

import (
	"context"
	"fmt"

	binance_connector "github.com/zok2/binance-connector-go"
)

func main() {
	UserDelegationHistory()
}

func UserDelegationHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// UserDelegationHistoryService - /sapi/v1/asset/custody/transfer-history
	userDelegationHistory, err := client.NewUserDelegationHistoryService().
		Email("email@email.com").StartTime(1664442061000).EndTime(1664442078000).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(userDelegationHistory))
}
