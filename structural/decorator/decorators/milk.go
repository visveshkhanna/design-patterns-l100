package decorators

import "design-patterns/structural/decorator/interfaces"

type Milk struct {
	base interfaces.IBeverage
}

func WithMilk(b interfaces.IBeverage) *Milk {
	return &Milk{base: b}
}

func (m *Milk) Description() string {
	return m.base.Description() + ", Milk"
}

func (m *Milk) Cost() int {
	return m.base.Cost() + 30
}

var _ interfaces.IBeverage = (*Milk)(nil)
