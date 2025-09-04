package handlers

import "design-patterns/behavioural/chainofres/models"

type IHandler interface {
	SetNext(next IHandler)
	Handle(msg *models.Message)
}

type BaseHandler struct {
	next IHandler
}

func (b *BaseHandler) SetNext(next IHandler) {
	b.next = next
}

func (b *BaseHandler) callNext(msg *models.Message) {
	if b.next != nil {
		b.next.Handle(msg)
	}
}
