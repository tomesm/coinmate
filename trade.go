package coinmate

type APITradingPairs struct {
	Error        bool                    `json:"error"`
	ErrorMessage interface{}             `json:"errorMessage"`
	Data         []APITradingPairsResult `json:"data"`
}

type APITradingPairsResult struct {
	Name                              string  `json:"name"`
	FirstCurrency                     string  `json:"firstCurrency"`
	SecondCurrency                    string  `json:"secondCurrency"`
	PriceDecimals                     int64   `json:"priceDecimals"`
	LotDecimals                       int64   `json:"lotDecimals"`
	MinAmount                         float64 `json:"minAmount"`
	TradesWebSocketChannelID          string  `json:"tradesWebSocketChannelId"`
	OrderBookWebSocketChannelID       string  `json:"orderBookWebSocketChannelId"`
	TradeStatisticsWebSocketChannelID string  `json:"tradeStatisticsWebSocketChannelId"`
}

func (api *APIClient) GetTradingPairs() (tp APITradingPairs, err error) {
	if err := api.Execute("GET", Endpoints{}.tradingPairs(), nil, &tp); err != nil {
		return tp, err
	}
	return tp, nil
}
