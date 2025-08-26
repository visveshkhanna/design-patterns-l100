## Decorator — Add behavior without changing the original type

Goal: attach additional responsibilities to an object dynamically.

This example starts with an `Espresso` and decorates it with `Milk` and `Sugar`. Each decorator wraps a `Beverage` and augments `Description()` and `Cost()` while keeping the same interface.

Run:
```bash
go run .
```

Why this design:
- **Open/Closed**: extend behavior via composition, not modification
- **Stackable**: any order/combination of decorators
- **Minimal**: tiny interface keeps focus on composition


