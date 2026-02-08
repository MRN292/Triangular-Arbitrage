package port

import (
	"triangular-arbitrage/internal/domain"
)

type Strategy interface {
	Name() string
	FindArbitrageOpportunities(exchange Exchange, orderBooks *domain.OrderBooks) (*domain.ArbitrageOpportunity, error)
}
