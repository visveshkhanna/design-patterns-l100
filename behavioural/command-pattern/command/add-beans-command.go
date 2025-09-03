package command

import "design-patterns/behavioural/command-pattern/machine"

type AddBeansCommand struct {
	Machine *machine.CoffeeMachine
	Gr      int
}

func (c *AddBeansCommand) Execute() { c.Machine.AddBeans(c.Gr) }
