package director

import (
	"design-patterns/creational/builder/interfaces"
	"design-patterns/creational/builder/models"
)

type Director struct {
	builder interfaces.IBuilder
}

func NewDirector(b interfaces.IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b interfaces.IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() models.House {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}
