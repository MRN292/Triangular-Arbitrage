package main

import (
	"context"
	"fmt"
	"log"
	"triangular-arbitrage/config"
	"triangular-arbitrage/helper"
	"triangular-arbitrage/internal/adapter/exchange/kraken"
	"triangular-arbitrage/internal/adapter/strategy/triangular"
	"triangular-arbitrage/internal/registry"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in loading configs %v", err)
	}

	userInputs, err := helper.GetInputs()
	if err != nil {
		log.Fatal(err)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	ctx := context.Background()

	kraken := kraken.NewKraken(cfg)

	exRegistry := registry.NewExchangeRegistry(kraken)
	ex, ok := exRegistry.Get("kraken")
	if !ok {
		log.Printf("exhange not found : %s", "kraken")
	}
	if err := ex.Ping(ctx); err != nil {
		log.Printf("Error in connection %v", err)
	}

	tri := triangular.NewTriangular(userInputs.Pairs[0], userInputs.Pairs[1], userInputs.Pairs[2], userInputs.StartAmount, userInputs.MinProfit)
	strgReg := registry.NewStrategyRegistry(tri)
	st, ok := strgReg.Get("triangular")
	if !ok {
		log.Printf("Strategy not found : %s", "trangular")
	}

	orderBooks, err := ex.GetOrderBook(ctx, userInputs.Pairs)

	opp, err := st.FindArbitrageOpportunities(ctx, ex, orderBooks)
	if err != nil {
		log.Printf("strategy error: %v", err)
	}

	fmt.Println(opp.EndAmount)

}
