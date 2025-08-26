package main

func main() {
	t := &DeliveryTemplate{d: Webhook{}}
	t.Send()
}
