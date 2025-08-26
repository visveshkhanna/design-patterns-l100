package main

type Button struct{ command Command }

func (b *Button) press() {
	if b.command != nil {
		b.command.execute()
	}
}

// Remote can queue multiple commands to run in sequence.
type Remote struct{ queue []Command }

func (r *Remote) add(c Command) { r.queue = append(r.queue, c) }

func (r *Remote) runAll() {
	for _, c := range r.queue {
		c.execute()
	}
	r.queue = nil
}
