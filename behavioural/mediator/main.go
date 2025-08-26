package main

func main() {
	dialog := &Dialog{}
	button := &Button{mediator: dialog}
	input := &Input{mediator: dialog}
	dialog.button = button
	dialog.input = input

	button.click()
	input.change()
}
