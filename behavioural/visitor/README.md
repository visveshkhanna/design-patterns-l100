### Visitor — Game Map Printer

Separates operations from data. We visit `Forest` and `City` zones with a `MapPrinter` to print summaries.

Run:
```bash
go run ./behavioural/visitor
```

Why Visitor here
- Add new operations (e.g., `ResourceCounter`) without changing the zone types.

