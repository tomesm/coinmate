package coinmate

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
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
	if api.Endpoint == "" {
		api.Endpoint = Endpoints{}.baseURL()
	}
	client := &http.Client{}
	data := url.Values{}
	if body != nil {
		data = createURLData(body)
	}
	request, err := http.NewRequest(method, api.Endpoint+path, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("ERROR creating request")
		return err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error in response")
		fmt.Println(response.Status)
		return err
	}
	err = decodeJSONBody(response, &result)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return nil
}

func (api *APIClient) signature(nonce string) string {
	message := nonce + api.ClientId + api.Key
	key := []byte(api.Secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	signature := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(signature)
}

func (api *APIClient) nonce() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
}

func (api *APIClient) convertString(str string) int64 {
	conv, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatal("Cannot convert str to int")
	}
	return conv
}
