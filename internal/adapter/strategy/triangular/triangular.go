package triangular

import (
	"context"
	"fmt"
	"time"
	"triangular-arbitrage/internal/domain"
	"triangular-arbitrage/internal/port"
)

type Triangular struct {
	StartAmount   float64
	MinimumProfit float64
	Pair1         string
	Pair2         string
	Pair3         string
}

func NewTriangular(pair1, pair2, pair3 string, startAmount, minimumProfit float64) *Triangular {
	return &Triangular{
		Pair1:         pair1,
		Pair2:         pair2,
		Pair3:         pair3,
		StartAmount:   startAmount,
		MinimumProfit: minimumProfit,
	}
}
func (t *Triangular) Name() string {
	return "triangular"
}

func (t *Triangular) FindArbitrageOpportunities(ctx context.Context, exchange port.Exchange, orderBooks *domain.OrderBooks) (*domain.ArbitrageOpportunity, error) {
	_ = ctx

	ob1, ok1 := orderBooks.Markets[t.Pair1]
	ob2, ok2 := orderBooks.Markets[t.Pair2]
	ob3, ok3 := orderBooks.Markets[t.Pair3]

	if !ok1 || !ok2 || !ok3 {
		return nil, fmt.Errorf("Pair not found in order book of %s", exchange.Name())
	}

	// Sell step
	firstStep := t.StartAmount / ob1.BestAsk.Price
	// Checking Liquidity
	if firstStep > ob1.BestAsk.Amount {
		return nil, fmt.Errorf("Liquidity not enough in %s", exchange.Name())
	}

	// Sell step
	secondStep := firstStep / ob2.BestAsk.Price
	// Checking Liquidity
	if secondStep > ob2.BestAsk.Amount {
		return nil, fmt.Errorf("Liquidity not enough in %s", exchange.Name())
	}

	// Buy step
	// Checking Liquidity
	if secondStep > ob3.BestBid.Amount {
		return nil, fmt.Errorf("Liquidity not enough in %s", exchange.Name())
	}
	finalAmount := secondStep * ob3.BestBid.Price
	if finalAmount <= t.StartAmount {
		return nil, fmt.Errorf("No profit in %s", exchange.Name())
	}

	profit := finalAmount - t.StartAmount
	profitPercent := (profit / t.StartAmount) * 100

	if profitPercent < t.MinimumProfit {
		return nil, fmt.Errorf("Profit presentage is less than %f in %s", t.MinimumProfit, exchange.Name())
	}

	return &domain.ArbitrageOpportunity{
		Exchange:      exchange.Name(),
		Path:          fmt.Sprintf("%s -> %s -> %s", t.Pair1, t.Pair2, t.Pair3),
		StartAmount:   t.StartAmount,
		EndAmount:     finalAmount,
		Profit:        profit,
		ProfitPercent: profitPercent,
		Timestamp:     time.Now(),
	}, nil

}
