package main

func main() {
	p := &Product{price: 10}
	s1 := &EmailSubscriber{email: "a@example.com"}
	s2 := &EmailSubscriber{email: "b@example.com"}
	p.register(s1)
	p.register(s2)
	p.setPrice(12.5)
	p.unregister(s1)
	p.setPrice(9.9)
}
