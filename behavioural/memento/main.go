package main

func main() {
	field := &TextField{}
	history := &History{}

	field.typeText("Hello ")
	history.push(field.save())

	field.typeText("World")
	field.moveCursor(3)
	history.push(field.save())

	field.typeText("!!!")
	field.print()

	if m, ok := history.pop(); ok {
		field.restore(m)
	}
	field.print()

	if m, ok := history.pop(); ok {
		field.restore(m)
	}
	field.print()
}
