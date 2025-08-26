package main

import "fmt"

func main() {
	c := &IntCollection{items: []int{1, 2, 3, 4}}
	it := c.createIterator()
	for it.hasNext() {
		fmt.Println(it.next())
	}
}
