package main

const (
	ENDOPOINT = "https://coinmate.io/api/"
)

type API struct {
	Key      string
	Secret   string
	Endpoint string
}

// NewAPIClient : Constructor for coinmate api client
func NewAPIClient(key, secret string) *API {
	return *API{
		Key:      key,
		Secret:   secret,
		Endpoint: ENDPOINT,
	}
}
