package coinmate

func TestGetTransactionHistory(t *Testing.T) {
	c := NewAPIClient("CpmRVUJL0OGByT2otAfCKeeDdU6yfi6OzvnXcAwaHvEgit ", "", "6")
	c.Endpoint = "https://private-anon-4cb33c8b65-coinmate.apiary-mock.com/api"
}
