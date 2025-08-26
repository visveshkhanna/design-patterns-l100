### Mediator — Greenhouse Controller

Coordinates interactions between a temperature sensor, a fan, and a heater. The mediator (`Greenhouse`) centralizes responses to events (`hot`, `cold`, `ok`).

Run:
```bash
go run ./behavioural/mediator
```

Why Mediator here
- Avoids tight coupling between `Sensor`, `Fan`, and `Heater`; logic for reactions lives in one place.

