package main

import (
	"design-patterns/structural/flyweight/factory"
	"design-patterns/structural/flyweight/models"
	"fmt"
)

func main() {
	iconFactory := factory.NewIconFactory()

	var markers []*models.MapMarker

	for i := 0; i < 5; i++ {
		markers = append(markers, models.NewMapMarker(12.9+float64(i)*0.01, 77.5, "Restaurant", iconFactory.Get("restaurant")))
	}
	for i := 0; i < 3; i++ {
		markers = append(markers, models.NewMapMarker(12.97+float64(i)*0.02, 77.6, "Hotel", iconFactory.Get("hotel")))
	}
	for i := 0; i < 4; i++ {
		markers = append(markers, models.NewMapMarker(13.0+float64(i)*0.015, 77.55, "Park", iconFactory.Get("park")))
	}

	for _, m := range markers {
		m.Render()
	}

	bytesIfUnique := len(markers) * 8 * 1024
	bytesWithFlyweight := iconFactory.Count() * 8 * 1024
	fmt.Printf("\nIcons created: %d, Markers: %d\n", iconFactory.Count(), len(markers))
	fmt.Printf("Memory if unique icons: %d bytes\n", bytesIfUnique)
	fmt.Printf("Memory with flyweight:  %d bytes\n", bytesWithFlyweight)
}
