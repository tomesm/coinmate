package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tomesm/go-coinmate/coinmate"
)

func main() {
	c := coinmate.NewAPIClient(os.Getenv("COINMATE_KEY"),
		os.Getenv("COINMATE_SECRET"),
		os.Getenv("COINMATE_CLIENT_ID"))

	// tradingPairs, err := c.GetTradingPairs()
	// if err != nil {
	// 	log.Fatal("Trading pairs not fetched")
	// }
	// fmt.Println(tradingPairs.Data[0].FirstCurrency)

	transactionHistory, err := c.GetTransactionHistory()
	if err != nil {
		log.Fatal("Transaction history not fetched")
	}

	for _, th := range transactionHistory.Data {
		fmt.Println(th.Status)
	}

	//fmt.Println(transactionHistory.Data[0].Description)
}
