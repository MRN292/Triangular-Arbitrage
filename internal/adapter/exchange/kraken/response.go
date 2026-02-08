package kraken

type KrakenOrderBook struct {
	Error  []any                      `json:"error"`
	Result map[string]KrakenMarket `json:"result"`
}

type KrakenMarket struct {
	Asks [][2]string `json:"asks"`
	Bids [][2]string `json:"bids"`
}