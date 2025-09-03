package expressions

// Predicates
type IsRole struct{ Role string }
type IsRegion struct{ Region string }
type IsBeta struct{}
type RequestsBelow struct{ Max int }

func (p *IsRole) Eval(c *Ctx) bool        { return c.Role == p.Role }
func (p *IsRegion) Eval(c *Ctx) bool      { return c.Region == p.Region }
func (p *IsBeta) Eval(c *Ctx) bool        { return c.Beta }
func (p *RequestsBelow) Eval(c *Ctx) bool { return c.Requests < p.Max }
