package coinmate

type apiOrderBook struct {
	Error        bool               `json:"error"`
	ErrorMessage interface{}        `json:"errorMessage"`
	Data         apiOrderBookResult `json:"data"`
}

type apiOrderBookResult struct {
	Asks []apiOderBookAsk `json:"asks"`
	Bids []apiOderBookAsk `json:"bids"`
}

type apiOderBookAsk struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

type APIOderBookBody struct {
	CurrencyPair      string
	GroupByPriceLimit string
}

func (api *APIClient) GetOrderBook(body APIOderBookBody) (ob apiOrderBook, err error) {
	if err := api.Execute("GET", Endpoints{}.orderBook(), body, &ob); err != nil {
		return ob, err
	}
	return ob, nil
}
