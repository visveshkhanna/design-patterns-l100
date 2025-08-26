package main

import "fmt"

type Handler interface {
	SetNext(next Handler)
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(next Handler) {
	b.next = next
}

func (b *BaseHandler) callNext(request string) {
	if b.next != nil {
		b.next.Handle(request)
	} else {
		fmt.Println("End of chain.")
	}
}
