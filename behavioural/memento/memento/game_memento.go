package memento

import (
	"fmt"
	"time"
)

type GameMemento struct {
	health    int
	level     int
	inventory []string
	position  Position
	timestamp time.Time
}

type Position struct {
	X, Y int
}

func NewGameMemento(health, level int, inventory []string, position Position) *GameMemento {

	inventoryCopy := make([]string, len(inventory))
	copy(inventoryCopy, inventory)

	return &GameMemento{
		health:    health,
		level:     level,
		inventory: inventoryCopy,
		position:  position,
		timestamp: time.Now(),
	}
}

func (m *GameMemento) GetHealth() int {
	return m.health
}

func (m *GameMemento) GetLevel() int {
	return m.level
}

func (m *GameMemento) GetInventory() []string {
	inventoryCopy := make([]string, len(m.inventory))
	copy(inventoryCopy, m.inventory)
	return inventoryCopy
}

func (m *GameMemento) GetPosition() Position {
	return m.position
}

func (m *GameMemento) GetTimestamp() time.Time {
	return m.timestamp
}

func (m *GameMemento) GetSaveInfo() string {
	return fmt.Sprintf("Level %d, Health %d, Items: %d, Saved: %s",
		m.level, m.health, len(m.inventory), m.timestamp.Format("15:04:05"))
}
