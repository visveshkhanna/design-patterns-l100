package main

import (
	"design-patterns/structural/facade/facade"
	"fmt"
)

func main() {
	orderFacade := facade.NewOrderFacade()
	items := map[string]int{
		"phone":   1,
		"case":    2,
		"charger": 1,
	}

	tracking, err := orderFacade.PlaceOrder("user-42", items)
	if err != nil {
		fmt.Println("Order failed:", err)
		return
	}

	fmt.Println("Order placed! Tracking:", tracking)
}
