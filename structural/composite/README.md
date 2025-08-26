## Composite — Treat part-whole hierarchies uniformly

Goal: compose objects into tree structures and let clients treat individual objects and compositions uniformly.

This example builds an order consisting of items and boxes of items. Both implement a common interface so you can `Print()` and get `Price()` from either a single `Item` or a `Box` containing many nested elements.

Run:
```bash
go run .
```

Why this design:
- **Uniformity**: leaves and composites share the same interface
- **Recursive operations**: price/print work the same at any level
- **Practical**: mirrors filesystem, menus, shopping carts


