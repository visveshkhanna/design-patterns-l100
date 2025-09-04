package main

import (
	"fmt"

	"design-patterns/structural/decorator/components"
	"design-patterns/structural/decorator/decorators"
)

func main() {
	base := components.NewEspresso()
	withMilk := decorators.WithMilk(base)
	withMilkAndSugar := decorators.WithSugar(withMilk)

	fmt.Printf("Order 1: %s => %d\n", base.Description(), base.Cost())
	fmt.Printf("Order 2: %s => %d\n", withMilk.Description(), withMilk.Cost())
	fmt.Printf("Order 3: %s => %d\n", withMilkAndSugar.Description(), withMilkAndSugar.Cost())
}
