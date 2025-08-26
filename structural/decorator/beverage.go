package main

// Beverage is the component interface decorated by add-ons like Milk and Sugar.
type Beverage interface {
	Description() string
	Cost() int
}
