package visitors

import (
	"design-patterns/behavioural/visitor/interfaces"
	"fmt"
)

type MapPrinter struct{}

func NewMapPrinter() *MapPrinter {
	return &MapPrinter{}
}

func (mp *MapPrinter) VisitForest(forest interfaces.ForestElement) {
	fmt.Printf("Forest: %d trees\n", forest.GetTrees())
}

func (mp *MapPrinter) VisitCity(city interfaces.CityElement) {
	fmt.Printf("City: %d buildings\n", city.GetBuildings())
}
