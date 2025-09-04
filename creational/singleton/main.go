package main

import (
	"fmt"

	"design-patterns/creational/singleton/singleton"
)

func main() {

	pen1 := singleton.GetSignaturePen()
	pen2 := singleton.GetSignaturePen()

	fmt.Println("Pen1 == Pen2:", pen1 == pen2)

	artwork1 := pen1.Sign("Mona Lisa")
	artwork2 := pen2.Sign("Starry Night")

	fmt.Println(artwork1)
	fmt.Println(artwork2)
}
