### Command — Smart Coffee Machine

Encapsulates requests (add water, add beans, brew, clean) as command objects. A `Button` triggers a single command, while a `Remote` queues multiple commands to run a recipe.

Commands:
- AddWaterCommand, AddBeansCommand, BrewEspressoCommand, CleanCommand

Run:
```bash
go run ./behavioural/command
```

Why Command here
- Decouples UI invokers from machine operations; recipes can be re-ordered/queued.

