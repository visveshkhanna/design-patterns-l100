package main

import (
	"design-patterns/behavioural/strategy/context"
	"design-patterns/behavioural/strategy/strategies"
)

func main() {

	r := context.NewRide(&strategies.Economy{})
	r.PrintQuote(12)

	r.SetStrategy(&strategies.Premium{})
	r.PrintQuote(12)
}
