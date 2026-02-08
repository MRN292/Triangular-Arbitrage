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

## Folder Structure (high level)

```
cmd/
    main.go                  # Entry point: wires config + exchange adapters + runs the app
config/
    config.go                # App configuration
    apiconfig.go             # Exchange/API configuration
helper/
    input.go                 # CLI input helpers
    print.go                 # Console output helpers
internal/
    app/                     # Use-cases / orchestration (runs the arbitrage flow)
    domain/                  # Core domain logic (opportunity calculation, models)
    adapter/                 # Exchange implementations (e.g., Kraken)
    port/                    # Interfaces (abstractions for adapters)
    registry/                # Wiring/registration of dependencies
    types/                   # Shared types used across layers
```

## Important parts (what to look at)

- **`cmd/main.go`**  
  Prints each found opportunity, and shows key fields like `Exchange`, `Path`, `EndAmount`, `Profit`, and `ProfitPercent`.

- **`internal/app`**  
  The application layer that coordinates the run: reading inputs/config, calling the exchange adapter(s), and executing the arbitrage detection.

- **`internal/domain`**  
  The core triangular-arbitrage logic: given market prices/order book data, compute the end amount and profit for a 3-leg path.

- **`internal/adapter/exchange/*`**  
  Concrete exchange integrations (e.g. Kraken). This is where request/response mapping and exchange-specific behavior lives.

## Usage

1) Create your local environment file:

```
cp .env.example .env
```

2) Run:

```
go run cmd/main.go
```