package arbitrage

import "time"

type Exchange struct {
	Name    string            `json:"name"`
	Markets map[string]Market `json:"markets"`
}

type Market struct {
	Bids []Order `json:"bids"`
	Asks []Order `json:"asks"`
}

type Order struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

type ArbitrageOpportunity struct {
	Exchange      string
	Path          string
	StartAmount   float64
	EndAmount     float64
	Profit        float64
	ProfitPercent float64
	Timestamp     time.Time
}
