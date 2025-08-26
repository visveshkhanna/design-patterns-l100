package main

import "fmt"

type PricingStrategy interface {
	price(distanceKm float64) float64
}

type Economy struct{}

type Premium struct{}

func (Economy) price(distanceKm float64) float64 { return 1.0 + 0.8*distanceKm }

func (Premium) price(distanceKm float64) float64 { return 3.0 + 1.2*distanceKm }

type Ride struct{ strategy PricingStrategy }

func (r *Ride) setStrategy(st PricingStrategy) { r.strategy = st }

func (r *Ride) quote(distanceKm float64) float64 { return r.strategy.price(distanceKm) }

func (r *Ride) printQuote(distanceKm float64) { fmt.Println("$", r.quote(distanceKm)) }
