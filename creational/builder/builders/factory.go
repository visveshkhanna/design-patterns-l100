package builders

import "design-patterns/creational/builder/interfaces"

func GetBuilder(builderType string) interfaces.IBuilder {
	switch builderType {
	case "normal":
		return NewNormalBuilder()
	case "igloo":
		return NewIglooBuilder()
	default:
		return nil
	}
}
