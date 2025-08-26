package main

func main() {
	zones := []Element{&Forest{trees: 120}, &City{buildings: 45}}
	printer := MapPrinter{}
	for _, z := range zones {
		z.accept(printer)
	}
}
