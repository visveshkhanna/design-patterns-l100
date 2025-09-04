package interfaces

type IState interface {
	Submit(ctx IContext)
	Approve(ctx IContext)
	Reject(ctx IContext)
}

type IContext interface {
	SetState(s IState)
}
