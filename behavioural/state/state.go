package main

import "fmt"

type State interface {
	submit(ctx *Document)
	approve(ctx *Document)
	reject(ctx *Document)
}

type Document struct {
	state State
	title string
}

func (d *Document) setState(s State) { d.state = s }
func (d *Document) submit()          { d.state.submit(d) }
func (d *Document) approve()         { d.state.approve(d) }
func (d *Document) reject()          { d.state.reject(d) }

type Draft struct{}

type Review struct{}

type Published struct{}

func (Draft) submit(ctx *Document)  { fmt.Println("-> Review"); ctx.setState(Review{}) }
func (Draft) approve(ctx *Document) { fmt.Println("noop") }
func (Draft) reject(ctx *Document)  { fmt.Println("noop") }

func (Review) submit(ctx *Document)  { fmt.Println("noop") }
func (Review) approve(ctx *Document) { fmt.Println("-> Published"); ctx.setState(Published{}) }
func (Review) reject(ctx *Document)  { fmt.Println("-> Draft"); ctx.setState(Draft{}) }

func (Published) submit(ctx *Document)  { fmt.Println("noop") }
func (Published) approve(ctx *Document) { fmt.Println("noop") }
func (Published) reject(ctx *Document)  { fmt.Println("noop") }
