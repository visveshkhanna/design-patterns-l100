## Abstract Factory — Themed UI (Light/Dark)

Goal: create families of related components without hard-coding their concrete classes.

This example builds a minimal UI with a `ThemeFactory` that produces a matching `Button` and `Card` for either a light or dark theme. The client uses only the abstract factory and product interfaces; swapping the family is one line.

Run:
```bash
go run .
```

## Structure:
- `interfaces/`: Core interfaces (Button, Card, ThemeFactory)
- `themes/`: Concrete implementations for light and dark themes
- `factory/`: Factory function to create theme instances
- `main.go`: Demo showing both themes in action

## Why this design:
- **Separation of concerns**: Each folder has a clear responsibility
- **Coherent families**: Light components always match light cards/buttons
- **Easy extensibility**: Add new themes without touching client code
- **Clean imports**: Clear dependencies between packages
- **Minimal API**: Keeps focus on the pattern, not implementation details


