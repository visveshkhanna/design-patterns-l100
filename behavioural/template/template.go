package main

import "fmt"

type Game interface {
	initialize()
	startPlay()
	endPlay()
}

type GameTemplate struct{ game Game }

func (g *GameTemplate) Play() {
	g.game.initialize()
	g.game.startPlay()
	g.game.endPlay()
}

type Football struct{}

func (Football) initialize() { fmt.Println("Football Game Initialized!") }
func (Football) startPlay()  { fmt.Println("Football Game Started. Enjoy the game!") }
func (Football) endPlay()    { fmt.Println("Football Game Finished!") }
