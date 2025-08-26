package main

import "fmt"

// CoffeeMachine is the receiver providing concrete operations.
type CoffeeMachine struct{ waterMl, beansGr int }

func (c *CoffeeMachine) addWater(ml int) {
	c.waterMl += ml
	fmt.Println("water:", c.waterMl, "ml")
}

func (c *CoffeeMachine) addBeans(gr int) {
	c.beansGr += gr
	fmt.Println("beans:", c.beansGr, "g")
}

func (c *CoffeeMachine) brewEspresso() {
	if c.waterMl < 30 || c.beansGr < 8 {
		fmt.Println("not enough resources")
		return
	}
	c.waterMl -= 30
	c.beansGr -= 8
	fmt.Println("espresso ready")
}

func (c *CoffeeMachine) clean() {
	fmt.Println("clean cycle done")
}
