package main

import "fmt"

// OrderFacade provides a simple API to place orders using multiple subsystems.
type OrderFacade struct {
	inventory *Inventory
	payment   *PaymentProcessor
	shipping  *ShippingService
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		inventory: NewInventory(),
		payment:   NewPaymentProcessor(),
		shipping:  NewShippingService(),
	}
}

// PlaceOrder checks stock, reserves items, charges payment, and ships.
func (f *OrderFacade) PlaceOrder(userID string, items map[string]int) (string, error) {
	// 1) Check availability
	for id, qty := range items {
		if !f.inventory.CheckAvailability(id, qty) {
			return "", fmt.Errorf("item %s is out of stock or insufficient qty", id)
		}
	}

	// 2) Calculate total
	amount := f.inventory.CalculateTotal(items)

	// 3) Reserve items
	for id, qty := range items {
		if ok := f.inventory.Reserve(id, qty); !ok {
			return "", fmt.Errorf("failed to reserve %s", id)
		}
	}

	// 4) Charge payment
	if !f.payment.Charge(userID, amount) {
		return "", fmt.Errorf("payment failed for user %s", userID)
	}

	// 5) Ship
	tracking := f.shipping.Ship(userID, items)
	return tracking, nil
}
