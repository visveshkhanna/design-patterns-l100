package main

// And returns true if both operands are true.
type And struct{ left, right Expr }

func (a *And) Eval(c *Ctx) bool { return a.left.Eval(c) && a.right.Eval(c) }
