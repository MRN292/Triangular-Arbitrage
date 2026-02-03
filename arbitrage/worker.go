package arbitrage

import (
	"context"
	"fmt"
	"sync"
)

// jobsChannelType holds the job details for each worker
type jobsChannelType struct {
	exchanges     string
	minimumProfit float64
	startAmount   float64
	pair1         string
	pair2         string
	pair3         string
}

// TriangularArbitrageWorker manages a pool of workers to check for triangular arbitrage opportunities
type TriangularArbitrageWorker struct {
	ctx           context.Context
	workersNumber int
	jobsChannel   chan *jobsChannelType
	resultChannel chan *ArbitrageOpportunity
	wg            sync.WaitGroup
}

func NewTriangularArbitrageWorker(ctx context.Context, exchangePaths []string, pairs []string, startAmount, minProfit float64, resultChannel chan *ArbitrageOpportunity) *TriangularArbitrageWorker {
	jobsChannel := make(chan *jobsChannelType, len(exchangePaths))

	// numbers of workers is set to 30% of total exchanges
	workersNumber := int(float64(len(exchangePaths)) * 0.3)
	if workersNumber < 1 {
		workersNumber = 1
	}

	// Add all jobs to the channel
	for _, path := range exchangePaths {
		jobsChannel <- &jobsChannelType{
			exchanges:     path,
			minimumProfit: minProfit,
			startAmount:   startAmount,
			pair1:         pairs[0],
			pair2:         pairs[1],
			pair3:         pairs[2],
		}
	}
	// closing channel to preventing workers to wait forever , in the other hand time out is set in contex
	close(jobsChannel)

	worker := &TriangularArbitrageWorker{
		ctx:           ctx,
		workersNumber: workersNumber,
		jobsChannel:   jobsChannel,
		resultChannel: resultChannel,
	}
	return worker
}

// Start starts the worker pool
func (w *TriangularArbitrageWorker) Start() {
	for i := 1; i <= w.workersNumber; i++ {
		w.wg.Add(1)
		go w.runWorker(i)
	}
}

// Wait waits for all workers to finish
func (w *TriangularArbitrageWorker) Wait() {
	w.wg.Wait()
}

func (w *TriangularArbitrageWorker) runWorker(workerID int) {
	defer w.wg.Done()

	for job := range w.jobsChannel {
		select {
		case <-w.ctx.Done():
			return
		default:
			w.processJob(workerID, job)
		}
	}
}

// processJob processes a single job, using FindArbitrageOpportunities() function in it
func (w *TriangularArbitrageWorker) processJob(workerID int, job *jobsChannelType) {
	ex, err := GetOrderBook(w.ctx, job.exchanges)
	if err != nil {
		fmt.Printf("Worker %d: Failed to load %s, error: %v\n", workerID, job.exchanges, err)
		return
	}

	// Check triangular arbitrage opportunity
	tra := TriangularArbitrage{
		Exchange:      ex,
		MinimumProfit: job.minimumProfit,
		StartAmount:   job.startAmount,
		Pair1:         job.pair1,
		Pair2:         job.pair2,
		Pair3:         job.pair3,
	}

	output, ok := tra.FindArbitrageOpportunities()
	if !ok {
		return
	}

	select {
	case <-w.ctx.Done():
		return
	case w.resultChannel <- output:
	}
}
