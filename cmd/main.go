package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"triangular-arbitrage/config"
	"triangular-arbitrage/internal/adapter/exchange/kraken"
	"triangular-arbitrage/internal/app"
	"triangular-arbitrage/internal/registry"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in loading configs %v", err)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	ctx := context.Background()

	kraken := kraken.NewKraken(cfg)
	exRegistry := registry.NewExchangeRegistry(kraken)

	// i implemented regitry for strategy but not usful
	trangularCore := app.NewTrangularCore(ctx, exRegistry)

	results := trangularCore.TrangularCoreRun()

	if len(results) == 0 {
		log.Printf("no results")
		return
	}
	for _, r := range results {
		if r.Err != nil {
			log.Printf("[%s] %v", r.Exchange, r.Err)
			continue
		}
		if r.ArbitrageOpp == nil {
			log.Printf("[%s] no opportunity", r.Exchange)
			continue
		}
		fmt.Println("================================")
		fmt.Printf("Exchange: %s\nPath: %s\nEndAmount: %.8f\nProfit: %.8f\nProfitPercent: %.4f\nTimeStamp: %s",
			r.ArbitrageOpp.Exchange,
			r.ArbitrageOpp.Path,
			r.ArbitrageOpp.EndAmount,
			r.ArbitrageOpp.Profit,
			r.ArbitrageOpp.ProfitPercent,
			time.Now().Format(time.RFC3339),
		)
	}
}
