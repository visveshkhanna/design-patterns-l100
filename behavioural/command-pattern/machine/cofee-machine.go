package machine

import "fmt"

type CoffeeMachine struct{ waterMl, beansGr int }

func (c *CoffeeMachine) AddWater(ml int) {
	c.waterMl += ml
	fmt.Println("water:", c.waterMl, "ml")
}

func (c *CoffeeMachine) AddBeans(gr int) {
	c.beansGr += gr
	fmt.Println("beans:", c.beansGr, "g")
}

func (c *CoffeeMachine) BrewEspresso() {
	if c.waterMl < 30 || c.beansGr < 8 {
		fmt.Println("not enough resources")
		return
	}
	c.waterMl -= 30
	c.beansGr -= 8
	fmt.Println("espresso ready")
}

func (c *CoffeeMachine) Clean() {
	fmt.Println("clean cycle done")
}
