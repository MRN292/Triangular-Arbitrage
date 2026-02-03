package arbitrage

import (
	"fmt"
	"time"
)

type TriangularArbitrage struct {
	Exchange      *Exchange
	MinimumProfit float64
	StartAmount   float64
	Pair1         string
	Pair2         string
	Pair3         string
}

func (t *TriangularArbitrage) FindArbitrageOpportunities() (*ArbitrageOpportunity, bool) {

	m1, ok := t.Exchange.Markets[t.Pair1]
	if !ok || len(m1.Asks) == 0 {
		return nil, false
	}

	m2, ok := t.Exchange.Markets[t.Pair2]
	if !ok || len(m2.Asks) == 0 {
		return nil, false
	}

	m3, ok := t.Exchange.Markets[t.Pair3]
	if !ok || len(m3.Bids) == 0 {
		return nil, false
	}

	// Sell step
	firstStep := t.StartAmount / m1.Asks[0].Price
	// Checking Liquidity
	if firstStep > m1.Asks[0].Amount {
		return nil, false
	}

	// Sell step
	secondStep := firstStep / m2.Asks[0].Price
	// Checking Liquidity
	if secondStep > m2.Asks[0].Amount {
		return nil, false
	}

	// Buy step
	// Checking Liquidity
	if secondStep > m3.Bids[0].Amount {
		return nil, false
	}
	finalAmount := secondStep * m3.Bids[0].Price
	if finalAmount <= t.StartAmount {
		return nil, false
	}

	profit := finalAmount - t.StartAmount
	profitPercent := (profit / t.StartAmount) * 100

	if profitPercent < t.MinimumProfit {
		return nil, false
	}
	return &ArbitrageOpportunity{
		Exchange:      t.Exchange.Name,
		Path:          fmt.Sprintf("%s -> %s -> %s", t.Pair1, t.Pair2, t.Pair3),
		StartAmount:   t.StartAmount,
		EndAmount:     finalAmount,
		Profit:        profit,
		ProfitPercent: profitPercent,
		Timestamp:     time.Now(),
	}, true
}
