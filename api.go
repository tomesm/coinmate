package coinmate

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type APIClient struct {
	Key      string
	Secret   string
	ClientId string
	Endpoint string
}

func NewAPIClient(apiKey, apiSecret, clientId string) *APIClient {
	return &APIClient{
		Key:      apiKey,
		Secret:   apiSecret,
		ClientId: clientId,
		Endpoint: "",
	}
}

func (api *APIClient) Execute(method string, path string, body interface{}, result interface{}) error {
	client := &http.Client{}
	// body and query and data can be empty
	// if body nil an empty string is assigned to the query. no need to check if body is nil
	query := createURLValues(body).Encode()
	data := []byte(query)

	if api.Endpoint == "" {
		api.Endpoint = Endpoints{}.baseURL()
	}
	// it works even if data is empty for GET method
	request, err := http.NewRequest(method, api.Endpoint+path, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("ERROR creating request")
		return nil
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.URL.RawQuery = query
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(response.Status)
		log.Fatal(err)
		return err
	}
	fmt.Println(response.Status)
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (api *APIClient) Signature(nonce string) string {
	message := nonce + api.ClientId + api.Key
	key := []byte(api.Secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	signature := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(signature)
}

func (api *APIClient) Nonce() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
}
