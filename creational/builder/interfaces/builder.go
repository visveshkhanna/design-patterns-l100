package interfaces

import "design-patterns/creational/builder/models"

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() models.House
}
