package expressions

// And returns true if both operands are true.
type And struct{ Left, Right Expr }

func (a *And) Eval(c *Ctx) bool { return a.Left.Eval(c) && a.Right.Eval(c) }
