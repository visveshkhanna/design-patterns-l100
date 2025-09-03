package main

import (
	"design-patterns/behavioural/command-pattern/command"
	"design-patterns/behavioural/command-pattern/machine"
	"design-patterns/behavioural/command-pattern/triggers"
)

func main() {
	machine := &machine.CoffeeMachine{}

	addWater := &command.AddWaterCommand{Machine: machine, Ml: 120}
	addBeans := &command.AddBeansCommand{Machine: machine, Gr: 16}
	brew := &command.BrewEspressoCommand{Machine: machine}
	clean := &command.CleanCommand{Machine: machine}

	// Buttons trigger single actions
	waterBtn := &triggers.Button{Command: addWater}
	beansBtn := &triggers.Button{Command: addBeans}
	brewBtn := &triggers.Button{Command: brew}
	cleanBtn := &triggers.Button{Command: clean}

	waterBtn.Press()
	beansBtn.Press()
	brewBtn.Press()
	cleanBtn.Press()

	// Remote queues a recipe
	remote := &triggers.Remote{}
	remote.Add(&command.AddWaterCommand{Machine: machine, Ml: 60})
	remote.Add(&command.AddBeansCommand{Machine: machine, Gr: 8})
	remote.Add(&command.BrewEspressoCommand{Machine: machine})
	remote.RunAll()
}
