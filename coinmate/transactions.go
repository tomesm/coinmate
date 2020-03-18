package coinmate

import "time"

const (
	TransactionHistory = "/transactionHistory"
)

type APITransactionBody struct {
	offset int `json:"offset`
	limit int `json:"limit`
	sort string `json:"sort"`
	// timestampFrom int64 `json:"timestampFrom`
	clientId int64 `json:"clientIDd"`
	publicKey string `json:"publicKey`
	nonce int64 `json:"nonce"`
	signature `json:"signature"`

}

func (api *APIClient) GetTransactionHistory() []byte {
	
	body := APITransactionBody {
		offset: 0,
		limit: 10,
		sort: "ASC",
		clientId: api.ClientId,
		publicKey: api.Key,
		nonce: api.Nonce
		signature: api.createSignature,

	}

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

	return []byte{}
}
