package interfaces

type IElement interface {
	Accept(visitor IVisitor)
}

type IForestElement interface {
	IElement
	GetTrees() int
}

type ICityElement interface {
	IElement
	GetBuildings() int
}
