package main

// Or returns true if either operand is true.
type Or struct{ left, right Expr }

func (o *Or) Eval(c *Ctx) bool { return o.left.Eval(c) || o.right.Eval(c) }
