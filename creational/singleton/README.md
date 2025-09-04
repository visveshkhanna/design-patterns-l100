## Singleton — One signature pen

Goal: ensure a class has only one instance and provide a global access point to it.

This example uses `sync.Once` to create a single `SignaturePen` instance. Multiple calls to `GetSignaturePen()` return the same object, verified by identity comparison in `main.go`.

Run:
```bash
go run .
```

## Structure

```
singleton/
├── models/
│   └── signature_pen.go    # SignaturePen struct and methods
├── singleton/
│   └── pen_manager.go      # Singleton logic with GetSignaturePen()
├── main.go                 # Demo and usage example
└── README.md              # This file
```

## Files

- `models/signature_pen.go`: Defines the `SignaturePen` struct and its methods
- `singleton/pen_manager.go`: Contains singleton logic using `sync.Once`
- `main.go`: Demonstrates singleton usage and verification

## Why this design

- **Thread-safe**: `sync.Once` guarantees one-time initialization
- **Simple API**: a single accessor function `GetSignaturePen()`
- **Separation of concerns**: Model, singleton logic, and demo are in separate files
- **Demonstrably single**: prints the equality of two retrieved instances


