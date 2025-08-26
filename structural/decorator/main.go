package main

import (
	"fmt"
)

func main() {
	base := NewEspresso()
	withMilk := WithMilk(base)
	withMilkAndSugar := WithSugar(withMilk)

	fmt.Printf("Order 1: %s => %d\n", base.Description(), base.Cost())
	fmt.Printf("Order 2: %s => %d\n", withMilk.Description(), withMilk.Cost())
	fmt.Printf("Order 3: %s => %d\n", withMilkAndSugar.Description(), withMilkAndSugar.Cost())
}
