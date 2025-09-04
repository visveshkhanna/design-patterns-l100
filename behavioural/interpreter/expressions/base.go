package expressions

type IExpr interface {
	Eval(*Ctx) bool
}

type Ctx struct {
	Role     string
	Region   string
	Beta     bool
	Requests int
}
