package states

import (
	"design-patterns/behavioural/state/interfaces"
	"fmt"
)

type Draft struct{}

func (Draft) Submit(ctx interfaces.Context) {
	fmt.Println("Document submitted for review")
	ctx.SetState(Review{})
}

func (Draft) Approve(ctx interfaces.Context) {
	fmt.Println("Cannot approve a draft document")
}

func (Draft) Reject(ctx interfaces.Context) {
	fmt.Println("Cannot reject a draft document")
}
