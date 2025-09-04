package main

import (
	"fmt"

	"design-patterns/creational/builder/builders"
	"design-patterns/creational/builder/director"
)

func main() {
	normalBuilder := builders.GetBuilder("normal")
	iglooBuilder := builders.GetBuilder("igloo")

	houseDirector := director.NewDirector(normalBuilder)
	normalHouse := houseDirector.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	houseDirector.SetBuilder(iglooBuilder)
	iglooHouse := houseDirector.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)
}
