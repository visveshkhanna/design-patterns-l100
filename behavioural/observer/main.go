package main

import (
	"fmt"

	"design-patterns/behavioural/observer/observers"
	"design-patterns/behavioural/observer/subjects"
)

func main() {
	fmt.Println("=== Observer Pattern Demo: Budget Alerts ===")

	budget := subjects.NewBudget(150)

	smsAlert := observers.NewSMSAlert("+1-555-0101", 50)
	emailAlert := observers.NewEmailAlert("user@example.com", 100)
	urgentSMS := observers.NewSMSAlert("+1-555-0199", 120)

	budget.Register(smsAlert)
	budget.Register(emailAlert)
	budget.Register(urgentSMS)

	fmt.Printf("Budget created with limit: $%.2f\n\n", budget.GetLimit())

	fmt.Println("Adding $30 expense...")
	budget.AddSpend(30)

	fmt.Println("\nAdding $25 expense...")
	budget.AddSpend(25)

	fmt.Println("\nAdding $50 expense...")
	budget.AddSpend(50)

	fmt.Println("\nRemoving SMS alert and adding $20 expense...")
	budget.Unregister(smsAlert)
	budget.AddSpend(20)

	fmt.Printf("\nFinal spending: $%.2f / $%.2f\n", budget.GetSpend(), budget.GetLimit())
}
