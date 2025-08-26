package main

import (
	"fmt"
)

func main() {
	factory := NewIconFactory()

	markers := []*MapMarker{}
	// Create many markers sharing a few icon types
	for i := 0; i < 5; i++ {
		markers = append(markers, NewMapMarker(12.9+float64(i)*0.01, 77.5, "Restaurant", factory.Get("restaurant")))
	}
	for i := 0; i < 3; i++ {
		markers = append(markers, NewMapMarker(12.97+float64(i)*0.02, 77.6, "Hotel", factory.Get("hotel")))
	}
	for i := 0; i < 4; i++ {
		markers = append(markers, NewMapMarker(13.0+float64(i)*0.015, 77.55, "Park", factory.Get("park")))
	}

	for _, m := range markers {
		m.Render()
	}

	// Memory illustration: without flyweight, each marker would carry icon bytes
	bytesIfUnique := len(markers) * 8 * 1024
	bytesWithFlyweight := factory.Count() * 8 * 1024
	fmt.Printf("\nIcons created: %d, Markers: %d\n", factory.Count(), len(markers))
	fmt.Printf("Memory if unique icons: %d bytes\n", bytesIfUnique)
	fmt.Printf("Memory with flyweight:  %d bytes\n", bytesWithFlyweight)
}
