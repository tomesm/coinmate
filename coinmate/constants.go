package coinmate

type Endpoints struct{}

func (e Endpoints) baseURL() string {
	return "https://coinmate.io/api/"
}

func (e Endpoints) tradingPairs() string {
	return "/tradingPairs"
}

func (e Endpoints) transactionHistory() string {
	return "/transactionHistory"
}
