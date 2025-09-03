package main

import (
	"design-patterns/behavioural/template/delivery"
	"design-patterns/behavioural/template/template"
)

func main() {
	webhook := delivery.Webhook{}
	t := &template.DeliveryTemplate{Delivery: webhook}
	t.Send()
}
