package subjects

import "design-patterns/behavioural/observer/interfaces"

type Budget struct {
	observers []interfaces.IObserver
	spend     float64
	limit     float64
}

func NewBudget(limit float64) *Budget {
	return &Budget{
		observers: make([]interfaces.IObserver, 0),
		spend:     0,
		limit:     limit,
	}
}

func (b *Budget) AddSpend(amount float64) {
	b.spend += amount
	b.NotifyAll()
}

func (b *Budget) GetSpend() float64 {
	return b.spend
}

func (b *Budget) GetLimit() float64 {
	return b.limit
}

func (b *Budget) Register(observer interfaces.IObserver) {
	b.observers = append(b.observers, observer)
}

func (b *Budget) Unregister(observer interfaces.IObserver) {
	for i, obs := range b.observers {
		if obs == observer {
			b.observers = append(b.observers[:i], b.observers[i+1:]...)
			break
		}
	}
}

func (b *Budget) NotifyAll() {
	for _, observer := range b.observers {
		observer.Update(b.spend)
	}
}
