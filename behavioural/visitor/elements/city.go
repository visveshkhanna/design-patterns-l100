package elements

import "design-patterns/behavioural/visitor/interfaces"

type City struct {
	buildings int
}

func NewCity(buildings int) *City {
	return &City{buildings: buildings}
}

func (c *City) Accept(visitor interfaces.IVisitor) {
	visitor.VisitCity(c)
}

func (c *City) GetBuildings() int {
	return c.buildings
}
