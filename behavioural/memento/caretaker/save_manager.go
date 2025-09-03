package caretaker

import (
	"design-patterns/behavioural/memento/memento"
	"fmt"
)

type SaveManager struct {
	saves    map[int]*memento.GameMemento
	maxSlots int
}

func NewSaveManager(maxSlots int) *SaveManager {
	return &SaveManager{
		saves:    make(map[int]*memento.GameMemento),
		maxSlots: maxSlots,
	}
}

func (sm *SaveManager) SaveGame(slot int, gameMemento *memento.GameMemento) error {
	if slot < 1 || slot > sm.maxSlots {
		return fmt.Errorf("invalid save slot %d. Must be between 1 and %d", slot, sm.maxSlots)
	}

	sm.saves[slot] = gameMemento
	fmt.Printf("Game saved to slot %d: %s\n", slot, gameMemento.GetSaveInfo())
	return nil
}

func (sm *SaveManager) LoadGame(slot int) (*memento.GameMemento, error) {
	if slot < 1 || slot > sm.maxSlots {
		return nil, fmt.Errorf("invalid save slot %d. Must be between 1 and %d", slot, sm.maxSlots)
	}

	save, exists := sm.saves[slot]
	if !exists {
		return nil, fmt.Errorf("no save found in slot %d", slot)
	}

	fmt.Printf("Loading game from slot %d: %s\n", slot, save.GetSaveInfo())
	return save, nil
}

func (sm *SaveManager) DeleteSave(slot int) error {
	if slot < 1 || slot > sm.maxSlots {
		return fmt.Errorf("invalid save slot %d. Must be between 1 and %d", slot, sm.maxSlots)
	}

	if _, exists := sm.saves[slot]; !exists {
		return fmt.Errorf("no save found in slot %d", slot)
	}

	delete(sm.saves, slot)
	fmt.Printf("Save deleted from slot %d\n", slot)
	return nil
}

func (sm *SaveManager) ListSaves() {
	fmt.Printf("\n=== Save Slots ===\n")
	hasAnySaves := false

	for slot := 1; slot <= sm.maxSlots; slot++ {
		if save, exists := sm.saves[slot]; exists {
			fmt.Printf("Slot %d: %s\n", slot, save.GetSaveInfo())
			hasAnySaves = true
		} else {
			fmt.Printf("Slot %d: [Empty]\n", slot)
		}
	}

	if !hasAnySaves {
		fmt.Printf("No saves found.\n")
	}
	fmt.Printf("==================\n\n")
}

func (sm *SaveManager) HasSave(slot int) bool {
	if slot < 1 || slot > sm.maxSlots {
		return false
	}
	_, exists := sm.saves[slot]
	return exists
}

func (sm *SaveManager) GetSaveCount() int {
	return len(sm.saves)
}

func (sm *SaveManager) GetMaxSlots() int {
	return sm.maxSlots
}

func (sm *SaveManager) QuickSave(gameMemento *memento.GameMemento) (int, error) {

	for slot := 1; slot <= sm.maxSlots; slot++ {
		if !sm.HasSave(slot) {
			err := sm.SaveGame(slot, gameMemento)
			return slot, err
		}
	}

	err := sm.SaveGame(1, gameMemento)
	fmt.Printf("All slots full, overwrote slot 1\n")
	return 1, err
}

func (sm *SaveManager) GetOldestSave() (int, *memento.GameMemento, error) {
	if len(sm.saves) == 0 {
		return 0, nil, fmt.Errorf("no saves available")
	}

	var oldestSlot int
	var oldestSave *memento.GameMemento

	for slot, save := range sm.saves {
		if oldestSave == nil || save.GetTimestamp().Before(oldestSave.GetTimestamp()) {
			oldestSlot = slot
			oldestSave = save
		}
	}

	return oldestSlot, oldestSave, nil
}

func (sm *SaveManager) GetNewestSave() (int, *memento.GameMemento, error) {
	if len(sm.saves) == 0 {
		return 0, nil, fmt.Errorf("no saves available")
	}

	var newestSlot int
	var newestSave *memento.GameMemento

	for slot, save := range sm.saves {
		if newestSave == nil || save.GetTimestamp().After(newestSave.GetTimestamp()) {
			newestSlot = slot
			newestSave = save
		}
	}

	return newestSlot, newestSave, nil
}
