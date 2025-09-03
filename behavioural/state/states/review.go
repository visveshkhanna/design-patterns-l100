package states

import (
	"design-patterns/behavioural/state/interfaces"
	"fmt"
)

type Review struct{}

func (Review) Submit(ctx interfaces.Context) {
	fmt.Println("Document is already under review")
}

func (Review) Approve(ctx interfaces.Context) {
	fmt.Println("Document approved and published")
	ctx.SetState(Published{})
}

func (Review) Reject(ctx interfaces.Context) {
	fmt.Println("Document rejected, back to draft")
	ctx.SetState(Draft{})
}
