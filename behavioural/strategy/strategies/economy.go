package strategies

type Economy struct{}

func (e Economy) Price(distanceKm float64) float64 {
	return 1.0 + 0.8*distanceKm
}
