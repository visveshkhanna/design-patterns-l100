package main

import "fmt"

type Observer interface{ update(spend float64) }

type Subject interface {
	register(o Observer)
	unregister(o Observer)
	notifyAll()
}

type Budget struct {
	observers []Observer
	spend     float64
	limit     float64
}

func (b *Budget) addSpend(amt float64) { b.spend += amt; b.notifyAll() }

func (b *Budget) register(o Observer) { b.observers = append(b.observers, o) }

func (b *Budget) unregister(o Observer) {
	for i, ob := range b.observers {
		if ob == o {
			b.observers = append(b.observers[:i], b.observers[i+1:]...)
			break
		}
	}
}

func (b *Budget) notifyAll() {
	for _, o := range b.observers {
		o.update(b.spend)
	}
}

type SMSAlert struct {
	phone     string
	threshold float64
}

func (s *SMSAlert) update(spend float64) {
	if spend >= s.threshold {
		fmt.Println("SMS to", s.phone, "budget reached:", spend)
	}
}
