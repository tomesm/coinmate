package coinmate

/*const (*/
//TransactionHistory = "/transactionHistory"
/*)*/

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
	PriceCurrency   string   `json:"priceCurrency"`
	Amount          float64  `json:"amount"`
	AmountCurrency  string   `json:"amountCurrency"`
	Fee             float64  `json:"fee"`
	FeeCurrency     string   `json:"feeCurrency"`
	Description     *string  `json:"description"`
	Status          Status   `json:"status"`
	OrderID         *int64   `json:"orderId"`
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
	offset int    `json:"offset`
	limit  int    `json:"limit`
	sort   string `json:"sort"`
	// timestampFrom int64 `json:"timestampFrom`
	clientID  int64  `json:"clientId"`
	publicKey string `json:"publicKey`
	nonce     int64  `json:"nonce"`
	signature string `json:"signature"`
}

func (api *APIClient) GetTransactionHistory() (th APITransactionHistory, err error) {

	body := APITransactionHistoryBody{
		offset:    0,
		limit:     10,
		sort:      "ASC",
		clientID:  api.convertString(api.ClientId),
		publicKey: api.Key,
		nonce:     api.convertString(api.Nonce),
		signature: api.createSignature(),
	}

	if err := api.Execute("POST", api.Endpoints.transactionHistory(), body, &th); err != nil {
		return th, err
	}
	return th, nil

	// body := []byte(offset=0&limit=10&sort=ASC&timestampFrom=1401390154803&clientId=&publicKey=CpmRVUJL0OGByT2otAfCKeeDdU6yfi6OzvnXcAwaHvE&nonce=0&signature=E4F27EAB0A836873CEE325A2526CC7476E2A3F2BE8026CADFB7A66B72D582DE8)

	// req, _ := http.NewRequest("POST", "https://coinmate.io/api/transactionHistory", bytes.NewBuffer(body))

	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// resp, err := client.Do(req)

	// if err != nil {
	// 	fmt.Println("Errored when sending request to the server")
	// 	return
	// }

	// defer resp.Body.Close()
	// resp_body, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(resp.Status)
	// fmt.Println(string(resp_body))

	//return []byte{}
}
