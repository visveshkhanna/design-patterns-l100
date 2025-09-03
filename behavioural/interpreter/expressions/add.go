package expressions

// Or returns true if either operand is true.
type Or struct{ Left, Right Expr }

func (o *Or) Eval(c *Ctx) bool { return o.Left.Eval(c) || o.Right.Eval(c) }
