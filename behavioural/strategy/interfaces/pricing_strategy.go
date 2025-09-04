package interfaces

type IPricingStrategy interface {
	Price(distanceKm float64) float64
}
