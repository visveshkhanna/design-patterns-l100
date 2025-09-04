package products

type BaseGun struct {
	name  string
	power int
}

func (g *BaseGun) SetName(name string) {
	g.name = name
}

func (g *BaseGun) GetName() string {
	return g.name
}

func (g *BaseGun) SetPower(power int) {
	g.power = power
}

func (g *BaseGun) GetPower() int {
	return g.power
}
