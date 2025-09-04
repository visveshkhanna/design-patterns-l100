package states

import (
	"design-patterns/behavioural/state/interfaces"
	"fmt"
)

type Published struct{}

func (Published) Submit(ctx interfaces.IContext) {
	fmt.Println("Document is already published")
}

func (Published) Approve(ctx interfaces.IContext) {
	fmt.Println("Document is already published")
}

func (Published) Reject(ctx interfaces.IContext) {
	fmt.Println("Cannot reject a published document")
}
