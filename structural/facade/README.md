## Facade — Simple API over complex subsystems

Goal: provide a unified interface to a set of interfaces in a subsystem.

This example exposes an `OrderFacade` that coordinates inventory, payment, and shipping. The client only calls `PlaceOrder`, while the facade orchestrates the multiple services and error handling.

Run:
```bash
go run .
```

Why this design:
- **Ease of use**: a single call instead of many
- **Encapsulation**: hides subsystem complexity and ordering
- **Safer**: centralizes validations and failure paths


