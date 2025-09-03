package elements

import "design-patterns/behavioural/visitor/interfaces"

type Forest struct {
	trees int
}

func NewForest(trees int) *Forest {
	return &Forest{trees: trees}
}

func (f *Forest) Accept(visitor interfaces.Visitor) {
	visitor.VisitForest(f)
}

func (f *Forest) GetTrees() int {
	return f.trees
}
