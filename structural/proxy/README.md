## Proxy — Control access with caching and rules

Goal: provide a surrogate or placeholder to control access to another object.

This example adds a `CachingProxy` in front of an `OriginServer`. The proxy caches responses and can block certain URLs. Clients call `Fetch` on the proxy instead of the origin; the proxy decides whether to serve from cache, fetch fresh, or deny.

Run:
```bash
go run .
```

Why this design:
- **Performance**: cache frequent requests
- **Policy enforcement**: block/allow lists without changing the origin
- **Transparency**: client speaks the same interface


