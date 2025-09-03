package command

import "design-patterns/behavioural/command-pattern/machine"

type BrewEspressoCommand struct{ Machine *machine.CoffeeMachine }

func (c *BrewEspressoCommand) Execute() { c.Machine.BrewEspresso() }
