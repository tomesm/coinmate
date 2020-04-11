package coinmate

type APITransactionHistory struct {
	Error        bool                          `json:"error"`
	ErrorMessage interface{}                   `json:"errorMessage"`
	Data         []APITransactionHistoryResult `json:"data"`
}

type APITransactionHistoryResult struct {
	Timestamp       int64    `json:"timestamp"`
	TransactionID   int64    `json:"transactionId"`
	TransactionType string   `json:"transactionType"`
	Price           *float64 `json:"price"`
	PriceCurrency   *string  `json:"priceCurrency"`
	Amount          float64  `json:"amount"`
	AmountCurrency  string   `json:"amountCurrency"`
	Fee             float64  `json:"fee"`
	FeeCurrency     string   `json:"feeCurrency"`
	Description     *string  `json:"description"`
	Status          string   `json:"status`
	OrderID         *int64   `json:"orderId"`
}

type APITransactionHistoryBody struct {
	Offset    string
	Limit     string
	Sort      string
	ClientId  string
	PublicKey string
	Nonce     string
	Signature string
}

func (api *APIClient) GetTransactionHistory(offset, limit, sort string) (th APITransactionHistory, err error) {
	body := api.getTransactionHistoryBody(offset, limit, sort)
	if err := api.Execute("POST", Endpoints{}.transactionHistory(), body, &th); err != nil {
		return th, err
	}
	return th, nil
}

func (api *APIClient) getTransactionHistoryBody(offset, limit, sort string) APITransactionHistoryBody {
	body := APITransactionHistoryBody{}
	body.Offset = offset
	body.Limit = limit
	body.Sort = sort
	body.ClientId = api.ClientId
	body.PublicKey = api.Key
	body.Nonce = api.nonce()
	body.Signature = api.signature(body.Nonce)
	return body
}
