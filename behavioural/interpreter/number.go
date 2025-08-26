package main

type Number struct{ value int }

func (n *Number) Interpret() int { return n.value }