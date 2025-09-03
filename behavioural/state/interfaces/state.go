package interfaces

type State interface {
	Submit(ctx Context)
	Approve(ctx Context)
	Reject(ctx Context)
}

type Context interface {
	SetState(s State)
}
