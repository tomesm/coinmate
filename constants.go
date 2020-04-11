package coinmate

type Endpoints struct{}

func (e Endpoints) baseURL() string {
	return "https://coinmate.io/api/"
}

func (e Endpoints) mockBaseURL() string {
	return "https://private-anon-5fc444e684-coinmate.apiary-mock.com"
}

func (e Endpoints) tradingPairs() string {
	return "/tradingPairs"
}

func (e Endpoints) transactionHistory() string {
	return "/transactionHistory"
}
