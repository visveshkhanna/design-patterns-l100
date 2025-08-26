package main

import "fmt"

type Observer interface{ update(price float64) }

type Subject interface {
	register(o Observer)
	unregister(o Observer)
	notifyAll()
}

type Product struct {
	observers []Observer
	price     float64
}

func (p *Product) setPrice(price float64) { p.price = price; p.notifyAll() }

func (p *Product) register(o Observer) { p.observers = append(p.observers, o) }

func (p *Product) unregister(o Observer) {
	for i, ob := range p.observers {
		if ob == o {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
			break
		}
	}
}

func (p *Product) notifyAll() {
	for _, o := range p.observers {
		o.update(p.price)
	}
}

type EmailSubscriber struct{ email string }

func (e *EmailSubscriber) update(price float64) {
	fmt.Println("Email to", e.email, "new price:", price)
}
