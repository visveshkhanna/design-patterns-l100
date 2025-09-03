package template

import "design-patterns/behavioural/template/interfaces"

type DeliveryTemplate struct {
	Delivery interfaces.Delivery
}

func (t *DeliveryTemplate) Send() {
	t.Delivery.Prepare()
	t.Delivery.Dispatch()
	t.Delivery.Report()
}
