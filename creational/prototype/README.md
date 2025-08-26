## Prototype — Clone a filesystem-like tree

Goal: create new objects by copying existing instances (prototypes) instead of building from scratch.

This example models a tiny filesystem. Every `Inode` (`File` or `Folder`) implements `clone()` and `print()`. Cloning a folder produces a deep copy with cloned children, demonstrated by printing both original and clone hierarchies.

Run:
```bash
go run .
```

Files:
- `inode.go`: prototype interface
- `file.go`: leaf node clone/print
- `folder.go`: composite node deep-clone/print
- `main.go`: E2E demo building and cloning a tree

Why this design:
- **Fast duplication**: copy existing structures without reconfiguration
- **Encapsulated copy logic**: each type knows how to clone itself
- **Simple and visible**: console output makes deep clone behavior obvious


