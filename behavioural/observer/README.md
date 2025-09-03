### Observer — Budget Alerts

The Observer pattern allows objects (observers) to subscribe to and receive notifications about changes in another object (subject). In this example, different alert systems observe budget spending and react when their thresholds are reached.

## Structure

- `interfaces/` - Contains Observer and Subject interface definitions
- `subjects/` - Contains Budget implementation (the observable subject)  
- `observers/` - Contains SMS and Email alert implementations
- `main.go` - Demonstrates the pattern with budget tracking

## How it works

1. Budget is the **Subject** that maintains spending state
2. SMS and Email alerts are **Observers** that watch for spending changes
3. When spending is added, all registered observers are notified
4. Each observer can react differently based on their own thresholds

Run:
```bash
go run ./behavioural/observer
```

## Why Observer Pattern

- **Decoupling**: Budget updates are decoupled from alert channels
- **Extensibility**: Easy to add new types of alerts without changing existing code
- **Independence**: Multiple alerts can react independently with different thresholds

