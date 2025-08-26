package main

// Sugar is a decorator that adds sugar to a beverage.
type Sugar struct {
	base Beverage
}

func WithSugar(b Beverage) *Sugar {
	return &Sugar{base: b}
}

func (s *Sugar) Description() string {
	return s.base.Description() + ", Sugar"
}

func (s *Sugar) Cost() int {
	return s.base.Cost() + 10
}
