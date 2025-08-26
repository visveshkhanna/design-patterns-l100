package main

type Subtract struct{ left, right Expression }

func (s *Subtract) Interpret() int { return s.left.Interpret() - s.right.Interpret() }
