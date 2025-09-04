package components

import (
	"design-patterns/structural/composite/interfaces"
	"fmt"
)

type Box struct {
	name       string
	components []interfaces.IComponent
}

func NewBox(name string) *Box {
	return &Box{name: name, components: make([]interfaces.IComponent, 0)}
}

func (b *Box) Add(component interfaces.IComponent) {
	b.components = append(b.components, component)
}

func (b *Box) Remove(index int) {
	if index >= 0 && index < len(b.components) {
		b.components = append(b.components[:index], b.components[index+1:]...)
	}
}

func (b *Box) Price() int {
	total := 0
	for _, c := range b.components {
		total += c.Price()
	}
	return total
}

func (b *Box) Print(indent string) {
	fmt.Printf("%s+ %s (subtotal: $%d)\n", indent, b.name, b.Price())
	deeper := indent + "  "
	for _, c := range b.components {
		c.Print(deeper)
	}
}
