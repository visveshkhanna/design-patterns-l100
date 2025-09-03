package main

import (
	"fmt"

	"design-patterns/behavioural/memento/caretaker"
	"design-patterns/behavioural/memento/originator"
)

func main() {
	fmt.Println("=== Memento Pattern: Game Save System ===")

	game := originator.NewGameState()
	saveManager := caretaker.NewSaveManager(3)

	fmt.Println("Starting new game...")
	game.ShowStatus()

	fmt.Println("Playing the game...")
	game.Move(5, 3)
	game.AddItem("Magic Ring")
	game.TakeDamage(20)

	save1 := game.CreateSave()
	saveManager.SaveGame(1, save1)

	fmt.Println("\nContinuing adventure...")
	game.Move(2, -1)
	game.LevelUp()
	game.AddItem("Dragon Sword")
	game.UseItem("Health Potion")

	save2 := game.CreateSave()
	saveManager.SaveGame(2, save2)

	game.ShowStatus()

	fmt.Println("Oh no! Encountered a powerful enemy!")
	game.TakeDamage(70)
	game.TakeDamage(50)

	if game.IsGameOver() {
		fmt.Println("GAME OVER! Loading previous save...")

		saveManager.ListSaves()

		if loadedSave, err := saveManager.LoadGame(2); err == nil {
			game.LoadSave(loadedSave)
			game.ShowStatus()
		}
	}

	fmt.Println("Demonstrating save management...")

	game.AddItem("Health Elixir")
	slot, _ := saveManager.QuickSave(game.CreateSave())
	fmt.Printf("Quick saved to slot %d\n", slot)

	saveManager.ListSaves()

	fmt.Println("Loading earlier save from slot 1...")
	if earlierSave, err := saveManager.LoadGame(1); err == nil {
		game.LoadSave(earlierSave)
		game.ShowStatus()
	}

	if newestSlot, newestSave, err := saveManager.GetNewestSave(); err == nil {
		fmt.Printf("Newest save is in slot %d: %s\n", newestSlot, newestSave.GetSaveInfo())
	}

	if oldestSlot, oldestSave, err := saveManager.GetOldestSave(); err == nil {
		fmt.Printf("Oldest save is in slot %d: %s\n", oldestSlot, oldestSave.GetSaveInfo())
	}

	fmt.Println("\nDeleting save from slot 1...")
	saveManager.DeleteSave(1)
	saveManager.ListSaves()

	fmt.Println("\n=== Memento Pattern Demo Complete ===")
	fmt.Println("Key benefits demonstrated:")
	fmt.Println("- Game state encapsulation (originator)")
	fmt.Println("- Immutable save snapshots (memento)")
	fmt.Println("- External save management (caretaker)")
	fmt.Println("- No external access to game internals")
}
