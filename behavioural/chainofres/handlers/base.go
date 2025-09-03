package handlers

import "design-patterns/behavioural/chainofres/models"

type Handler interface {
	SetNext(next Handler)
	Handle(msg *models.Message)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(next Handler) {
	b.next = next
}

func (b *BaseHandler) callNext(msg *models.Message) {
	if b.next != nil {
		b.next.Handle(msg)
	}
}
