package command

import "design-patterns/behavioural/command-pattern/machine"

type AddWaterCommand struct {
	Machine *machine.CoffeeMachine
	Ml      int
}

func (c *AddWaterCommand) Execute() { c.Machine.AddWater(c.Ml) }
