package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"triangular-arbitrage/arbitrage"
	"triangular-arbitrage/helper"
)

func main() {

	userInputs, err := helper.GetInputs()
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Channel for saving outputs
	resultCh := make(chan *arbitrage.ArbitrageOpportunity, len(userInputs.ExchangesPaths))

	// Create and start worker pool
	workerPool := arbitrage.NewTriangularArbitrageWorker(
		ctx,
		userInputs.ExchangesPaths,
		userInputs.Pairs,
		userInputs.StartAmount,
		userInputs.MinProfit,
		resultCh,
	)
	workerPool.Start()

	go func() {
		workerPool.Wait()
		close(resultCh)
	}()

	// Collect results
	var results []*arbitrage.ArbitrageOpportunity
	for res := range resultCh {
		results = append(results, res)
	}

	if len(results) == 0 {
		fmt.Println("No arbitrage opportunity found in any exchange")
		return
	}

	helper.PrettyPrint(results)
}
