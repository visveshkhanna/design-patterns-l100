## Factory Method — Create Guns by type

Goal: define an interface for creating an object, but let subclasses decide which class to instantiate.

This example exposes `getGun(type)` which returns an `IGun`. Different concrete implementations (`Ak47`, `musket`) are selected by a simple key; client only depends on the interface.

Run:
```bash
go run .
```

Files:
- `iGun.go`: product interface
- `gun.go`: base fields for products
- `ak47.go`, `musket.go`: concrete products
- `gunFactory.go`: selects which product to instantiate
- `main.go`: E2E demo

Why this design:
- **Decouples creation from usage**: callers ask for behavior, not classes
- **Extensible**: add a new gun by implementing `IGun` and updating the selector
- **Minimal**: tiny interface and simple key-based dispatch


