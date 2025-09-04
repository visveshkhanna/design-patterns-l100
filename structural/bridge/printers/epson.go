package printers

import "fmt"

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by epson")
}
