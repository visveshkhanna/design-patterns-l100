package computers

import (
	"design-patterns/structural/bridge/interfaces"
	"fmt"
)

type Windows struct {
	printer interfaces.IPrinter
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p interfaces.IPrinter) {
	w.printer = p
}
