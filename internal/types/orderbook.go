package types

type OrderBook struct {
	Markets map[string]MarketBook `json:"result"`
}

type MarketBook struct {
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}