### Interpreter — Access Rule Evaluator

A tiny boolean expression language interprets access rules against a request context.

Grammar (informal):
- Predicates: `IsRole(admin)`, `IsRegion(EU)`, `IsBeta`, `RequestsBelow(100)`
- Combinators: `And(a, b)`, `Or(a, b)`

Run:
```bash
go run ./behavioural/interpreter
```

Why Interpreter here
- Encodes rule logic as a tree of expressions; new predicates/combinators can be added without changing existing ones.

