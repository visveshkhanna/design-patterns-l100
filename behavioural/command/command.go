package main

// Command defines the executable action
type Command interface {
	execute()
}

type AddWaterCommand struct {
	machine *CoffeeMachine
	ml      int
}

type AddBeansCommand struct {
	machine *CoffeeMachine
	gr      int
}

type BrewEspressoCommand struct{ machine *CoffeeMachine }

type CleanCommand struct{ machine *CoffeeMachine }

func (c *AddWaterCommand) execute()     { c.machine.addWater(c.ml) }
func (c *AddBeansCommand) execute()     { c.machine.addBeans(c.gr) }
func (c *BrewEspressoCommand) execute() { c.machine.brewEspresso() }
func (c *CleanCommand) execute()        { c.machine.clean() }
