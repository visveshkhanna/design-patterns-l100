### Interpreter — Access Rule Evaluator

A tiny boolean expression language interprets access rules against a request context with a complete parser implementation.

Features:
- **Parser**: Converts string expressions into executable expression trees
- **Lexer**: Tokenizes input expressions  
- **Grammar**: Supports boolean logic with proper operator precedence
- **Expressions**: Role-based, region-based, feature flag, and numeric comparisons

Grammar (EBNF):
```
expression     = or_expression
or_expression  = and_expression ("OR" and_expression)*
and_expression = primary ("AND" primary)*
primary        = "(" expression ")" | comparison | identifier
comparison     = identifier ("==" string | "<" number)
identifier     = [a-zA-Z_][a-zA-Z0-9_]*
```

Supported Syntax:
- Equality: `role==admin`, `region==EU`
- Boolean flags: `beta`
- Numeric comparison: `requests<100`
- Logical operators: `AND`, `OR`
- Grouping: `(expression)`

Examples:
```
role==admin
beta AND region==EU
(role==admin OR beta) AND requests<100
```

Run:
```bash
go run ./behavioural/interpreter
```

Why Interpreter here:
- Encodes rule logic as a tree of expressions; new predicates/combinators can be added without changing existing ones
- Parser allows dynamic rule creation from string expressions
- Clean separation between parsing logic and evaluation logic

