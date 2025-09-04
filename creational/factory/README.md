## Factory Method — Create Guns by type

Goal: define an interface for creating an object, but let subclasses decide which class to instantiate.

This example exposes `factory.GetGun(type)` which returns a `Gun` interface. Different concrete implementations (`AK47`, `Musket`) are selected by a simple key; client only depends on the interface.

Run:
```bash
go run .
```

Folder Structure:
- `interfaces/`: Contains the Gun interface definition
  - `gun.go`: product interface
- `products/`: Contains concrete gun implementations
  - `base.go`: base struct with common gun fields and methods
  - `ak47.go`, `musket.go`: concrete gun products
- `factory/`: Contains the factory logic
  - `gun_factory.go`: selects which product to instantiate
- `main.go`: demonstration of the factory pattern

Why this design:
- **Decouples creation from usage**: callers ask for behavior, not classes
- **Extensible**: add a new gun by implementing `Gun` interface and updating the factory
- **Organized**: clear separation of concerns with logical folder structure
- **Minimal**: simple interface and key-based dispatch


