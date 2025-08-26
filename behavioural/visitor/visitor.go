package main

import "fmt"

type Visitor interface {
	visitCircle(*Circle)
	visitSquare(*Square)
}

type Element interface{ accept(Visitor) }

type Circle struct{ radius int }

type Square struct{ side int }

func (c *Circle) accept(v Visitor) { v.visitCircle(c) }

func (s *Square) accept(v Visitor) { v.visitSquare(s) }

type AreaCalculator struct{}

func (AreaCalculator) visitCircle(c *Circle) {
	fmt.Println("Area of circle:", 3.14*float64(c.radius*c.radius))
}

func (AreaCalculator) visitSquare(s *Square) { fmt.Println("Area of square:", s.side*s.side) }
