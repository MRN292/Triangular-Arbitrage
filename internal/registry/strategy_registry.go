package registry

import "triangular-arbitrage/internal/port"

type StrategyRegistry struct {
	strgMap map[string]port.Strategy
}

func NewStrategyRegistry(strategies ...port.Strategy) *StrategyRegistry {
	registry := &StrategyRegistry{
		strgMap: make(map[string]port.Strategy, len(strategies)),
	}
	for _, strategy := range strategies {
		registry.strgMap[strategy.Name()] = strategy
	}
	return registry
}

func (r *StrategyRegistry) Get(name string) (port.Strategy, bool) {
	s, ok := r.strgMap[name]
	return s, ok
}

func (r *StrategyRegistry) List() (output []string) {

	for strg := range r.strgMap {
		output = append(output, strg)
	}
	return output
}
