package main

type Button struct{ command Command }

func (b *Button) press() {
	if b.command != nil {
		b.command.execute()
	}
}
