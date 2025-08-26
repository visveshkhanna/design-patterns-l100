### State — Document Workflow

Models a document moving through `Draft -> Review -> Published` with explicit state objects controlling allowed transitions.

Run:
```bash
go run ./behavioural/state
```

Why State here
- Encapsulates transition rules inside states; caller uses a stable API (`submit/approve/reject`).

