package main

// PaymentProcessor subsystem handles charging customers.
type PaymentProcessor struct{}

func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{}
}

func (p *PaymentProcessor) Charge(userID string, amount int) bool {
	// In a real system, integrate with a payment gateway.
	// Here we pretend the charge always succeeds if amount > 0.
	return amount > 0 && userID != ""
}
