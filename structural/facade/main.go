package main

import (
	"fmt"
)

func main() {
	facade := NewOrderFacade()
	items := map[string]int{
		"phone":   1,
		"case":    2,
		"charger": 1,
	}

	tracking, err := facade.PlaceOrder("user-42", items)
	if err != nil {
		fmt.Println("Order failed:", err)
		return
	}

	fmt.Println("Order placed! Tracking:", tracking)
}
