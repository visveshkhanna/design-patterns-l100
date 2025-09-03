package observers

import "fmt"

type SMSAlert struct {
	phone     string
	threshold float64
}

func NewSMSAlert(phone string, threshold float64) *SMSAlert {
	return &SMSAlert{
		phone:     phone,
		threshold: threshold,
	}
}

func (s *SMSAlert) Update(spend float64) {
	if spend >= s.threshold {
		fmt.Printf("📱 SMS Alert to %s: Budget threshold reached! Current spending: $%.2f\n",
			s.phone, spend)
	}
}
