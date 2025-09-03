package observers

import "fmt"

type EmailAlert struct {
	email     string
	threshold float64
}

func NewEmailAlert(email string, threshold float64) *EmailAlert {
	return &EmailAlert{
		email:     email,
		threshold: threshold,
	}
}

func (e *EmailAlert) Update(spend float64) {
	if spend >= e.threshold {
		fmt.Printf("📧 Email Alert to %s: Budget alert! Current spending: $%.2f\n",
			e.email, spend)
	}
}
