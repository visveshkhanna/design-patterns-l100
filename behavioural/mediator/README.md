### Mediator — Greenhouse Controller

Coordinates interactions between a temperature sensor, a fan, and a heater. The mediator (`Greenhouse`) centralizes responses to events (`hot`, `cold`, `ok`).

## Project Structure

```
mediator/
├── interfaces/
│   └── mediator.go       # Mediator interface definition
├── components/
│   ├── fan.go           # Fan component
│   ├── heater.go        # Heater component
│   └── sensor.go        # Temperature sensor component
├── mediator/
│   └── greenhouse.go    # Concrete greenhouse mediator implementation
├── main.go              # Example usage
└── README.md
```

## Run
```bash
go run ./behavioural/mediator
```

## Why Mediator Pattern
- **Loose Coupling**: Avoids tight coupling between `Sensor`, `Fan`, and `Heater`
- **Centralized Logic**: All interaction logic lives in the mediator (`Greenhouse`)
- **Easier Maintenance**: Adding new components or changing interactions only requires modifying the mediator
- **Single Responsibility**: Each component only knows about the mediator, not other components

