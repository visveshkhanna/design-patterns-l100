package main

import (
	"design-patterns/structural/composite/components"
	"fmt"
)

func main() {

	phone := components.NewItem("Phone", 699)
	caseItem := components.NewItem("Case", 29)
	charger := components.NewItem("Charger", 39)

	cables := components.NewBox("Cables")
	cables.Add(components.NewItem("USB-C Cable", 12))
	cables.Add(components.NewItem("Lightning Cable", 15))

	accessories := components.NewBox("Accessories")
	accessories.Add(caseItem)
	accessories.Add(charger)
	accessories.Add(cables)

	order := components.NewBox("Order #1001")
	order.Add(phone)
	order.Add(accessories)

	fmt.Println("Order contents:")
	order.Print("  ")
	fmt.Printf("\nGrand total: $%d\n", order.Price())
}
