package main

// Milk is a decorator that adds milk to a beverage.
type Milk struct {
	base Beverage
}

func WithMilk(b Beverage) *Milk {
	return &Milk{base: b}
}

func (m *Milk) Description() string {
	return m.base.Description() + ", Milk"
}

func (m *Milk) Cost() int {
	return m.base.Cost() + 30
}
