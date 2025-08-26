package main

import "fmt"

func main() {

	hpPrinter := &HP{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	windowsComputer := &Windows{}
	
	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()
	fmt.Println()

	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
	fmt.Println()
	
}