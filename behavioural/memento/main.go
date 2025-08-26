package main

func main() {
	editor := &Editor{}
	history := &History{}

	editor.typeText("Hello ")
	history.push(editor.save())

	editor.typeText("World")
	history.push(editor.save())

	editor.typeText("!!!")
	editor.print()

	if m, ok := history.pop(); ok {
		editor.restore(m)
	}
	editor.print()

	if m, ok := history.pop(); ok {
		editor.restore(m)
	}
	editor.print()
}
