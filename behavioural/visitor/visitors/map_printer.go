package visitors

import (
	"design-patterns/behavioural/visitor/interfaces"
	"fmt"
)

type MapPrinter struct{}

func NewMapPrinter() *MapPrinter {
	return &MapPrinter{}
}

func (mp *MapPrinter) VisitForest(forest interfaces.IForestElement) {
	fmt.Printf("Forest: %d trees\n", forest.GetTrees())
}

func (mp *MapPrinter) VisitCity(city interfaces.ICityElement) {
	fmt.Printf("City: %d buildings\n", city.GetBuildings())
}
