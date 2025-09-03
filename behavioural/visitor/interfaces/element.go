package interfaces

type Element interface {
	Accept(visitor Visitor)
}

type ForestElement interface {
	Element
	GetTrees() int
}

type CityElement interface {
	Element
	GetBuildings() int
}
