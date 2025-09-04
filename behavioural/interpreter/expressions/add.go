package expressions

type Or struct{ Left, Right IExpr }

func (o *Or) Eval(c *Ctx) bool { return o.Left.Eval(c) || o.Right.Eval(c) }
