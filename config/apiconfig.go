package config

type API struct {
	BinancePingUrl      string `env:"BINANCE_PING_URL" required:"true"`
	BinanceOrderBookUrl string `env:"BINANCE_ORDER_BOOK_URL" required:"true"`
	
	KrakenPingUrl      string `env:"KRAKEN_PING_URL" required:"true"`
	KrakenOrderBookUrl string `env:"KRAKEN_ORDER_BOOK_URL" required:"true"`
}
