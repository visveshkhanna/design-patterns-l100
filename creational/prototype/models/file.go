package models

import (
	"fmt"

	"design-patterns/creational/prototype/interfaces"
)

type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) Clone() interfaces.IPrototype {
	return &File{Name: f.Name + "_clone"}
}
