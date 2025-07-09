package main

import "fmt"

type Folder struct {
	Children []Inode
	Name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.Children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}
