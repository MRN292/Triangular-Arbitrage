package kraken

import (
	"strconv"
	"triangular-arbitrage/internal/domain"
)

func (dto KrakenOrderBook) ToDomain() (output *domain.TopOfBook,err error) {
	
	for _, MarketBook := range dto.Result {

		bestAskAmount, err := strconv.ParseFloat(MarketBook.Asks[0][1], 64)
		if err != nil {
			return output, err
		}
		bestAskPrice, err := strconv.ParseFloat(MarketBook.Asks[0][0], 64)
		if err != nil {
			return output, err
		}
		bestBidsAmount, err := strconv.ParseFloat(MarketBook.Bids[0][1], 64)
		if err != nil {
			return output, err
		}
		bestBidsPrice, err := strconv.ParseFloat(MarketBook.Bids[0][0], 64)
		if err != nil {
			return output, err
		}

		output = &domain.TopOfBook{
			BestAsk: &domain.Side{Price: bestAskPrice, Amount: bestAskAmount},
			BestBid: &domain.Side{Price: bestBidsPrice, Amount: bestBidsAmount},
		}
	}

	return output, nil
}
