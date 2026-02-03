# Introduction

Until recently, I had very little understanding or practical perspective on cryptocurrencies and how digital asset markets operate.</br>
Through my recent learning experience, I have gained a solid understanding of several key concepts in the cryptocurrency</br>
domain, including market structures, order books, and triangular arbitrage.

# Here is a summary of what I have learned:

1. **Cryptocurrency Basics**

    * Cryptocurrencies are digital, decentralized assets that operate on blockchain technology.
    * Trading occurs in pairs, such as BTC/USDT or ETH/BTC.
    * Exchanges like Binance or Coinbase facilitate buying and selling.
    * Orders specify the price and amount for buying or selling assets.

2. **Order Book**

    * An order book is a live ledger of buy (bids) and sell (asks) orders.
    * The top of the book represents the best bid and ask prices.
    * Price refers to the cost per unit of the asset.
    * Amount indicates the maximum quantity available at that price.
    * Total value = price × amount.

3. Market Factors

    * Liquidity determines how easily an asset can be bought or sold without moving the price.
    * Slippage is the difference between the expected execution price and the actual price.
    * Spread is the difference between the best bid and best ask.

4. Triangular Arbitrage

    * This involves exploiting price differences among three trading pairs within a single exchange.
    * Example path: USDT → BTC → ETH → USDT
    * Profit is calculated by comparing the final amount in the starting currency with the initial amount.
    * In simplified exercises, only the top-of-book prices are considered, ignoring liquidity constraints.


# Implementation

## Folder Structure

```
cmd/
    main.go              # Application entry point
arbitrage/
    types.go             # Data structures (Exchange, Market, Order, ArbitrageOpportunity)
    triangular.go        # Arbitrage calculation logic
    loader.go            # JSON data loader (simulates API calls)
    worker.go            # Concurrent worker pool
helper/
    input.go             # User input handling
    print.go             # Result formatting
exchanges/
    binance.json         # Mock order book data
```

## Key Functions

**`FindArbitrageOpportunities()`** (arbitrage/worker.go)
- Entry point that orchestrates concurrent exchange processing
- Creates worker pool (30% of exchanges, min 1) and distributes work via channels
- Uses `context.Context` for 5-second timeout and cancellation
- Collects results from all workers using `sync.WaitGroup`

**`calculateTriangularArbitrage()`** (arbitrage/triangular.go)
- Executes the three-step arbitrage calculation:
  1. StartAmount / Pair1.AskPrice = Amount1
  2. Amount1 / Pair2.AskPrice = Amount2
  3. Amount2 * Pair3.BidPrice = FinalAmount
- Validates liquidity at each step (checks if order amount ≥ required amount)
- Returns opportunity if profit exceeds minimum threshold

**`GetOrderBook()`** (arbitrage/loader.go)
- Loads order book data from JSON files (simulates API calls)
- Parses market data including bids/asks for each trading pair

## Usage

```
go run cmd/main.go
```

**Example:**
```
Enter exchanges: binance
Enter 3 pairs: BTC/USDT,ETH/BTC,ETH/USDT
Enter start amount: 1000
Enter minimum profit %: 0.5
```