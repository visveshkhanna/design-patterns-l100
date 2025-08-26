package main

import "fmt"

// MapMarker holds extrinsic state (coordinates, label) and references a shared Icon.
type MapMarker struct {
	lat   float64
	lng   float64
	label string
	icon  *Icon
}

func NewMapMarker(lat, lng float64, label string, icon *Icon) *MapMarker {
	return &MapMarker{lat: lat, lng: lng, label: label, icon: icon}
}

func (m *MapMarker) Render() {
	fmt.Printf("(%.4f, %.4f) %s [%s]\n", m.lat, m.lng, m.label, m.icon.name)
}
