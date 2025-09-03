package triggers

import "design-patterns/behavioural/command-pattern/command"

type Remote struct{ queue []command.Command }

func (r *Remote) Add(c command.Command) { r.queue = append(r.queue, c) }

func (r *Remote) RunAll() {
	for _, c := range r.queue {
		c.Execute()
	}
	r.queue = nil
}