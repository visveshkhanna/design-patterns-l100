package command

import "design-patterns/behavioural/command-pattern/machine"

type CleanCommand struct{ Machine *machine.CoffeeMachine }

func (c *CleanCommand) Execute() { c.Machine.Clean() }
