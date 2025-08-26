package main

func main() {
	machine := &CoffeeMachine{}

	addWater := &AddWaterCommand{machine: machine, ml: 120}
	addBeans := &AddBeansCommand{machine: machine, gr: 16}
	brew := &BrewEspressoCommand{machine: machine}
	clean := &CleanCommand{machine: machine}

	// Buttons trigger single actions
	waterBtn := &Button{command: addWater}
	beansBtn := &Button{command: addBeans}
	brewBtn := &Button{command: brew}
	cleanBtn := &Button{command: clean}

	waterBtn.press()
	beansBtn.press()
	brewBtn.press()
	cleanBtn.press()

	// Remote queues a recipe
	remote := &Remote{}
	remote.add(&AddWaterCommand{machine: machine, ml: 60})
	remote.add(&AddBeansCommand{machine: machine, gr: 8})
	remote.add(&BrewEspressoCommand{machine: machine})
	remote.runAll()
}
