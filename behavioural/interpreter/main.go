package main

import "fmt"

func main() {
	// (5 + 3) - 2 = 6
	expr := &Subtract{
		left:  &Add{left: &Number{value: 5}, right: &Number{value: 3}},
		right: &Number{value: 2},
	}
	fmt.Println("Result:", expr.Interpret())
}
