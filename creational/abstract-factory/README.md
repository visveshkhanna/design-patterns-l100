## Abstract Factory — Themed UI (Light/Dark)

Goal: create families of related components without hard-coding their concrete classes.

This example builds a minimal UI with a `ThemeFactory` that produces a matching `Button` and `Card` for either a light or dark theme. The client uses only the abstract factory and product interfaces; swapping the family is one line.

Run:
```bash
go run .
```

Files:
- `primitives.go`: `ThemeFactory` + product interfaces and light/dark implementations
- `main.go`: E2E demo rendering both themes

Why this design:
- Ensures coherent families (light components always match light cards/buttons)
- Easy to add new themes without touching client code
- Minimal API keeps focus on the pattern, not UI code


