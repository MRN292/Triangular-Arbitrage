package helper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UserInput struct {
	ExchangesPaths []string
	Pairs          []string
	StartAmount    float64
	MinProfit      float64
}

func GetInputs() (*UserInput, error) {
	reader := bufio.NewReader(os.Stdin)

	// Get exchanges from user through console
	fmt.Print("Enter exchanges (comma separated, example: binance,forex): ")
	exInput, _ := reader.ReadString('\n')
	exInput = strings.ToLower(strings.TrimSpace(exInput))
	exFiles := strings.Split(exInput, ",")
	if len(exInput) == 0 {
		return nil, fmt.Errorf("invalid exchanges")
	}
	for i := range exFiles {
		exFiles[i] = strings.TrimSpace(exFiles[i])
	}

	// Get pairs from user through console
	fmt.Print("Enter 3 pairs (comma separated, example: BTC/USDT,ETH/BTC,ETH/USDT): ")
	pairInput, _ := reader.ReadString('\n')
	pairInput = strings.TrimSpace(pairInput)
	pairs := strings.Split(pairInput, ",")
	if len(pairs) != 3 {
		return nil, fmt.Errorf("invalid pairs")
	}

	// Get user's start amount
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

	return &UserInput{
		ExchangesPaths: exFiles,
		Pairs:          pairs,
		StartAmount:    startAmount,
		MinProfit:      minProfit,
	}, nil
}
