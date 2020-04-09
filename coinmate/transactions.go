package coinmate

//"strconv"
//"time"

/*const (*/
//TransactionHistory = "/transactionHistory"
/*)*/

type APITransactionHistory struct {
	Error        bool                          `json:"error"`
	ErrorMessage interface{}                   `json:"errorMessage"`
	Data         []APITransactionHistoryResult `json:"data"`
}

type APITransactionHistoryResult struct {
	Timestamp       int64   `json:"timestamp"`
	TransactionID   int64   `json:"transactionId"`
	TransactionType string  `json:"transactionType"`
	Price           float64 `json:"price"`
	PriceCurrency   string  `json:"priceCurrency"`
	Amount          float64 `json:"amount"`
	AmountCurrency  string  `json:"amountCurrency"`
	Fee             float64 `json:"fee"`
	FeeCurrency     string  `json:"feeCurrency"`
	Description     string  `json:"description"`
	Status          Status  `json:"status`
	OrderID         int64   `json:"orderId"`
}

/*type Currency string*/

//const (
//Btc Currency = "BTC"
//Czk Currency = "CZK"
//Eur Currency = "EUR"
/*)*/

type Status string

const (
	Ok      Status = "OK"
	Pending Status = "PENDING"
)

type APITransactionHistoryBody struct {
	Offset    string
	Limit     string
	Sort      string
	ClientId  string
	PublicKey string
	Nonce     string
	Signature string
}

func (api *APIClient) GetTransactionHistory() (th APITransactionHistory, err error) {
	body := APITransactionHistoryBody{}
	body.Offset = "0"
	body.Limit = "10"
	body.Sort = "ASC"
	body.ClientId = api.ClientId
	body.PublicKey = api.Key
	body.Nonce = api.nonce()
	body.Signature = api.signature(body.Nonce)

	if err := api.Execute("POST", api.Endpoints.transactionHistory(), body, &th); err != nil {
		return th, err
	}
	return th, nil
}
