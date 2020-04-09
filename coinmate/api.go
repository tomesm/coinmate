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
	Key       string
	Secret    string
	ClientId  string
	Endpoints Endpoints
}

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

//  : Constructor for coinmate api client
func NewAPIClient(apiKey, apiSecret, clientId string) *APIClient {
	return &APIClient{
		Key:       apiKey,
		Secret:    apiSecret,
		ClientId:  clientId,
		Endpoints: Endpoints{},
	}
}

func (api *APIClient) Execute(method string, path string, body interface{}, result interface{}) error {
	client := &http.Client{}
	data := url.Values{}

	if body != nil {
		data = createURLData(body)
		// data.Set("offset", "0")
		// data.Set("limit", "10")
		// data.Set("sort", "ASC")
		// data.Set("clientId", "33058")
		// data.Set("nonce", nonce)
		// data.Set("publicKey", "X_89yOm0Nj0CtGG3yScN5WPssbcKPRnEWanKaxAQgGs")
		// data.Set("signature", signature)
	}

	request, err := http.NewRequest(method, api.Endpoints.baseURL()+path, strings.NewReader(data.Encode()))
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

	// err = json.NewDecoder(response.Body).Decode(result)
	// if err != nil {
	// 	return err
	// }
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
