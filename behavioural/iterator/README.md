### Iterator — Log Entries

Iterates over structured log entries without exposing the underlying slice. The iterator returns `LogEntry` values (level + message).

Run:
```bash
go run ./behavioural/iterator
```

Why Iterator here
- Provides a simple traversal API; future filters or pagination can be added without changing clients.

