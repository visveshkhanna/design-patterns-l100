package main

import "fmt"

type Item struct {
	name  string
	price int
}

func NewItem(name string, price int) *Item {
	return &Item{name: name, price: price}
}

func (i *Item) Price() int {
	return i.price
}

func (i *Item) Print(indent string) {
	fmt.Printf("%s- %s: $%d\n", indent, i.name, i.price)
}
