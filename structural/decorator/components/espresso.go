package components

import "design-patterns/structural/decorator/interfaces"

type Espresso struct{}

func NewEspresso() *Espresso {
	return &Espresso{}
}

func (e *Espresso) Description() string {
	return "Espresso"
}

func (e *Espresso) Cost() int {
	return 200
}

var _ interfaces.IBeverage = (*Espresso)(nil)
