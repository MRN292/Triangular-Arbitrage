package app

import (
	"context"
	"fmt"
	"log"
	"sync"
	"triangular-arbitrage/internal/adapter/cli"
	"triangular-arbitrage/internal/adapter/strategy/triangular"
	"triangular-arbitrage/internal/domain"
	"triangular-arbitrage/internal/registry"
)

type TriangularCore struct {
	exchangeRegistry   *registry.ExchangeRegistry
	ctx                context.Context
	userInput          *cli.TriangularUserInput
	resultChannel      chan TriangularWorkerResult
	wg                 sync.WaitGroup
	triangularStrategy *triangular.Triangular
}

type TriangularWorkerResult struct {
	Exchange     string
	ArbitrageOpp *domain.ArbitrageOpportunity
	Err          error
}

func NewTrangularCore(ctx context.Context, exchangeRegistry *registry.ExchangeRegistry) *TriangularCore {
	return &TriangularCore{
		ctx:              ctx,
		exchangeRegistry: exchangeRegistry,
	}
}

func (t *TriangularCore) TrangularCoreRun() (output []TriangularWorkerResult) {
	userInput, err := cli.GetTriangularUserInput()
	if err != nil {
		log.Fatal(err)
	}
	t.userInput = userInput

	t.triangularStrategy = triangular.NewTriangular(
		t.userInput.Pairs[0],
		t.userInput.Pairs[1],
		t.userInput.Pairs[2],
		t.userInput.StartAmount,
		t.userInput.MinProfit,
	)

	t.resultChannel = make(chan TriangularWorkerResult, len(t.userInput.Exchanges))

	for _, exchange := range t.userInput.Exchanges {
		t.wg.Add(1)
		go t.trangularWorker(exchange)
	}

	t.wg.Wait()
	close(t.resultChannel)

	for r := range t.resultChannel {
		output = append(output, r)
	}

	return output
}

func (t *TriangularCore) trangularWorker(exchange string) {
	defer t.wg.Done()

	if err := t.ctx.Err(); err != nil {
		t.resultChannel <- TriangularWorkerResult{Exchange: exchange, Err: err}
		return
	}

	ex, ok := t.exchangeRegistry.Get(exchange)
	if !ok {
		t.resultChannel <- TriangularWorkerResult{
			Exchange: exchange,
			Err:      fmt.Errorf("exchange not found"),
		}
		return
	}

	if err := t.ctx.Err(); err != nil {
		t.resultChannel <- TriangularWorkerResult{Exchange: ex.Name(), Err: err}
		return
	}

	if err := ex.Ping(t.ctx); err != nil {
		t.resultChannel <- TriangularWorkerResult{Exchange: ex.Name(), Err: err}
		return
	}

	orderBook, err := ex.GetOrderBook(t.ctx, t.userInput.Pairs)
	if err != nil {
		t.resultChannel <- TriangularWorkerResult{Exchange: ex.Name(), Err: err}
		return
	}

	opp, err := t.triangularStrategy.FindArbitrageOpportunities(ex, orderBook)
	if err != nil {
		t.resultChannel <- TriangularWorkerResult{Exchange: ex.Name(), Err: err}
		return
	}

	t.resultChannel <- TriangularWorkerResult{Exchange: ex.Name(), ArbitrageOpp: opp}
}
