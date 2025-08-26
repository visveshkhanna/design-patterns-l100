## Builder — Construct a House step-by-step

Goal: separate the construction of a complex object from its representation so the same construction process can create different variants.

This example builds a `House` by letting a `Director` orchestrate build steps defined by `IBuilder`. Two builders (`NormalBuilder`, `IglooBuilder`) produce different houses with the same sequence of steps.

Run:
```bash
go run .
```

Files:
- `iBuilder.go`: `IBuilder` interface + `getBuilder`
- `normalBuilder.go`, `iglooBuilder.go`: concrete builders
- `director.go`: coordinates the build sequence
- `house.go`: product
- `main.go`: E2E demo

Why this design:
- **Stable process, varied products**: Director drives steps; builders vary details
- **Extensible**: add builders without changing the Director or client code
- **Simple**: tiny API keeps focus on the pattern, not domain specifics


