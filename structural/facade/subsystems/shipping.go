package subsystems

import (
	"fmt"
	"time"
)

type ShippingService struct{}

func NewShippingService() *ShippingService {
	return &ShippingService{}
}

func (s *ShippingService) Ship(userID string, items map[string]int) string {

	return fmt.Sprintf("TRK-%s-%d", userID, time.Now().Unix())
}
