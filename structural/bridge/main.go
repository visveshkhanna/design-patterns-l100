package main

import (
	"design-patterns/structural/bridge/computers"
	"design-patterns/structural/bridge/printers"
	"fmt"
)

func main() {
	hpPrinter := &printers.HP{}
	epsonPrinter := &printers.Epson{}

	macComputer := &computers.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	windowsComputer := &computers.Windows{}

	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()
	fmt.Println()

	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
	fmt.Println()
}
