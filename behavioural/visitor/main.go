package main

import (
	"design-patterns/behavioural/visitor/elements"
	"design-patterns/behavioural/visitor/interfaces"
	"design-patterns/behavioural/visitor/visitors"
)

func main() {

	zones := []interfaces.IElement{
		elements.NewForest(120),
		elements.NewCity(45),
	}

	printer := visitors.NewMapPrinter()

	for _, zone := range zones {
		zone.Accept(printer)
	}
}
