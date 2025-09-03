### Memento Pattern — Game Save System

The Memento pattern captures and externalizes an object's internal state so that it can be restored later, without violating encapsulation.

## Implementation

This example demonstrates a game save system with three key components:

- **Originator** (`GameState`): The game that can create and restore from save snapshots
- **Memento** (`GameMemento`): Immutable snapshot of game state (health, level, inventory, position)
- **Caretaker** (`SaveManager`): Manages multiple save slots without accessing memento internals

## Key Features

- **Multiple save slots** with slot management
- **Quick save** functionality
- **Save metadata** (timestamp, game info)
- **Complete state restoration** including health, level, inventory, and position
- **Encapsulation** - external code cannot access game internals directly

## Run

```bash
go run ./behavioural/memento
```

## Why Memento Here

- **Undo/Redo capability**: Players can return to previous game states
- **Safe state management**: Game internals remain encapsulated
- **Flexible save system**: Caretaker can implement various save strategies (auto-save, manual saves, etc.)
- **Immutable snapshots**: Prevents accidental modification of saved states

The pattern separates state capture from state management, allowing complex save systems without exposing game implementation details.