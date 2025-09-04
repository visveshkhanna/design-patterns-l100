package expressions

type And struct{ Left, Right IExpr }

func (a *And) Eval(c *Ctx) bool { return a.Left.Eval(c) && a.Right.Eval(c) }
