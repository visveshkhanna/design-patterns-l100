package products

import "design-patterns/creational/factory/interfaces"

type AK47 struct {
	BaseGun
}

func NewAK47() interfaces.IGun {
	return &AK47{
		BaseGun: BaseGun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
