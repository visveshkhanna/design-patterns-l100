package main

import "fmt"

type Delivery interface {
	prepare()
	dispatch()
	report()
}

type DeliveryTemplate struct{ d Delivery }

func (t *DeliveryTemplate) Send() {
	t.d.prepare()
	t.d.dispatch()
	t.d.report()
}

type Webhook struct{}

func (Webhook) prepare()  { fmt.Println("resolving endpoint & payload") }
func (Webhook) dispatch() { fmt.Println("POST /webhook -> 200 OK") }
func (Webhook) report()   { fmt.Println("metrics: delivered") }
