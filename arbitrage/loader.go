package arbitrage

import (
	"context"
	"encoding/json"
	"os"
)

func GetOrderBook(ctx context.Context, exchange string) (*Exchange, error) {
	
    if err := ctx.Err(); err != nil {
        return nil, err
    }

	data, err := os.ReadFile("exchanges/" + exchange + ".json")
	if err != nil {
		return nil, err
	}

	var ex = new(Exchange)
	if err := json.Unmarshal(data, ex); err != nil {
		return nil, err
	}

	return ex, nil
}
