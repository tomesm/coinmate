package coinmate

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Endpoint string

const (
	// Endpoints
	BaseURL = "https://coinmate.io/api"
)

type APIClient struct {
	Key      string
	Secret   string
	ClientId string
	Nonce    string
}

//  : Constructor for coinmate api client
func NewAPIClient(apiKey, apiSecret, clientId string) *APIClient {
	return &APIClient{
		Key:      apiKey,
		Secret:   apiSecret,
		ClientId: clientId,
	}
}

func (api *APIClient) Execute(method string, path Endpoint, body interface{}, result interface{}) error {
	client := &http.Client{}
	var bodyBuffer io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyBuffer = bytes.NewBuffer([]byte(data))
	}
	request, err := http.NewRequest(method, BaseURL+string(path), bodyBuffer)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return err
	}
	return nil
}

func (api *APIClient) createSignature() string {
	api.Nonce = string(time.Now().Format("20060102150405"))
	message := api.Nonce + api.ClientId + api.Key
	key := []byte(api.Secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	signature := hex.EncodeToString(h.Sum(nil))
	return signature

	// return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (api *APIClient) convertString(str string) int64 {
	conv, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal("Cannot convert str to int")
	}
	return conv
}

//  signature = "{}{}{}".format(nonce, self.clientId, self.publicApiKey)
//     dig = hmac.new(
//         self.privateApiKey,
//         msg=signature.encode('utf-8'),
//         digestmod=hashlib.sha256
//     ).hexdigest()
//     signature = dig.encode('utf-8')
//     signature = signature.upper()
//     return signature

// def createSignature(clientId, apiKey, privateKey, nonce): message = str(nonce) + str(clientId) + apiKey signature = hmac.new(privateKey, message, digestmod=hashlib.sha256).hexdigest()
