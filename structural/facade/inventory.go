package main

// Inventory subsystem manages stock and pricing.
type Inventory struct {
	stock map[string]int
	price map[string]int
}

func NewInventory() *Inventory {
	return &Inventory{
		stock: map[string]int{
			"phone":   5,
			"case":    20,
			"charger": 10,
		},
		price: map[string]int{
			"phone":   699,
			"case":    29,
			"charger": 39,
		},
	}
}

func (i *Inventory) CheckAvailability(id string, qty int) bool {
	available, ok := i.stock[id]
	if !ok {
		return false
	}
	return available >= qty
}

func (i *Inventory) Reserve(id string, qty int) bool {
	if !i.CheckAvailability(id, qty) {
		return false
	}
	i.stock[id] -= qty
	return true
}

func (i *Inventory) CalculateTotal(items map[string]int) int {
	total := 0
	for id, qty := range items {
		p := i.price[id]
		total += p * qty
	}
	return total
}
