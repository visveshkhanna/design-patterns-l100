package main

import (
	"fmt"

	"design-patterns/creational/prototype/interfaces"
	"design-patterns/creational/prototype/models"
)

func main() {

	file1 := &models.File{Name: "File1"}
	file2 := &models.File{Name: "File2"}
	file3 := &models.File{Name: "File3"}

	folder1 := &models.Folder{
		Children: []interfaces.IPrototype{file1},
		Name:     "Folder1",
	}

	folder2 := &models.Folder{
		Children: []interfaces.IPrototype{folder1, file2, file3},
		Name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2:")
	folder2.Print("")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for cloned Folder:")
	cloneFolder.Print("")
}
