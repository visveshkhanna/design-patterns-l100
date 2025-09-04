package singleton

import (
	"fmt"
	"sync"

	"design-patterns/creational/singleton/models"
)

var (
	instance *models.SignaturePen
	once     sync.Once
)

func GetSignaturePen() *models.SignaturePen {
	once.Do(func() {
		instance = &models.SignaturePen{Color: "Gold"}
		fmt.Println("Signature Pen created")
	})
	return instance
}
