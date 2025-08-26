package main

// Predicates
type IsRole struct{ role string }
type IsRegion struct{ region string }
type IsBeta struct{}
type RequestsBelow struct{ max int }

func (p *IsRole) Eval(c *Ctx) bool        { return c.Role == p.role }
func (p *IsRegion) Eval(c *Ctx) bool      { return c.Region == p.region }
func (p *IsBeta) Eval(c *Ctx) bool        { return c.Beta }
func (p *RequestsBelow) Eval(c *Ctx) bool { return c.Requests < p.max }
