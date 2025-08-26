package main

type Component interface {
	Price() int
	Print(indent string)
}
