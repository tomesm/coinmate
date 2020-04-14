package coinmate

import (
	"fmt"
	"testing"
)

func TestGetTransactionHistory(t *testing.T) {
	c := NewAPIClient("CpmRVUJL0OGByT2otAfCKeeDdU6yfi6OzvnXcAwaHvEgit ", "", "6")
	c.Endpoint = "https://private-anon-4cb33c8b65-coinmate.apiary-mock.com/api"

	body := getTransactionHistoryBody("0", "10", "ASC")

	res, err := c.GetTransactionHistory(body)
	if err != nil {
		t.Error("Expected nil, got ", err.Error())
	}
	if len(res.Data)-1 != 10 {
		t.Errorf("want length %d; got %d", 10, len(res.Data)-1)
	}

	tests := []struct {
		name   string
		result string
		want   string
	}{
		{
			name:   "CZK deposit",
			result: res.Data[0].TransactionType,
			want:   "DEPOSIT",
		},
		{
			name:   "CZK amount currency",
			result: res.Data[0].AmountCurrency,
			want:   "CZK",
		},
		{
			name:   "CZK amount",
			result: fmt.Sprintf("%f", res.Data[0].Amount),
			want:   "1000.000000",
		},
		{
			name:   "BTC withdrawal",
			result: res.Data[1].TransactionType,
			want:   "WITHDRAWAL",
		},
		{
			name:   "BTC status pending",
			result: res.Data[1].Status,
			want:   "PENDING",
		},
		{
			name:   "BTC fee",
			result: fmt.Sprintf("%f", res.Data[1].Fee),
			want:   "0.000300",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.result != tc.want {
				t.Errorf("Expected %s trading pair, got: %s\n", tc.want, tc.result)
			}
		})
	}

}

func getTransactionHistoryBody(offset, limit, sort string) APITransactionHistoryBody {
	body := APITransactionHistoryBody{}
	body.Offset = offset
	body.Limit = limit
	body.Sort = sort
	body.ClientId = "6"
	body.PublicKey = "CpmRVUJL0OGByT2otAfCKeeDdU6yfi6OzvnXcAwaHvE"
	body.Nonce = "0"
	body.Signature = "E4F27EAB0A836873CEE325A2526CC7476E2A3F2BE8026CADFB7A66B72D582DE8"
	return body
}
