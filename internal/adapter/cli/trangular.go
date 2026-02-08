package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TriangularUserInput struct {
	Exchanges   []string
	Pairs       []string
	StartAmount float64
	MinProfit   float64
}

func GetTriangularUserInput() (*TriangularUserInput, error) {
	r := os.Stdin

	reader := bufio.NewReader(r)

	// Get exchanges from user through console
	fmt.Print("Enter exchanges (comma separated, example: kraken,binance): ")
	exInput, _ := reader.ReadString('\n')
	exInput = strings.ToLower(strings.TrimSpace(exInput))
	exchanges := strings.Split(exInput, ",")
	if len(exInput) == 0 {
		return nil, fmt.Errorf("invalid exchanges")
	}
	for i := range exchanges {
		exchanges[i] = strings.TrimSpace(exchanges[i])
	}

	fmt.Print("Enter 3 pairs (comma separated, example: BTC/USDT,ETH/BTC,ETH/USDT): ")
	pairInput, _ := reader.ReadString('\n')
	pairInput = strings.TrimSpace(pairInput)
	pairs := strings.Split(pairInput, ",")
	if len(pairs) != 3 {
		return nil, fmt.Errorf("invalid pairs")
	}

	fmt.Print("Enter start amount in dollar (example: 1000): ")
	var startAmount float64
	_, err := fmt.Scanf("%f\n", &startAmount)
	if err != nil {
		return nil, fmt.Errorf("invalid start amount")
	}

	// Get minimum profit from user
	fmt.Print("Enter minimum profit percent (example: 1.5): ")
	var minProfit float64
	_, err = fmt.Scanf("%f\n", &minProfit)
	if err != nil {
		return nil, fmt.Errorf("invalid profit percent")
	}

	return &TriangularUserInput{
		Exchanges:   exchanges,
		Pairs:       pairs,
		StartAmount: startAmount,
		MinProfit:   minProfit,
	}, nil
}
