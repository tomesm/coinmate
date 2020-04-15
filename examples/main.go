package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tomesm/coinmate"
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

	th := coinmate.APITransactionHistoryBody{}
	th.Offset = "0"
	th.Limit = "2"
	th.Sort = "ASC"
	th.ClientId = c.ClientId
	th.PublicKey = c.Key
	th.Nonce = c.Nonce()
	th.Signature = c.Signature(th.Nonce)

	transactionHistory, err := c.GetTransactionHistory(th)
	if err != nil {
		log.Fatal("Transaction history not fetched")
	}

	for _, th := range transactionHistory.Data {
		fmt.Println(th.Status)
	}

	orderBookBody := coinmate.APIOderBookBody{}
	orderBookBody.CurrencyPair = "BTC_EUR"
	orderBookBody.GroupByPriceLimit = "False"

	orderBook, err := c.GetOrderBook(orderBookBody)
	if err != nil {
		log.Fatal("Order book not fetched")
	}
	fmt.Println(orderBook.Data.Asks[0].Price)
}
