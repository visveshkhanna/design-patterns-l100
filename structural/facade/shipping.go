package main

import (
	"fmt"
	"time"
)

// ShippingService subsystem handles shipping and returns a tracking number.
type ShippingService struct{}

func NewShippingService() *ShippingService {
	return &ShippingService{}
}

func (s *ShippingService) Ship(userID string, items map[string]int) string {
	// Return a pseudo tracking number.
	return fmt.Sprintf("TRK-%s-%d", userID, time.Now().Unix())
}
