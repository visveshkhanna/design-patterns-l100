/*

Idea: The artist has only one signature pen used to sign all artworks.

Concept: Ensures a class has only one instance and provides a global access point.

Example: A signature pen used by the artist must be single and reused.

*/

package main

import (
	"fmt"
	"sync"
)

type SignaturePen struct {
	Color string
}

var instance *SignaturePen
var once sync.Once

func GetSignaturePen() *SignaturePen {
	once.Do(func() {
		instance = &SignaturePen{Color: "Gold"}
		fmt.Println("Signature Pen created")
	})
	return instance
}

func main() {
	pen1 := GetSignaturePen()
	pen2 := GetSignaturePen()

	fmt.Println("Pen1 == Pen2:", pen1 == pen2)
}
