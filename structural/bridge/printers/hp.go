package printers

import "fmt"

type HP struct{}

func (h *HP) PrintFile() {
	fmt.Println("Printing by hp")
}
