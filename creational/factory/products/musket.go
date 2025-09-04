package products

import "design-patterns/creational/factory/interfaces"

type Musket struct {
	BaseGun
}

func NewMusket() interfaces.IGun {
	return &Musket{
		BaseGun: BaseGun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
