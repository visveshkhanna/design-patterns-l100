package main

import "fmt"

type State interface{ pressPlay(ctx *Player) }

type Player struct{ state State }

func (p *Player) setState(s State) { p.state = s }
func (p *Player) pressPlay()       { p.state.pressPlay(p) }

type PlayingState struct{}

type PausedState struct{}

func (PlayingState) pressPlay(ctx *Player) { fmt.Println("Pausing"); ctx.setState(PausedState{}) }

func (PausedState) pressPlay(ctx *Player) { fmt.Println("Playing"); ctx.setState(PlayingState{}) }
