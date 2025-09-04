package factory

import (
	"design-patterns/creational/factory/interfaces"
	"design-patterns/creational/factory/products"
	"fmt"
)

func GetGun(gunType string) (interfaces.IGun, error) {
	switch gunType {
	case "ak47":
		return products.NewAK47(), nil
	case "musket":
		return products.NewMusket(), nil
	default:
		return nil, fmt.Errorf("unknown gun type: %s", gunType)
	}
}
