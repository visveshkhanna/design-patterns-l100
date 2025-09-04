## Builder — Construct a House step-by-step

Goal: separate the construction of a complex object from its representation so the same construction process can create different variants.

This example builds a `House` by letting a `Director` orchestrate build steps defined by `Builder` interface. Two builders (`NormalBuilder`, `IglooBuilder`) produce different houses with the same sequence of steps.

Run:
```bash
go run .
```

## Structure:

```
builder/
├── main.go                 # Demo application
├── interfaces/
│   └── builder.go         # Builder interface
├── models/
│   └── house.go          # House product
├── builders/
│   ├── factory.go        # Builder factory
│   ├── normal_builder.go # Concrete normal house builder
│   └── igloo_builder.go  # Concrete igloo house builder
└── director/
    └── director.go       # Coordinates build sequence
```

## Components:

- **`interfaces/builder.go`**: `Builder` interface defining construction steps
- **`builders/`**: Concrete builders implementing different house types
- **`director/director.go`**: Orchestrates the build sequence
- **`models/house.go`**: The product being built
- **`main.go`**: End-to-end demonstration

## Why this design:
- **Stable process, varied products**: Director drives steps; builders vary details
- **Extensible**: add builders without changing the Director or client code
- **Clean separation**: Each component has a single responsibility
- **Simple**: Clear API keeps focus on the pattern, not domain specifics


