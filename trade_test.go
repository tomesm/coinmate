package coinmate

import (
	"reflect"
	"testing"
)

func TestGetTradingPairs(t *testing.T) {
	c := NewAPIClient("CpmRVUJL0OGByT2otAfCKeeDdU6yfi6OzvnXcAwaHvE", "", "6")
	c.Endpoint = "https://private-anon-1e146dbd9c-coinmate.apiary-mock.com/api"

	res, err := c.GetTradingPairs()
	if err != nil {
		t.Error("Expected nil, got ", err.Error())
	}
	tests := []struct {
		name   string
		result apiTradingPairsResult
		want   apiTradingPairsResult
	}{
		{
			name:   "BTC_EUR Trading pair",
			result: res.Data[0],
			want:   expectedData("EUR"),
		},
		{
			name:   "BTC_CZK Trading pair",
			result: res.Data[1],
			want:   expectedData("CZK"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if !reflect.DeepEqual(tc.result, tc.want) {
				t.Errorf("Expected %s trading pair, got: %s\n", tc.want.Name, tc.result.Name)
			}
		})
	}
}

func expectedData(curr string) apiTradingPairsResult {
	return apiTradingPairsResult{
		Name:                              "BTC_" + curr,
		FirstCurrency:                     "BTC",
		SecondCurrency:                    curr,
		PriceDecimals:                     2,
		LotDecimals:                       8,
		MinAmount:                         0.001,
		TradesWebSocketChannelID:          "trades-BTC_" + curr,
		OrderBookWebSocketChannelID:       "order_book-BTC_" + curr,
		TradeStatisticsWebSocketChannelID: "statistics-BTC_" + curr,
	}
}
