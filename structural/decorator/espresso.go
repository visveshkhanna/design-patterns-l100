package main

// Espresso is a concrete component.
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
