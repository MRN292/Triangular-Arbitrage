package domain

import "time"

type ArbitrageOpportunity struct {
	Exchange      string
	Path          string
	StartAmount   float64
	EndAmount     float64
	Profit        float64
	ProfitPercent float64
	Timestamp     time.Time
}