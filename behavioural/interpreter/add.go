package main

type Add struct{ left, right Expression }

func (a *Add) Interpret() int { return a.left.Interpret() + a.right.Interpret() }