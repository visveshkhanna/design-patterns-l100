### Chain of Responsibility — Message Moderation

This example models a tiny message moderation pipeline for a chat app. Each handler either rejects a message or enriches it and passes it along. The chain is composable and each step is isolated.

Handlers used:
- ProfanityFilter: rejects messages containing banned words.
- SpamDetector: rejects obvious spam.
- LengthLimiter: trims long messages and adds a "trimmed" flag.
- AuditLogger: appends an "audited" flag.

Run:
```bash
go run ./behavioural/chainofres
```

Expected output (similar):
```text
accepted: Hello team, standup at 10 flags: audited
rejected: spam detected
accepted: This is a darn long mess flags: trimmed,audited
```

Why CoR here
- Decouples moderation steps; you can add/remove/reorder filters without changing others.
- Keeps handlers simple and testable.

