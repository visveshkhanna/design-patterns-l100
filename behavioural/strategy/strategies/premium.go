package strategies

type Premium struct{}

func (p Premium) Price(distanceKm float64) float64 {
	return 3.0 + 1.2*distanceKm
}
