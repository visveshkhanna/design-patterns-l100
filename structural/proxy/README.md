## Proxy Pattern — Control access with caching and rules

**Goal**: Provide a surrogate or placeholder to control access to another object.

This example demonstrates a `CachingProxy` that sits in front of an origin server. The proxy adds caching and URL blocking capabilities while maintaining the same interface as the origin server.

### Structure
```
proxy/
├── interfaces/          # Common interface definitions
│   └── fetcher.go       # Fetcher interface
├── origin/              # Real subject implementation
│   └── server.go        # OriginServer that does actual fetching
├── proxy/               # Proxy implementation
│   └── caching_proxy.go # CachingProxy with caching and blocking
├── main.go             # Demo usage
└── README.md           # This file
```

### Run
```bash
go run .
```

### Key Benefits
- **Performance**: Cache frequent requests to avoid repeated work
- **Policy enforcement**: Block/allow URLs without changing the origin server
- **Transparency**: Client uses the same interface whether talking to proxy or origin
- **Control**: Add logging, authentication, or other cross-cutting concerns


