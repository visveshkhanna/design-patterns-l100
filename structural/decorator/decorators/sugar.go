package decorators

import "design-patterns/structural/decorator/interfaces"

type Sugar struct {
	base interfaces.IBeverage
}

func WithSugar(b interfaces.IBeverage) *Sugar {
	return &Sugar{base: b}
}

func (s *Sugar) Description() string {
	return s.base.Description() + ", Sugar"
}

func (s *Sugar) Cost() int {
	return s.base.Cost() + 10
}

var _ interfaces.IBeverage = (*Sugar)(nil)
