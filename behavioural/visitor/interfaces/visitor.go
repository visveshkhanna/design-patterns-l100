package interfaces

type IVisitor interface {
	VisitForest(forest IForestElement)
	VisitCity(city ICityElement)
}
