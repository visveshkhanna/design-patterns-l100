package facade

import (
	"design-patterns/structural/facade/subsystems"
	"fmt"
)

type OrderFacade struct {
	inventory *subsystems.Inventory
	payment   *subsystems.PaymentProcessor
	shipping  *subsystems.ShippingService
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		inventory: subsystems.NewInventory(),
		payment:   subsystems.NewPaymentProcessor(),
		shipping:  subsystems.NewShippingService(),
	}
}

func (f *OrderFacade) PlaceOrder(userID string, items map[string]int) (string, error) {

	for id, qty := range items {
		if !f.inventory.CheckAvailability(id, qty) {
			return "", fmt.Errorf("item %s is out of stock or insufficient qty", id)
		}
	}

	amount := f.inventory.CalculateTotal(items)

	for id, qty := range items {
		if ok := f.inventory.Reserve(id, qty); !ok {
			return "", fmt.Errorf("failed to reserve %s", id)
		}
	}

	if !f.payment.Charge(userID, amount) {
		return "", fmt.Errorf("payment failed for user %s", userID)
	}

	tracking := f.shipping.Ship(userID, items)
	return tracking, nil
}
