package main

type Handler interface {
	SetNext(next Handler)
	Handle(msg *Message)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(next Handler) {
	b.next = next
}

func (b *BaseHandler) callNext(msg *Message) {
	if b.next != nil {
		b.next.Handle(msg)
	}
}

type Message struct {
	Text     string
	Flags    []string
	Rejected bool
	Reason   string
}
