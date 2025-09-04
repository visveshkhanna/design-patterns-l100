package computers

import (
	"design-patterns/structural/bridge/interfaces"
	"fmt"
)

type Mac struct {
	printer interfaces.IPrinter
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p interfaces.IPrinter) {
	m.printer = p
}
