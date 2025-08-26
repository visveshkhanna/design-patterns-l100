### Template Method — Webhook Delivery

Defines the skeleton of a webhook delivery (`prepare -> dispatch -> report`) and lets implementations fill the steps.

Run:
```bash
go run ./behavioural/template
```

Why Template here
- Reuses the algorithm structure while varying step implementations.

