package main

import "fmt"

type Box struct {
	name       string
	components []Component
}

func NewBox(name string) *Box {
	return &Box{name: name, components: make([]Component, 0)}
}

func (b *Box) Add(component Component) {
	b.components = append(b.components, component)
}

func (b *Box) RemoveByName(name string) {
	filtered := make([]Component, 0, len(b.components))
	for _, c := range b.components {
		switch v := c.(type) {
		case *Item:
			if v.name != name {
				filtered = append(filtered, c)
			}
		case *Box:
			if v.name != name {
				filtered = append(filtered, c)
			}
		default:
			filtered = append(filtered, c)
		}
	}
	b.components = filtered
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
