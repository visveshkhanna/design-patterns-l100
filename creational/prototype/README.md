## Prototype — Clone a filesystem-like tree

Goal: create new objects by copying existing instances (prototypes) instead of building from scratch.

This example models a tiny filesystem. Every `Inode` (`File` or `Folder`) implements `clone()` and `print()`. Cloning a folder produces a deep copy with cloned children, demonstrated by printing both original and clone hierarchies.

Run:
```bash
go run .
```

Files:
- `interfaces/prototype.go`: prototype interface definition
- `models/file.go`: file implementation with clone/print methods
- `models/folder.go`: folder implementation with deep-clone/print methods
- `main.go`: demonstration of building and cloning a filesystem tree

Why this design:
- **Fast duplication**: copy existing structures without reconfiguration
- **Encapsulated copy logic**: each type knows how to clone itself
- **Simple and visible**: console output makes deep clone behavior obvious


