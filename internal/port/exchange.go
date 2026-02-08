package port

import (
	"context"
	"triangular-arbitrage/internal/domain"
)

type Exchange interface {
	Name() string
	Ping(ctx context.Context) error
	GetOrderBook(ctx context.Context, symbols []string) (*domain.OrderBooks, error)
}
