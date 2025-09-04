package models

import (
	"design-patterns/structural/flyweight/flyweight"
	"fmt"
)

type MapMarker struct {
	lat   float64
	lng   float64
	label string
	icon  *flyweight.Icon
}

func NewMapMarker(lat, lng float64, label string, icon *flyweight.Icon) *MapMarker {
	return &MapMarker{lat: lat, lng: lng, label: label, icon: icon}
}

func (m *MapMarker) Render() {
	fmt.Printf("(%.4f, %.4f) %s [%s]\n", m.lat, m.lng, m.label, m.icon.GetName())
}

func (m *MapMarker) GetPosition() (float64, float64) {
	return m.lat, m.lng
}

func (m *MapMarker) GetLabel() string {
	return m.label
}
