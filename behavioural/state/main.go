package main

import (
	"design-patterns/behavioural/state/context"
	"design-patterns/behavioural/state/states"
)

func main() {

	d := context.NewDocument("Proposal", states.Draft{})

	d.Submit()
	d.Approve()
	d.Reject()
}
