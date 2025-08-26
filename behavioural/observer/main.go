package main

func main() {
	b := &Budget{limit: 100}
	alert1 := &SMSAlert{phone: "+111", threshold: 50}
	alert2 := &SMSAlert{phone: "+222", threshold: 90}
	b.register(alert1)
	b.register(alert2)
	b.addSpend(40)
	b.addSpend(20)
	b.unregister(alert1)
	b.addSpend(50)
}
