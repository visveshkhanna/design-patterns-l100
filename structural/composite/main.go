package main

import (
	"fmt"
)

func main() {
	phone := NewItem("Phone", 699)
	caseItem := NewItem("Case", 29)
	charger := NewItem("Charger", 39)

	cables := NewBox("Cables")
	cables.Add(NewItem("USB-C Cable", 12))
	cables.Add(NewItem("Lightning Cable", 15))

	accessories := NewBox("Accessories")
	accessories.Add(caseItem)
	accessories.Add(charger)
	accessories.Add(cables)

	order := NewBox("Order #1001")
	order.Add(phone)
	order.Add(accessories)

	fmt.Println("Order contents:")
	order.Print("  ")
	fmt.Printf("\nGrand total: $%d\n", order.Price())
}
