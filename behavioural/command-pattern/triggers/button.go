package triggers

import "design-patterns/behavioural/command-pattern/command"

type Button struct{ Command command.Command }

func (b *Button) Press() {
	if b.Command != nil {
		b.Command.Execute()
	}
}
