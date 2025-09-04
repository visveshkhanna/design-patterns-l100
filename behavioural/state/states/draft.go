package states

import (
	"design-patterns/behavioural/state/interfaces"
	"fmt"
)

type Draft struct{}

func (Draft) Submit(ctx interfaces.IContext) {
	fmt.Println("Document submitted for review")
	ctx.SetState(Review{})
}

func (Draft) Approve(ctx interfaces.IContext) {
	fmt.Println("Cannot approve a draft document")
}

func (Draft) Reject(ctx interfaces.IContext) {
	fmt.Println("Cannot reject a draft document")
}
