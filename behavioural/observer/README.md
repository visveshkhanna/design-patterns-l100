### Observer — Budget Alerts

Observers subscribe to spending updates and get notified when thresholds are crossed.

Run:
```bash
go run ./behavioural/observer
```

Why Observer here
- Decouples budget updates from alert channels; multiple alerts can react independently.

