package interfaces

type Visitor interface {
	VisitForest(forest ForestElement)
	VisitCity(city CityElement)
}
