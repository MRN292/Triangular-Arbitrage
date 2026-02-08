package registry

import "triangular-arbitrage/internal/port"

type ExchangeRegistry struct {
	exMap map[string]port.Exchange // place for savng ex :)
}

func NewExchangeRegistry(exchanges ...port.Exchange) *ExchangeRegistry {
	registry := &ExchangeRegistry{
		make(map[string]port.Exchange),
	}

	for _, exchange := range exchanges {
		registry.exMap[exchange.Name()] = exchange
	}
	return registry
}

func (r *ExchangeRegistry) Get(name string) (port.Exchange, bool) {
	ex, ok := r.exMap[name]
	return ex, ok
}

func (r *ExchangeRegistry) List() (output []string) {

	for ex := range r.exMap {
		output = append(output, ex)
	}
	return output
}
