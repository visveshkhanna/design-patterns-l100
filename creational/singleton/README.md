## Singleton — One signature pen

Goal: ensure a class has only one instance and provide a global access point to it.

This example uses `sync.Once` to create a single `SignaturePen` instance. Multiple calls to `GetSignaturePen()` return the same object, verified by identity comparison in `main.go`.

Run:
```bash
go run .
```

Files:
- `main.go`: singleton type, `GetSignaturePen`, and E2E demo

Why this design:
- **Thread-safe**: `sync.Once` guarantees one-time initialization
- **Simple API**: a single accessor function
- **Demonstrably single**: prints the equality of two retrieved instances


