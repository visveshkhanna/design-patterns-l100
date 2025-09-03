package delivery

import "fmt"

type Webhook struct{}

func (w Webhook) Prepare() {
	fmt.Println("resolving endpoint & payload")
}

func (w Webhook) Dispatch() {
	fmt.Println("POST /webhook -> 200 OK")
}

func (w Webhook) Report() {
	fmt.Println("metrics: delivered")
}
