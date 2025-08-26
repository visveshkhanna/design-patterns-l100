package main

import "fmt"

type Visitor interface {
	visitForest(*Forest)
	visitCity(*City)
}

type Element interface{ accept(Visitor) }

type Forest struct{ trees int }

type City struct{ buildings int }

func (f *Forest) accept(v Visitor) { v.visitForest(f) }

func (c *City) accept(v Visitor) { v.visitCity(c) }

type MapPrinter struct{}

func (MapPrinter) visitForest(f *Forest) { fmt.Println("Forest:", f.trees, "trees") }

func (MapPrinter) visitCity(c *City) { fmt.Println("City:", c.buildings, "buildings") }
