package main

import (
	"fmt"
	"os"
	"log"
	"github.com/tomesm/go-coinmate/coinmate"
)

func main() {
	c := coinmate.NewAPIClient(os.Getenv("COINMATE_KEY"), 
						   os.Getenv("COINMATE_SECRET"),
						   os.Getenv("COINMATE_CLIENT_ID"))

	tradingPairs, err := c.GetTradingPairs()
	if err != nil {
		log.Fatal("Trading pairs not fetched")
	}
	fmt.Println(tradingPairs.Data[0].FirstCurrency)

	
}