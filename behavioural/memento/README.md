### Memento — TextField History

Captures `TextField` snapshots (content + cursor) to enable undo-like behavior without exposing internals.

Run:
```bash
go run ./behavioural/memento
```

Why Memento here
- Separates snapshot storage from the `TextField`; allows restoring previous states safely.

