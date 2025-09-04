package models

import (
	"fmt"

	"design-patterns/creational/prototype/interfaces"
)

type Folder struct {
	Children []interfaces.IPrototype
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, child := range f.Children {
		child.Print(indentation + "  ")
	}
}

func (f *Folder) Clone() interfaces.IPrototype {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []interfaces.IPrototype
	for _, child := range f.Children {
		copy := child.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}
