package subsystems

type PaymentProcessor struct{}

func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{}
}

func (p *PaymentProcessor) Charge(userID string, amount int) bool {

	return amount > 0 && userID != ""
}
