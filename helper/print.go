package helper

import (
	"fmt"
	"strings"
	"triangular-arbitrage/internal/domain"
)

func PrettyPrint(results []*domain.ArbitrageOpportunity) {
	if len(results) == 0 {
		fmt.Println("No arbitrage opportunities found.")
		return
	}

	fmt.Println(strings.Repeat("=", 60))

	for i, r := range results {
		fmt.Printf("#%d\n", i+1)
		fmt.Printf("Exchange       : %s\n", r.Exchange)
		fmt.Printf("Triangle path  : %s\n", r.Path)
		fmt.Printf("Start Amount   : %.2f\n", r.StartAmount)
		fmt.Printf("Final Amount   : %.2f\n", r.EndAmount)
		fmt.Printf("Profit         : %.2f\n", r.Profit)
		fmt.Printf("Profit Percent : %.2f%%\n", r.ProfitPercent)
		fmt.Printf("Timestamp      : %s\n", r.Timestamp.Format("2006-01-02 15:04:05"))
		fmt.Println(strings.Repeat("-", 60))
	}
}
