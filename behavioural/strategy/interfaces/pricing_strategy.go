package interfaces

type PricingStrategy interface {
	Price(distanceKm float64) float64
}
