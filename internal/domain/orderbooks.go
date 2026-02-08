package domain

type OrderBooks struct {
	Markets map[string]*TopOfBook
}

type TopOfBook struct {
	BestBid *Side
	BestAsk *Side
}

type Side struct {
	Price  float64
	Amount float64
}