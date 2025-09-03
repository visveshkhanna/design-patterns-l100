package originator

import (
	"fmt"
	"strings"

	"design-patterns/behavioural/memento/memento"
)

type GameState struct {
	health    int
	level     int
	inventory []string
	position  memento.Position
}

func NewGameState() *GameState {
	return &GameState{
		health:    100,
		level:     1,
		inventory: []string{"Wooden Sword", "Health Potion"},
		position:  memento.Position{X: 0, Y: 0},
	}
}

func (g *GameState) TakeDamage(damage int) {
	g.health -= damage
	if g.health < 0 {
		g.health = 0
	}
	fmt.Printf("Took %d damage! Health: %d\n", damage, g.health)
}

func (g *GameState) Heal(amount int) {
	g.health += amount
	if g.health > 100 {
		g.health = 100
	}
	fmt.Printf("Healed %d points! Health: %d\n", amount, g.health)
}

func (g *GameState) LevelUp() {
	g.level++
	g.health = 100
	fmt.Printf("Level up! Now level %d with full health!\n", g.level)
}

func (g *GameState) AddItem(item string) {
	g.inventory = append(g.inventory, item)
	fmt.Printf("Found %s! Added to inventory.\n", item)
}

func (g *GameState) UseItem(item string) bool {
	for i, invItem := range g.inventory {
		if invItem == item {
			g.inventory = append(g.inventory[:i], g.inventory[i+1:]...)
			fmt.Printf("Used %s!\n", item)
			return true
		}
	}
	fmt.Printf("Don't have %s in inventory!\n", item)
	return false
}

func (g *GameState) Move(deltaX, deltaY int) {
	g.position.X += deltaX
	g.position.Y += deltaY
	fmt.Printf("Moved to position (%d, %d)\n", g.position.X, g.position.Y)
}

func (g *GameState) CreateSave() *memento.GameMemento {
	return memento.NewGameMemento(g.health, g.level, g.inventory, g.position)
}

func (g *GameState) LoadSave(save *memento.GameMemento) {
	g.health = save.GetHealth()
	g.level = save.GetLevel()
	g.inventory = save.GetInventory()
	g.position = save.GetPosition()
	fmt.Printf("Game loaded from save: %s\n", save.GetSaveInfo())
}

func (g *GameState) ShowStatus() {
	fmt.Printf("\n=== Game Status ===\n")
	fmt.Printf("Health: %d/100\n", g.health)
	fmt.Printf("Level: %d\n", g.level)
	fmt.Printf("Position: (%d, %d)\n", g.position.X, g.position.Y)
	fmt.Printf("Inventory: [%s]\n", strings.Join(g.inventory, ", "))
	fmt.Printf("==================\n\n")
}

func (g *GameState) IsGameOver() bool {
	return g.health <= 0
}
