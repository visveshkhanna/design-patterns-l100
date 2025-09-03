package context

import (
	"design-patterns/behavioural/strategy/interfaces"
	"fmt"
)

type Ride struct {
	strategy interfaces.PricingStrategy
}

func NewRide(strategy interfaces.PricingStrategy) *Ride {
	return &Ride{strategy: strategy}
}

func (r *Ride) SetStrategy(strategy interfaces.PricingStrategy) {
	r.strategy = strategy
}

func (r *Ride) Quote(distanceKm float64) float64 {
	return r.strategy.Price(distanceKm)
}

func (r *Ride) PrintQuote(distanceKm float64) {
	fmt.Println("$", r.Quote(distanceKm))
}
