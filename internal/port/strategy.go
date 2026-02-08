package port

import (
	"context"
	"triangular-arbitrage/internal/domain"
)

type Strategy interface {
	Name() string
	FindArbitrageOpportunities(ctx context.Context, exchange Exchange, orderBooks *domain.OrderBooks) (*domain.ArbitrageOpportunity, error)
}
